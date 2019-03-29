package job_worker_mode

import "fmt"

type Goods struct {
	Data []byte
}

func (g Goods) UpdateServer() {
	fmt.Println("UpdateServer", string(g.Data))
}
