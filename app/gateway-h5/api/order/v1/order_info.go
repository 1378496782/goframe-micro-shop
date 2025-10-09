package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 订单分页查询

// GetList请求  精简参数
type OrderInfoGetListReq struct {
	g.Meta `path:"/order/list" method:"get" tags:"订单管理" sm:"订单分页列表"`
	Page   uint32 `json:"page" d:"1" v:"min:1" dc:"页码"`
	Size   uint32 `json:"size" d:"10" v:"max:50" dc:"每页数量"`
	UserId uint32 `json:"user_id"  v:"required"  dc:"用户ID"`
	Status uint32 `json:"status"  v:"required" dc:"订单状态：1待支付 2已支付待发货 3已发货 4已收货待评价"`
}

type OrderInfoGetListRes struct {
	List  []*OrderInfoItem `json:"list" dc:"订单列表"`
	Page  uint32           `json:"page" dc:"当前页码"`
	Size  uint32           `json:"size" dc:"每页数量"`
	Total uint32           `json:"total" dc:"总数"`
}

// 本orderInfo Item
type OrderInfoItem struct {
	Id          uint32                `json:"id" dc:"订单ID"`
	UserId      uint32                `json:"user_id" dc:"用户ID"`
	Number      string                `json:"number" dc:"订单编号"`
	Price       uint32                `json:"price" dc:"订单金额"`
	ActualPrice uint32                `json:"actual_price" dc:"实际支付金额"`
	Status      uint32                `json:"status" dc:"订单状态"`
	GoodsInfo   []*OrderListGoodsInfo `json:"goods_info" dc:"订单关联的商品信息"`
}

// orderInfo 商品信息专用
type OrderListGoodsInfo struct {
	GoodsId   uint32 `json:"goods_id" dc:"商品ID"`
	Count     uint32 `json:"count" dc:"商品数量"`
	GoodsName string `json:"goods_name" dc:"商品名称"`
	PicUrl    string `json:"pic_url" dc:"商品图片URL"`
}

// 创建订单
type OrderInfoCreateReq struct {
	g.Meta           `path:"/order" method:"post" tags:"订单管理" sm:"创建订单"`
	Price            uint32            `json:"price" v:"required|min:0" dc:"订单金额"`
	CouponPrice      uint32            `json:"coupon_price" d:"0" dc:"优惠券金额"`
	ActualPrice      uint32            `json:"actual_price" v:"required|min:0" dc:"实际支付金额"`
	ConsigneeName    string            `json:"consignee_name"  dc:"收货人姓名"`
	ConsigneePhone   string            `json:"consignee_phone"  dc:"收货人手机号"`
	ConsigneeAddress string            `json:"consignee_address"  dc:"收货人详细地址"`
	Remark           string            `json:"remark" dc:"备注"`
	OrderGoodsInfo   []*OrderGoodsItem `json:"order_goods_info" v:"required" dc:"订单商品信息"`
	CouponId         uint32            `json:"coupon_id" dc:"优惠券ID"`
}

type OrderInfoCreateRes struct {
	Id     uint32 `json:"id" dc:"订单ID"`
	Number string `json:"number" dc:"订单编号"`
}

type OrderGoodsItem struct {
	GoodsId        uint32 `json:"goods_id" v:"required" dc:"商品ID"`
	GoodsOptionsId uint32 `json:"goods_options_id" dc:"商品规格ID"`
	Count          uint32 `json:"count" v:"required|min:1" dc:"商品数量"`
	Remark         string `json:"remark" dc:"备注"`
	Price          uint32 `json:"price" v:"required|min:0" dc:"商品金额"`
	CouponPrice    uint32 `json:"coupon_price" d:"0" dc:"商品优惠券金额"`
	ActualPrice    uint32 `json:"actual_price" v:"required|min:0" dc:"商品实际支付金额"`
}

// PaymentReq 支付请求体
type PaymentReq struct {
	g.Meta `path:"/payment" method:"post" tags:"订单管理" sm:"发起支付"`
	OpenId string `json:"openId" v:"required" dc:"用户登录凭证"`
	Amount int64  `json:"amount" v:"required" dc:"金额，单位为分"`
	Number string `json:"number" v:"required" dc:"订单编号"`
}

// PaymentRes 支付响应体
type PaymentRes struct {
	TimeStamp  string `json:"timeStamp" dc:"时间戳，单位秒，字符串格式"`
	NonceStr   string `json:"nonceStr" dc:"随机字符串，防重放攻击"`
	Package    string `json:"package" dc:"统一下单接口返回的预支付交易会话标识，格式为 prepay_id=***"`
	SignType   string `json:"signType" dc:"签名类型，微信支付 v3 常用 RSA"`
	PaySign    string `json:"paySign" dc:"支付签名，商户后端使用私钥生成"`
	OutTradeNo string `json:"out_trade_no" dc:"商户订单号，用于后续查单、退款"`
}

// NotifyReq 回调请求体
type NotifyReq struct {
	g.Meta `path:"/notify" method:"post" tags:"订单管理" sm:"支付回调"`
	// 注意：这些字段不由框架自动绑定（因为微信是 POST 原始 JSON），需要手动读取 body 并赋值
	RawBody string            `json:"-" dc:"回调原始body（由框架手动读取）"`
	Headers map[string]string `json:"-" dc:"回调请求头（由框架手动读取）"`
}

// NotifyRes 回调响应体
type NotifyRes struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// 获取订单详情
type OrderInfoGetDetailReq struct {
	g.Meta `path:"/order/{id}" method:"get" tags:"订单管理" sm:"获取订单详情"`
	Id     uint32 `json:"id" v:"min:1" dc:"订单ID"`
}

// OrderInfoDetail 订单详情主信息
type OrderInfoDetail struct {
	Id               uint32      `json:"id" dc:"订单ID"`
	UserId           uint32      `json:"user_id" dc:"用户ID"`
	Number           string      `json:"number" dc:"订单编号"`
	Price            uint32      `json:"price" dc:"订单金额"`
	CouponPrice      uint32      `json:"coupon_price" dc:"优惠券金额"`
	ActualPrice      uint32      `json:"actual_price" dc:"实际支付金额"`
	PayType          uint32      `json:"pay_type" dc:"支付方式"`
	Remark           string      `json:"remark" dc:"备注"`
	Status           uint32      `json:"status" dc:"订单状态"`
	ConsigneeName    string      `json:"consignee_name" dc:"收货人姓名"`
	ConsigneePhone   string      `json:"consignee_phone" dc:"收货人手机号"`
	ConsigneeAddress string      `json:"consignee_address" dc:"收货人地址"`
	CreatedAt        *gtime.Time `json:"created_at" dc:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updated_at" dc:"更新时间"`
	PayAt            *gtime.Time `json:"pay_at" dc:"支付时间"`
}

type OrderInfoGetDetailRes struct {
	OrderInfo       *OrderInfoDetail  `json:"order_info" dc:"订单详情"`
	OrderGoodsInfos []*OrderGoodsItem `json:"order_goods_infos" dc:"订单商品列表"`
}

// OrderInfoGetCountReq 获取订单数量请求
type OrderInfoGetCountReq struct {
	g.Meta `path:"/order/count" method:"get" tags:"订单管理" sm:"获取订单数量"`
}

// OrderInfoGetCountRes 获取订单数量响应
type OrderInfoGetCountRes struct {
	Pending   uint32 `json:"pending"`
	Shipping  uint32 `json:"shipping"`
	Delivered uint32 `json:"delivered"`
	Completed uint32 `json:"completed"`
	AfterSale uint32 `json:"afterSale"`
}
