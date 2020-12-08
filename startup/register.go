package startup

import (
	"context"
	"net/http"

	"gitee.com/cristiane/micro-mall-comments/http_server"
	"gitee.com/cristiane/micro-mall-comments/proto/micro_mall_comments_proto/comments_business"
	"gitee.com/cristiane/micro-mall-comments/server"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// RegisterGRPCServer 此处注册pb的Server
func RegisterGRPCServer(grpcServer *grpc.Server) error {
	comments_business.RegisterCommentsBusinessServiceServer(grpcServer, server.NewCommentsServer())
	return nil
}

// RegisterGateway 此处注册pb的Gateway
func RegisterGateway(ctx context.Context, gateway *runtime.ServeMux, endPoint string, dopts []grpc.DialOption) error {
	if err := comments_business.RegisterCommentsBusinessServiceHandlerFromEndpoint(ctx, gateway, endPoint, dopts); err != nil {
		return err
	}
	return nil
}

// RegisterHttpRoute 此处注册http接口
func RegisterHttpRoute(serverMux *http.ServeMux) error {
	serverMux.HandleFunc("/swagger/", http_server.SwaggerHandler)
	return nil
}
