package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 退款分页查询
type RefundInfoGetListReq struct {
	g.Meta `path:"/refund" method:"get" tags:"退款管理" sm:"退款分页列表"`
	Page   uint32 `json:"page" d:"1" v:"min:1" dc:"页码"`
	Size   uint32 `json:"size" d:"10" v:"max:50" dc:"每页数量"`
}

type RefundInfoGetListRes struct {
	List  []*RefundInfoItem `json:"list" dc:"退款列表"`
	Page  uint32            `json:"page" dc:"当前页码"`
	Size  uint32            `json:"size" dc:"每页数量"`
	Total uint32            `json:"total" dc:"总数"`
}

type RefundInfoItem struct {
	Id        uint32                 `json:"id" dc:"退款ID"`
	Number    string                 `json:"number" dc:"退款单号"`
	OrderId   uint32                 `json:"order_id" dc:"订单ID"`
	GoodsId   uint32                 `json:"goods_id" dc:"商品ID"`
	Reason    string                 `json:"reason" dc:"退款原因"`
	Status    uint32                 `json:"status" dc:"状态 1待处理 2同意退款 3拒绝退款"`
	UserId    uint32                 `json:"user_id" dc:"用户ID"`
	CreatedAt *timestamppb.Timestamp `json:"created_at" dc:"创建时间"`
	UpdatedAt *timestamppb.Timestamp `json:"updated_at" dc:"更新时间"`
}

// 退款详情查询
type RefundInfoGetDetailReq struct {
	g.Meta `path:"/refund/detail/{id}" method:"get" tags:"退款管理" sm:"退款详情"`
	Id     uint32 `json:"id" v:"required" dc:"退款ID"`
}

type RefundInfoGetDetailRes struct {
	*RefundInfoItem // 复用列表项结构
}

// 创建退款
type RefundInfoCreateReq struct {
	g.Meta  `path:"/refund" method:"post" tags:"退款管理" sm:"创建退款"`
	UserId  uint32 `json:"user_id" dc:"用户ID"`
	OrderId uint32 `json:"order_id" v:"required" dc:"订单ID"`
	GoodsId uint32 `json:"goods_id" v:"required" dc:"商品ID"`
	Reason  string `json:"reason" v:"required" dc:"退款原因"`
}

type RefundInfoCreateRes struct {
	Id uint32 `json:"id" dc:"退款ID"`
}
