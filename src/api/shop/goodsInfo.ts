import request from '/@/utils/request'
// 查询商品列表
export function listGoodsInfo(query:object) {
  return request({
    url: '/api/v1/shop/goodsInfo/list',
    method: 'get',
    params: query
  })
}
// 查询商品详细
export function getGoodsInfo(id:number) {
  return request({
    url: '/api/v1/shop/goodsInfo/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增商品
export function addGoodsInfo(data:object) {
  return request({
    url: '/api/v1/shop/goodsInfo/add',
    method: 'post',
    data: data
  })
}
// 修改商品
export function updateGoodsInfo(data:object) {
  return request({
    url: '/api/v1/shop/goodsInfo/edit',
    method: 'put',
    data: data
  })
}
// 删除商品
export function delGoodsInfo(ids:number[]) {
  return request({
    url: '/api/v1/shop/goodsInfo/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
//相关连表查询数据
export function linkedDataSearch(){
  return request({
    url: '/api/v1/shop/goodsInfo/linkedData',
    method: 'get'
  })
}
