import request from '/@/utils/request'
// 查询已读记录列表
export function listSysNoticeRead(query:object) {
  return request({
    url: '/api/v1/system/sysNoticeRead/list',
    method: 'get',
    params: query
  })
}
// 查询已读记录详细
export function getSysNoticeRead(id:number) {
  return request({
    url: '/api/v1/system/sysNoticeRead/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增已读记录
export function addSysNoticeRead(data:object) {
  return request({
    url: '/api/v1/system/sysNoticeRead/add',
    method: 'post',
    data: data
  })
}
export function readNotice(data:object) {
  return request({
    url: '/api/v1/system/sysNoticeRead/readNotice',
    method: 'post',
    data: data
  })
}
// 修改已读记录
export function updateSysNoticeRead(data:object) {
  return request({
    url: '/api/v1/system/sysNoticeRead/edit',
    method: 'put',
    data: data
  })
}
// 删除已读记录
export function delSysNoticeRead(ids:number[]) {
  return request({
    url: '/api/v1/system/sysNoticeRead/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
