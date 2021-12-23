package main

import (
	"testing"

	"github.com/hwholiday/learning_tools/grpc/simple_rpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func Test(t *testing.T) {

}

//BenchmarkA-16    	    4872	    279341 ns/op
func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := grpc.Dial("127.0.0.1:8099", grpc.WithInsecure())
		if err != nil {
			b.Error(err)
		}
		c := proto.NewHowieClient(conn)
		_, err = c.LoL(context.Background(), &proto.HowieUp{Name: "howie"})
		if err != nil {
			b.Error(err)
		}

	}
}

//BenchmarkB-16    	   21667	     56072 ns/op
func BenchmarkB(b *testing.B) {
	conn, err := grpc.Dial("127.0.0.1:8099", grpc.WithInsecure())
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		c := proto.NewHowieClient(conn)
		_, err = c.LoL(context.Background(), &proto.HowieUp{Name: "howie"})
		if err != nil {
			b.Error(err)
		}
	}
}
