package discovery

import (
	"fmt"
	"github.com/hwholiday/learning_tools/hlb-grpc/register"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/resolver"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const CustomizeLB = "customize"
const customizeCtx = "customizeCtx"

var limit sync.Map

// NewBuilder creates a new weight balancer builder.
func newCustomizeBuilder(opt *Options) {
	//balancer.Builder
	builder := base.NewBalancerBuilderV2(CustomizeLB, &rrPickerBuilder{opt: opt}, base.Config{HealthCheck: true})
	balancer.Register(builder)
	return
}

//move discovery init
/*func init() {
	newBuilder(nil)
}*/

type rrPickerBuilder struct {
	opt *Options // discovery Options info
}

func (r *rrPickerBuilder) Build(info base.PickerBuildInfo) balancer.V2Picker {
	if len(info.ReadySCs) == 0 {
		return base.NewErrPickerV2(balancer.ErrNoSubConnAvailable)
	}
	var scs = make(map[balancer.SubConn]*register.Options, len(info.ReadySCs))
	for conn, addr := range info.ReadySCs {
		nodeInfo := GetNodeInfo(addr.Address)
		if nodeInfo != nil {
			scs[conn] = nodeInfo
		}
	}
	if len(scs) == 0 {
		return base.NewErrPickerV2(balancer.ErrNoSubConnAvailable)
	}
	return &rrPicker{
		node: scs,
		opt:  r.opt,
	}
}

type rrPicker struct {
	node map[balancer.SubConn]*register.Options
	mu   sync.Mutex
	opt  *Options // discovery Options info
}

func (p *rrPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	var subConns []balancer.SubConn
	if len(p.node) == 0 {
		return balancer.PickResult{}, ErrNoMatchFoundConn
	}
	for conn, node := range p.node {
		if p.filterNode(node, p.buildFilter(info)) {
			subConns = append(subConns, conn)
		}
	}
	if len(subConns) == 0 {
		return balancer.PickResult{}, ErrNotGetRightConn
	}
	index := rand.Intn(len(subConns))
	sc := subConns[index]
	/*return balancer.PickResult{SubConn: sc, Done: func(data balancer.DoneInfo) {
		fmt.Println("test", info.FullMethodName, "end", data.Err, "time", time.Now().UnixNano()/1e6-t)
	}}, nil*/
	return balancer.PickResult{SubConn: sc}, nil
}

func (p *rrPicker) buildFilter(info balancer.PickInfo) map[string]string {
	filterData := getCtxFilter(info.Ctx)
	methodArray := strings.Split(info.FullMethodName, "/")
	filterData["method"] = methodArray[len(methodArray)-1]
	if _, ok := filterData["version"]; !ok {
		filterData["version"] = p.opt.Version
	}
	return filterData
}

func (p *rrPicker) filterNode(connNode *register.Options, filter map[string]string) bool {
	//node在线
	if !connNode.Node.Online {
		return false
	}
	//版本
	if filter["version"] != connNode.Node.Version {
		return false
	}
	for k, v := range connNode.Metadata {
		if val, ok := filter[k]; ok {
			if val != v {
				return false
			}
		}
	}
	var endpoint register.Endpoints
	for k, v := range connNode.Endpoints {
		if k == filter["method"] {
			endpoint = v
		}
	}
	//接口在线
	if !endpoint.Online {
		return false
	}
	//接口限流
	if endpoint.LimitingSwitch {
		//间隔
		t := time.Now().UnixNano() / 1e6
		key := fmt.Sprintf("%s/%s", connNode.Node.Id, filter["method"])
		val, ok := limit.Load(key)
		if ok {
			valInt := val.(int64)
			interval := 1000 / endpoint.Limiting
			if t-valInt < interval {
				return false
			} else {
				limit.Store(key, t)
			}
		} else {
			limit.Store(key, t)
		}
	}
	//TODO 接口熔断
	return true
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
