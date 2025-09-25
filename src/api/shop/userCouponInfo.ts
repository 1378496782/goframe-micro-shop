import request from '/@/utils/request'
// 查询用户优惠券列表
export function listUserCouponInfo(query:object) {
  return request({
    url: '/api/v1/shop/userCouponInfo/list',
    method: 'get',
    params: query
  })
}
// 查询用户优惠券详细
export function getUserCouponInfo(id:number) {
  return request({
    url: '/api/v1/shop/userCouponInfo/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增用户优惠券
export function addUserCouponInfo(data:object) {
  return request({
    url: '/api/v1/shop/userCouponInfo/add',
    method: 'post',
    data: data
  })
}
// 修改用户优惠券
export function updateUserCouponInfo(data:object) {
  return request({
    url: '/api/v1/shop/userCouponInfo/edit',
    method: 'put',
    data: data
  })
}
// 删除用户优惠券
export function delUserCouponInfo(ids:number[]) {
  return request({
    url: '/api/v1/shop/userCouponInfo/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
//相关连表查询数据
export function linkedDataSearch(){
  return request({
    url: '/api/v1/shop/userCouponInfo/linkedData',
    method: 'get'
  })
}
