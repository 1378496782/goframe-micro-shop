import request from '/@/utils/request'
// 查询轮播图列表
export function listRotationInfo(query:object) {
  return request({
    url: '/api/v1/shop/rotationInfo/list',
    method: 'get',
    params: query
  })
}
// 查询轮播图详细
export function getRotationInfo(id:number) {
  return request({
    url: '/api/v1/shop/rotationInfo/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增轮播图
export function addRotationInfo(data:object) {
  return request({
    url: '/api/v1/shop/rotationInfo/add',
    method: 'post',
    data: data
  })
}
// 修改轮播图
export function updateRotationInfo(data:object) {
  return request({
    url: '/api/v1/shop/rotationInfo/edit',
    method: 'put',
    data: data
  })
}
// 删除轮播图
export function delRotationInfo(ids:number[]) {
  return request({
    url: '/api/v1/shop/rotationInfo/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
