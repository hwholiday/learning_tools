package discovery

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/resolver"
	"math/rand"
	"sync"
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
	var scs = make(map[balancer.SubConn]*NodeInfo, len(info.ReadySCs))
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
	node map[balancer.SubConn]*NodeInfo
	mu   sync.Mutex
}

func (p *rrPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	p.mu.Lock()
	version := info.Ctx.Value("version")
	var subConns []balancer.SubConn
	for conn, node := range p.node {
		fmt.Println("node", node)
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
	p.mu.Unlock()
	return balancer.PickResult{SubConn: sc}, nil
}

type attrKey struct{}

func SetNodeInfo(addr resolver.Address, hInfo *NodeInfo) resolver.Address {
	addr.Attributes = addr.Attributes.WithValues(attrKey{}, hInfo)
	return addr
}

func GetNodeInfo(attr resolver.Address) *NodeInfo {
	v := attr.Attributes.Value(attrKey{})
	hi, _ := v.(*NodeInfo)
	return hi
}
