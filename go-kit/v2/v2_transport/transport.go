package v2_transport

import (
	"context"
	"encoding/json"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"learning_tools/go-kit/v2/v2_endpoint"
	"learning_tools/go-kit/v2/v2_service"
	"net/http"
	"strconv"
)

func NewHttpHandler(endpoint v2_endpoint.EndPointServer, log *zap.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			log.Warn(fmt.Sprint(ctx.Value(v2_service.ContextReqUUid), zap.Error(err)))
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
		}), //程序中的全部报错都会走这里面
		//httptransport.ServerErrorHandler(NewZapLogErrorHandler(log)),
		httptransport.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			UUID := uuid.NewV5(uuid.Must(uuid.NewV4()), "req_uuid").String()
			ctx = context.WithValue(ctx, v2_service.ContextReqUUid, UUID)
			return ctx
		}),
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
	var (
		in  v2_service.Add
		err error
	)
	in.A, err = strconv.Atoi(r.FormValue("a"))
	in.B, err = strconv.Atoi(r.FormValue("b"))
	if err != nil {
		return in, err
	}
	return in, nil
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorWrapper struct {
	Error string `json:"errors"`
}
