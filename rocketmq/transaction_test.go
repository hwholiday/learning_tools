package rocketmq

import (
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func Test(t *testing.T) {
	rocketmq.NewTransactionProducer(&UserListener{
		localTrans: new(sync.Map),
	}, producer.WithNsResovler(primitive.NewPassthroughResolver([]string{"172.13.3.160:9876"})),
		producer.WithRetry(2))
}

type UserListener struct {
	localTrans       *sync.Map
	transactionIndex int32
}

func (dl *UserListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	fmt.Println("开始执行本地事务")
	var nextIndex = atomic.AddInt32(&dl.transactionIndex, 1)
	fmt.Printf("nextIndex: %v for transactionID: %v\n", nextIndex, msg.TransactionId)
	status := nextIndex % 3
	dl.localTrans.Store(msg.TransactionId, primitive.LocalTransactionState(status+1))
	fmt.Println("开始执行本地事务 结束", status+1)
	return primitive.UnknowState
}

func (dl *UserListener) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	fmt.Println("开始回查本地事务状态")
	fmt.Printf("%v msg transactionID : %v\n", time.Now(), msg.TransactionId)
	v, existed := dl.localTrans.Load(msg.TransactionId)
	if !existed {
		fmt.Printf("unknow msg: %v, return Commit", msg)
		return primitive.CommitMessageState
	}
	state := v.(primitive.LocalTransactionState)
	switch state {
	case 1:
		fmt.Printf("结束本地事务状态查询 checkLocalTransaction COMMIT_MESSAGE: %v\n", msg)
		return primitive.CommitMessageState
	case 2:
		fmt.Printf("结束本地事务状态查询 checkLocalTransaction ROLLBACK_MESSAGE: %v\n", msg)
		return primitive.RollbackMessageState
	case 3:
		fmt.Printf("结束本地事务状态查询 checkLocalTransaction unknow: %v\n", msg)
		return primitive.UnknowState
	default:
		fmt.Printf("结束本地事务状态查询 checkLocalTransaction default COMMIT_MESSAGE: %v\n", msg)
		return primitive.CommitMessageState
	}
}
