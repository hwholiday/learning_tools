package library

import (
	"testing"
	"time"
)

func TestReporting1(t *testing.T) {
	c, err := InitNode([]string{"172.12.17.161:2379"}, "gateway", "1", "1111111111", 15)
	defer c.Close()
	if err != nil {
		panic(err)
	}
	t.Log(c.UpNode())
	time.Sleep(10 * time.Second)

}

func TestReporting2(t *testing.T) {
	c, err := InitNode([]string{"172.12.17.161:2379"}, "gateway", "2", "222222222222", 15)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	t.Log(c.UpNode())
	select {}
}

func TestInitDiscoveryNode(t *testing.T) {
	c, err := InitDiscoveryNode([]string{"172.12.17.161:2379"}, "gateway")
	if err != nil {
		panic(err)
	}
	t.Log(c.GetNodes())
}
