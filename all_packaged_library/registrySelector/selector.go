package registrySelector

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Selector interface {
	Next() (Node, error)
}

type selectorServer struct {
	cli     *clientv3.Client
	node    []Node
	options SelectorOptions
}

type SelectorOptions struct {
	name   string
	config clientv3.Config
}

func NewSelector(options SelectorOptions) (Selector, error) {
	cli, err := clientv3.New(options.config)
	if err != nil {
		return nil, err
	}
	var s = &selectorServer{
		options: options,
		cli:     cli,
	}
	go s.Watch()
	return s, nil
}

func (s *selectorServer) Next() (Node, error) {
	if len(s.node) == 0 {
		return Node{}, fmt.Errorf("no node found on the %s", s.options.name)
	}
	i := rand.Int() % len(s.node)
	return s.node[i], nil
}

func (s *selectorServer) Watch() {
	res, err := s.cli.Get(context.TODO(), s.GetKey(), clientv3.WithPrefix(), clientv3.WithSerializable())
	if err != nil {
		log.Printf("[Watch] err : %s", err.Error())
		return
	}
	for _, kv := range res.Kvs {
		node, err := s.GetVal(kv.Value)
		if err != nil {
			log.Printf("[GetVal] err : %s", err.Error())
			continue
		}
		s.node = append(s.node, node)
	}
	ch := s.cli.Watch(context.TODO(), prefix, clientv3.WithPrefix())
	for {
		select {
		case c := <-ch:
			for _, e := range c.Events {
				switch e.Type {
				case clientv3.EventTypePut:
					node, err := s.GetVal(e.Kv.Value)
					if err != nil {
						log.Printf("[EventTypePut] err : %s", err.Error())
						continue
					}
					s.AddNode(node)
				case clientv3.EventTypeDelete:
					keyArray := strings.Split(string(e.Kv.Key), "/")
					if len(keyArray) <= 0 {
						log.Printf("[EventTypeDelete] key Split err : %s", err.Error())
						return
					}
					nodeId, err := strconv.Atoi(keyArray[len(keyArray)-1])
					if err != nil {
						log.Printf("[EventTypePut] key Atoi : %s", err.Error())
						continue
					}
					s.DelNode(uint32(nodeId))
				}
			}
		}
	}
}

func (s *selectorServer) DelNode(id uint32) {
	var node []Node
	for _, v := range s.node {
		if v.Id != id {
			node = append(node, v)
		}
	}
	s.node = node
}

func (s *selectorServer) AddNode(node Node) {
	var exist bool
	for _, v := range s.node {
		if v.Id == node.Id {
			exist = true
		}
	}
	if !exist {
		s.node = append(s.node, node)
	}
}

func (s *selectorServer) GetKey() string {
	return fmt.Sprintf("%s%s", prefix, s.options.name)
}

func (s *selectorServer) GetVal(val []byte) (Node, error) {
	var node Node
	err := json.Unmarshal(val, &node)
	if err != nil {
		return node, err
	}
	return node, nil
}
