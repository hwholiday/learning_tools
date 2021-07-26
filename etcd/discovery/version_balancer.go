package discovery

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/resolver"
	"learning_tools/etcd/register"
	"math/rand"
	"sync"
	"time"
)

// NewBuilder creates a new weight balancer builder.
func newBuilder() balancer.Builder {
	return base.NewBalancerBuilder("version", &rrPickerBuilder{}, base.Config{HealthCheck: true})
}

func init() {
	balancer.Register(newBuilder())
}

type rrPickerBuilder struct{}

func (*rrPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}
	var scs = make(map[balancer.SubConn]*register.Options, len(info.ReadySCs))
	for conn, addr := range info.ReadySCs {
		nodeInfo := GetNodeInfo(addr.Address)
		if nodeInfo != nil {
			scs[conn] = nodeInfo
		}
	}
	if len(scs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}
	return &rrPicker{
		node: scs,
	}
}

type rrPicker struct {
	node map[balancer.SubConn]*register.Options
	mu   sync.Mutex
}

func (p *rrPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	p.mu.Lock()
	t := time.Now().UnixNano() / 1e6
	defer p.mu.Unlock()
	version := info.Ctx.Value("version")
	var subConns []balancer.SubConn
	for conn, node := range p.node {
		if version != "" {
			if node.Node.Version == version.(string) {
				subConns = append(subConns, conn)
			}
		}
	}
	if len(subConns) == 0 {
		return balancer.PickResult{}, errors.New("no match found conn")
	}
	index := rand.Intn(len(subConns))
	sc := subConns[index]
	return balancer.PickResult{SubConn: sc, Done: func(data balancer.DoneInfo) {
		fmt.Println("test", info.FullMethodName, "end", data.Err, "time", time.Now().UnixNano()/1e6-t)
	}}, nil
}

type attrKey struct{}

func SetNodeInfo(addr resolver.Address, hInfo *register.Options) resolver.Address {
	addr.Attributes = attributes.New()
	addr.Attributes = addr.Attributes.WithValues(attrKey{}, hInfo)
	return addr
}

func GetNodeInfo(attr resolver.Address) *register.Options {
	v := attr.Attributes.Value(attrKey{})
	hi, _ := v.(*register.Options)
	return hi
}
