package perf

import "testing"

func TestStartPprof(t *testing.T) {
	StartPprof([]string{"127.0.0.1:8077"})
	select {}
}
