package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ConsigneeInfoCreateReq 创建收货地址请求
type ConsigneeInfoCreateReq struct {
	g.Meta    `path:"/consignee" method:"post" tags:"收货地址管理" summary:"创建收货地址"`
	IsDefault uint32 `json:"isDefault"   dc:"默认地址1 非默认0"`
	Name      string `json:"name"        dc:"收货人姓名"`
	Phone     string `json:"phone"       dc:"联系电话"`
	Province  string `json:"province"    dc:"省份"`
	City      string `json:"city"        dc:"城市"`
	Town      string `json:"town"        dc:"县区"`
	Street    string `json:"street"      dc:"街道乡镇"`
	Detail    string `json:"detail"      dc:"详细地址"`
}

// ConsigneeInfoCreateRes 创建收货地址响应
type ConsigneeInfoCreateRes struct {
	Id uint32 `json:"id" dc:"收货地址ID"`
}

// ConsigneeInfoGetListReq 获取收货地址列表请求
type ConsigneeInfoGetListReq struct {
	g.Meta `path:"/consignee" method:"get" tags:"收货地址管理" summary:"获取收货地址列表"`
	Page   uint32 `json:"page" v:"min:1" dc:"页码" d:"1"`
	Size   uint32 `json:"size" v:"max:100" dc:"每页数量" d:"10"`
}

// ConsigneeInfoGetListRes 获取收货地址列表响应
type ConsigneeInfoGetListRes struct {
	List  []*ConsigneeInfoItem `json:"list" dc:"收货地址列表"`
	Page  uint32               `json:"page" dc:"当前页码"`
	Size  uint32               `json:"size" dc:"每页数量"`
	Total uint32               `json:"total" dc:"总数"`
}

// ConsigneeInfoItem 收货地址项
type ConsigneeInfoItem struct {
	Id        uint32                 `json:"id"          dc:"收货地址ID"`
	UserId    uint32                 `json:"userId"      dc:"用户ID"`
	IsDefault uint32                 `json:"isDefault"   dc:"默认地址1 非默认0"`
	Name      string                 `json:"name"        dc:"收货人姓名"`
	Phone     string                 `json:"phone"       dc:"联系电话"`
	Province  string                 `json:"province"    dc:"省份"`
	City      string                 `json:"city"        dc:"城市"`
	Town      string                 `json:"town"        dc:"县区"`
	Street    string                 `json:"street"      dc:"街道乡镇"`
	Detail    string                 `json:"detail"      dc:"详细地址"`
	CreatedAt *timestamppb.Timestamp `json:"createdAt"   dc:"创建时间"`
	UpdatedAt *timestamppb.Timestamp `json:"updatedAt"   dc:"更新时间"`
	DeletedAt *timestamppb.Timestamp `json:"deletedAt"   dc:"删除时间"`
}

// ConsigneeInfoUpdateReq 更新收货地址请求
type ConsigneeInfoUpdateReq struct {
	g.Meta    `path:"/consignee" method:"put" tags:"收货地址管理" summary:"更新收货地址"`
	Id        uint32 `json:"id"        v:"required" dc:"收货地址ID"`
	IsDefault uint32 `json:"isDefault" dc:"默认地址1 非默认0"`
	Name      string `json:"name"      dc:"收货人姓名"`
	Phone     string `json:"phone"     dc:"联系电话"`
	Province  string `json:"province"  dc:"省份"`
	City      string `json:"city"      dc:"城市"`
	Town      string `json:"town"      dc:"县区"`
	Street    string `json:"street"    dc:"街道乡镇"`
	Detail    string `json:"detail"    dc:"详细地址"`
}

// ConsigneeInfoUpdateRes 更新收货地址响应
type ConsigneeInfoUpdateRes struct {
	Id uint32 `json:"id" dc:"收货地址ID"`
}

// ConsigneeInfoDeleteReq 删除收货地址请求
type ConsigneeInfoDeleteReq struct {
	g.Meta `path:"/consignee" method:"delete" tags:"收货地址管理" summary:"删除收货地址"`
	Id     uint32 `json:"id" v:"required" dc:"收货地址ID"`
}

// ConsigneeInfoDeleteRes 删除收货地址响应
type ConsigneeInfoDeleteRes struct {
}
