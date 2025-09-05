import request from '/@/utils/request'
import {SysJobLog} from "/@/views/system/sysJob/list/component/model";
// 查询定时任务列表
export function listSysJob(query:object) {
  return request({
    url: '/api/v1/system/sysJob/list',
    method: 'get',
    params: query
  })
}
// 查询定时任务详细
export function getSysJob(jobId:number) {
  return request({
    url: '/api/v1/system/sysJob/get',
    method: 'get',
    params: {
      jobId: jobId.toString()
    }
  })
}
// 新增定时任务
export function addSysJob(data:object) {
  return request({
    url: '/api/v1/system/sysJob/add',
    method: 'post',
    data: data
  })
}
// 修改定时任务
export function updateSysJob(data:object) {
  return request({
    url: '/api/v1/system/sysJob/edit',
    method: 'put',
    data: data
  })
}
// 删除定时任务
export function delSysJob(jobIds:number[]) {
  return request({
    url: '/api/v1/system/sysJob/delete',
    method: 'delete',
    data:{
      jobIds:jobIds
    }
  })
}
//获取用户信息列表
export function getUserList(uIds:number[]){
    return request({
     url: '/api/v1/system/user/getUsers',
     method: 'get',
     params: {ids:uIds}
   })
}

export function runJob(jobId:number){
  return request({
    url: '/api/v1/system/sysJob/run',
    method: 'put',
    data: {jobId}
  })
}

export function startJob(jobId:number){
  return request({
    url: '/api/v1/system/sysJob/start',
    method: 'put',
    data: {jobId}
  })
}


export function stopJob(jobId:number){
  return request({
    url: '/api/v1/system/sysJob/stop',
    method: 'put',
    data: {jobId}
  })
}


export function getJobLogs(req:object){
  return request({
    url: '/api/v1/system/sysJob/logs',
    method: 'get',
    params: req
  })
}

// 删除定时任务日志
export function delSysJobLog(targetName:string) {
  return request({
    url: '/api/v1/system/sysJob/deleteLogs',
    method: 'delete',
    data:{
      targetName
    }
  })
}
