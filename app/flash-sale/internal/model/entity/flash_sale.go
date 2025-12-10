package entity

// FlashSaleGoods 秒杀商品实体
type FlashSaleGoods struct {
	Id            uint32 `json:"id"`             // 主键ID
	GoodsId       uint32 `json:"goods_id"`       // 商品ID
	ActivityId    uint32 `json:"activity_id"`    // 活动ID
	Title         string `json:"title"`          // 商品标题
	Description   string `json:"description"`    // 商品描述
	OriginalPrice int64  `json:"original_price"` // 原价（分）
	SalePrice     int64  `json:"sale_price"`     // 秒杀价（分）
	TotalStock    int32  `json:"total_stock"`    // 总库存
	AvailableStock int32 `json:"available_stock"` // 可用库存
	StartTime     uint32 `json:"start_time"`     // 开始时间
	EndTime       uint32 `json:"end_time"`       // 结束时间
	Status        uint32 `json:"status"`         // 状态：1-进行中，2-已结束，3-未开始
	ImageUrl      string `json:"image_url"`      // 商品图片URL
	CreatedAt     uint32 `json:"created_at"`     // 创建时间
	UpdatedAt     uint32 `json:"updated_at"`     // 更新时间
}

// FlashSaleOrder 秒杀订单实体
type FlashSaleOrder struct {
	Id         uint32 `json:"id"`          // 主键ID
	OrderNo    string `json:"order_no"`    // 订单号
	GoodsId    uint32 `json:"goods_id"`    // 商品ID
	ActivityId uint32 `json:"activity_id"` // 活动ID
	UserId     uint32 `json:"user_id"`     // 用户ID
	Count      int32  `json:"count"`       // 购买数量
	Amount     int64  `json:"amount"`      // 订单金额（分）
	Status     uint32 `json:"status"`      // 状态：0-处理中，1-成功，2-失败
	CreatedAt  uint32 `json:"created_at"`  // 创建时间
	UpdatedAt  uint32 `json:"updated_at"`  // 更新时间
}

// FlashSaleResult 秒杀结果实体
type FlashSaleResult struct {
	Id        uint32 `json:"id"`         // 主键ID
	ResultId  string `json:"result_id"`  // 结果查询ID
	UserId    uint32 `json:"user_id"`    // 用户ID
	GoodsId   uint32 `json:"goods_id"`   // 商品ID
	OrderNo   string `json:"order_no"`   // 订单号
	Status    uint32 `json:"status"`     // 状态：0-处理中，1-成功，2-失败
	Message   string `json:"message"`    // 提示信息
	CreatedAt uint32 `json:"created_at"` // 创建时间
	UpdatedAt uint32 `json:"updated_at"` // 更新时间
}