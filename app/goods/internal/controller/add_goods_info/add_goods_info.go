package add_goods_info

import (
	"context"
	v1 "shop-goframe-micro-service-refacotor/app/goods/api/add_goods_info/v1"
	"shop-goframe-micro-service-refacotor/app/goods/api/pbentity"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/app/goods/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility"
	"shop-goframe-micro-service-refacotor/utility/consts"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type Controller struct {
	v1.UnimplementedAddGoodsInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAddGoodsInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.AddGoodsInfoGetListReq) (res *v1.AddGoodsInfoGetListRes, err error) {
	response := &v1.AddGoodsInfoListResponse{
		List:  make([]*pbentity.GoodsInfo, 0),
		Total: 5,
	}

	infoError := consts.InfoError(consts.AddGoodsInfo, consts.GetListFail)
	records, err := dao.GoodsInfo.Ctx(ctx).OrderAsc("price").Limit(5).All()
	if err != nil {
		g.Log().Errorf(ctx, infoError, err)
		return res, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	for _, record := range records {
		var addGoodsInfo *entity.GoodsInfo
		err = gconv.Struct(record, &addGoodsInfo)
		if err != nil {
			return nil, err
		}

		var pbAddGoodsInfo *pbentity.GoodsInfo
		err = gconv.Struct(addGoodsInfo, &pbAddGoodsInfo)
		if err != nil {
			return nil, err
		}
		pbAddGoodsInfo.CreatedAt = utility.SafeConvertTime(addGoodsInfo.CreatedAt)
		pbAddGoodsInfo.UpdatedAt = utility.SafeConvertTime(addGoodsInfo.UpdatedAt)

		response.List = append(response.List, pbAddGoodsInfo)
	}

	return &v1.AddGoodsInfoGetListRes{Data: response}, nil
}
