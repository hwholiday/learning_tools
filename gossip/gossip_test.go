package gossip

import (
	"fmt"
	"github.com/hashicorp/memberlist"
	"os"
	"strconv"
	"testing"
	"time"
)

func Test_gossip(t *testing.T) {
	ports := []int{8001, 8002, 8003, 8004, 8005}
	for _, v := range ports {
		go gossip(v)
	}
	select {}
}

type MyDelegate struct {
	msgCh chan []byte
}

func (d *MyDelegate) NotifyMsg(msg []byte) {
	d.msgCh <- msg
}
func (d *MyDelegate) NodeMeta(limit int) []byte {
	// not use, noop
	return []byte("")
}
func (d *MyDelegate) LocalState(join bool) []byte {
	// not use, noop
	return []byte("")
}
func (d *MyDelegate) GetBroadcasts(overhead, limit int) [][]byte {
	// not use, noop
	return nil
}
func (d *MyDelegate) MergeRemoteState(buf []byte, join bool) {
	// not use
}

func gossip(bindPort int) {
	msgCh := make(chan []byte)
	d := new(MyDelegate)
	d.msgCh = msgCh
	hostname, _ := os.Hostname()
	config := memberlist.DefaultLocalConfig()
	config.Name = hostname + "-" + strconv.Itoa(bindPort)
	config.BindAddr = "127.0.0.1"
	config.BindPort = bindPort
	config.AdvertisePort = bindPort
	config.Delegate = d
	list, err := memberlist.Create(config)
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}
	// Join an ex
	//isting cluster by specifying at least one known member.
	//当节点第一次启动时，它会去查配置文件cassandra.yaml从而得到它属于的集群名称，
	//但是它如何获得集群中其他节点的信息呢？就是通过种子节点（seed node).记住，
	//同一集群中所有的节点的cassandra.yaml中必须有相同的种子节点列表。
	//选派谁做种子节点没什么特别的意义，仅仅在于新节点加入到集群中时走gossip流程时有用，所以它们没什么特权。
	_, err = list.Join([]string{"127.0.0.1:8001"})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}
	go func() {
		local := list.LocalNode()
		for {
			select {
			case info := <-msgCh:
				/*var data = make(map[string]interface{})
				if err := json.Unmarshal(info, &data); err != nil {
					panic(err)
				}*/
				fmt.Println("readinfo ", local.Port, string(info))

			}
		}
	}()
	for {
		for _, member := range list.Members() {
			data := member
			fmt.Printf("Member: %s %s %d\n", data.Name, data.Addr, data.Port)

		}
		for _, member := range list.Members() {
			data := member
			go list.SendReliable(data, []byte(fmt.Sprintf("member.Addr %s %d send", data.Addr, data.Port)))

		}
		time.Sleep(time.Second * 3)

	}
}
