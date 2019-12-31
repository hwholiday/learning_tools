package v1_transport

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"learning_tools/go-kit/v1/v1_endpoint"
	"net/http"
	"strconv"
)

func NewHttpHandler(endpoint v1_endpoint.EndPointServer) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder), //程序中的全部报错都会走这里面
		/*httptransport.ServerErrorEncoder(transport.NewLogErrorHandler(logger))*/
	}

	m := http.NewServeMux()
	m.Handle("/sum", httptransport.NewServer(
		endpoint.AddEndPoint,
		decodeHTTPADDRequest,      //解析请求值
		encodeHTTPGenericResponse, //返回值
		options...,
	))
	return m
}
func decodeHTTPADDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	d, err := strconv.Atoi(r.FormValue("a"))
	if err != nil {
		return d, err
	}
	return d, nil
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println("encodeHTTPGenericResponse", response)
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	fmt.Println("errorEncoder", err.Error())
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

type errorWrapper struct {
	Error string `json:"errors"`
}
