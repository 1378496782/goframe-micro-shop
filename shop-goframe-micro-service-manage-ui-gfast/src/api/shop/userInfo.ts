import request from '/@/utils/request'
// 查询用户列表
export function listUserInfo(query:object) {
  return request({
    url: '/api/v1/shop/userInfo/list',
    method: 'get',
    params: query
  })
}
// 查询用户详细
export function getUserInfo(id:number) {
  return request({
    url: '/api/v1/shop/userInfo/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增用户
export function addUserInfo(data:object) {
  return request({
    url: '/api/v1/shop/userInfo/add',
    method: 'post',
    data: data
  })
}
// 修改用户
export function updateUserInfo(data:object) {
  return request({
    url: '/api/v1/shop/userInfo/edit',
    method: 'put',
    data: data
  })
}
// 删除用户
export function delUserInfo(ids:number[]) {
  return request({
    url: '/api/v1/shop/userInfo/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
