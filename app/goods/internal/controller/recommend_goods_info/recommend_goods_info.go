package recommend_goods_info

import (
	"context"
	"shop-goframe-micro-service-refacotor/app/goods/api/pbentity"
	v1 "shop-goframe-micro-service-refacotor/app/goods/api/recommend_goods_info/v1"
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
	v1.UnimplementedRecommendGoodsInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterRecommendGoodsInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.RecommendGoodsInfoGetListReq) (res *v1.RecommendGoodsInfoGetListRes, err error) {
	response := &v1.RecommendGoodsInfoListResponse{
		List:  make([]*pbentity.GoodsInfo, 0),
		Total: req.Count,
	}
	goodsCount := int(req.Count)

	infoError := consts.InfoError(consts.RecommendGoodsInfo, consts.GetListFail)
	records, err := dao.GoodsInfo.Ctx(ctx).Where("id != ?", req.Id).OrderAsc("price").Limit(goodsCount).All()
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

		var pbRecommendGoodsInfo *pbentity.GoodsInfo
		err = gconv.Struct(addGoodsInfo, &pbRecommendGoodsInfo)
		if err != nil {
			return nil, err
		}
		pbRecommendGoodsInfo.CreatedAt = utility.SafeConvertTime(addGoodsInfo.CreatedAt)
		pbRecommendGoodsInfo.UpdatedAt = utility.SafeConvertTime(addGoodsInfo.UpdatedAt)

		response.List = append(response.List, pbRecommendGoodsInfo)
	}

	return &v1.RecommendGoodsInfoGetListRes{Data: response}, nil
}
