package perf

import (
	"net/http"
	"net/http/pprof"
	"fmt"
)

// StartPprof start http pprof
func StartPprof(addrs []string) {
	pprofServeMux := http.NewServeMux()
	pprofServeMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	pprofServeMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	pprofServeMux.HandleFunc("/debug/pprof/", pprof.Index)
	pprofServeMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	for _, addr := range addrs {
		go func() {
			if err := http.ListenAndServe(addr, pprofServeMux); err != nil {
				fmt.Printf("http.ListenAndServe(\"%s\", pprofServeMux) error(%v)", addr, err)
				panic(err)
			}
		}()
	}
}
