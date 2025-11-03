package goods

import (
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	goodsInfo "shop-goframe-micro-service-refacotor/app/goods/api/goods_info/v1"
	"shop-goframe-micro-service-refacotor/utility/middleware"
)

var Client goodsInfo.GoodsInfoClient

func Register() {
	conn := grpcx.Client.MustNewGrpcClientConn("goods", grpcx.Client.ChainUnary(
		middleware.GrpcClientTimeout,
	))
	Client = goodsInfo.NewGoodsInfoClient(conn)
}
