package network

import "sync"

var connPool = sync.Pool{
	New: func() interface{} {
		return &conn{}
	},
}

func getConn() *conn {
	cli := connPool.Get().(*conn)
	return cli
}

func putConn(cli *conn) {
	connPool.Put(cli)
}
