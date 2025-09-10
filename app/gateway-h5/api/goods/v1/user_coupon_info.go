package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserCouponInfoGetListReq 优惠券选项分页查询
type UserCouponInfoGetListReq struct {
	g.Meta `path:"/user_coupon" method:"get" tags:"优惠券管理" sm:"优惠券选项列表"`
	Page   uint32 `json:"page" d:"1"  v:"min:1" dc:"页码"`
	Size   uint32 `json:"size" d:"10" v:"max:100" dc:"每页数量"`
}

type UserCouponInfoGetListRes struct {
	List  []*UserCouponInfoItem `json:"list" dc:"优惠券选项列表"`
	Page  uint32                `json:"page" dc:"当前页码"`
	Size  uint32                `json:"size" dc:"每页数量"`
	Total uint32                `json:"total" dc:"总数"`
}

type UserCouponInfoItem struct {
	Id        uint32                 `json:"id" dc:"优惠券ID"`
	UserId    int32                  `json:"user_id" dc:"用户ID"`
	GoodsId   uint32                 `json:"goods_id" dc:"商品id"`
	Count     uint32                 `json:"count" dc:"商品数量"`
	CreatedAt *timestamppb.Timestamp `json:"created_at" dc:"创建时间"`
	UpdatedAt *timestamppb.Timestamp `json:"updated_at" dc:"更新时间"`
}
