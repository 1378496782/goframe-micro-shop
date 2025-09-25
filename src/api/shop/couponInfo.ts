import request from '/@/utils/request'
// 查询优惠券列表
export function listCouponInfo(query:object) {
  return request({
    url: '/api/v1/shop/couponInfo/list',
    method: 'get',
    params: query
  })
}
// 查询优惠券详细
export function getCouponInfo(id:number) {
  return request({
    url: '/api/v1/shop/couponInfo/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增优惠券
export function addCouponInfo(data:object) {
  return request({
    url: '/api/v1/shop/couponInfo/add',
    method: 'post',
    data: data
  })
}
// 修改优惠券
export function updateCouponInfo(data:object) {
  return request({
    url: '/api/v1/shop/couponInfo/edit',
    method: 'put',
    data: data
  })
}
// 删除优惠券
export function delCouponInfo(ids:number[]) {
  return request({
    url: '/api/v1/shop/couponInfo/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
