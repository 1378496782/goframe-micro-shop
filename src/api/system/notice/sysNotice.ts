import request from '/@/utils/request'
// 查询通知私信列表
export function listSysNotice(query:object) {
  return request({
    url: '/api/v1/system/sysNotice/list',
    method: 'get',
    params: query
  })
}

export function listShowNotice(query:object) {
  return request({
    url: '/api/v1/system/sysNotice/listShow',
    method: 'get',
    params: query
  })
}
// 查询通知私信详细
export function getSysNotice(id:number) {
  return request({
    url: '/api/v1/system/sysNotice/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增通知私信
export function addSysNotice(data:object) {
  return request({
    url: '/api/v1/system/sysNotice/add',
    method: 'post',
    data: data
  })
}
// 修改通知私信
export function updateSysNotice(data:object) {
  return request({
    url: '/api/v1/system/sysNotice/edit',
    method: 'put',
    data: data
  })
}
// 删除通知私信
export function delSysNotice(ids:number[]) {
  return request({
    url: '/api/v1/system/sysNotice/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}


export  function  getIndexData(){
  return request(
      {
        url:'/api/v1/system/sysNotice/getIndexData',
        method:'get'
      }
  )
}

//获取维度通知私信公告信息
export function unReadCount(){
  return request({
    url:'/api/v1/system/sysNotice/unReadCount',
        method:'get'
  })
}

//获取私信待选择的用户
export function getUserList(query:string){
  return request({
    url:'/api/v1/system/sysNotice/userList',
    method:'get',
    params:{userNickname:query}
  })
}
