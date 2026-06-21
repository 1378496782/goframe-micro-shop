import request from '/@/utils/request'
// 查询商品分类列表
export function listCategoryInfo(query:object) {
  return request({
    url: '/api/v1/demo/categoryInfo/list',
    method: 'get',
    params: query
  })
}
// 查询商品分类详细
export function getCategoryInfo(id:number) {
  return request({
    url: '/api/v1/demo/categoryInfo/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增商品分类
export function addCategoryInfo(data:object) {
  return request({
    url: '/api/v1/demo/categoryInfo/add',
    method: 'post',
    data: data
  })
}
// 修改商品分类
export function updateCategoryInfo(data:object) {
  return request({
    url: '/api/v1/demo/categoryInfo/edit',
    method: 'put',
    data: data
  })
}
// 删除商品分类
export function delCategoryInfo(ids:number[]) {
  return request({
    url: '/api/v1/demo/categoryInfo/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
