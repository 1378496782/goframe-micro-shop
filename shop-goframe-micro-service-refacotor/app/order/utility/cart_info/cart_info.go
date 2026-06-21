package cart_info

import (
	cartInfo "shop-goframe-micro-service-refacotor/app/goods/api/cart_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

var Client cartInfo.CartInfoClient

func Register() {
	conn := grpcx.Client.MustNewGrpcClientConn("goods", grpcx.Client.ChainUnary(
		middleware.GrpcClientTimeout,
	))
	Client = cartInfo.NewCartInfoClient(conn)
}
