import request from '/@/utils/request'
import axios, {AxiosProgressEvent, CancelTokenSource} from "axios";
import {baseURL} from "/@/utils/gfast";
// 查询附件管理列表
export function listSysAttachment(query:object) {
  return request({
    url: '/api/v1/system/sysAttachment/list',
    method: 'get',
    params: query
  })
}
// 查询附件管理详细
export function getSysAttachment(id:number) {
  return request({
    url: '/api/v1/system/sysAttachment/get',
    method: 'get',
    params: {
      id: id.toString()
    }
  })
}
// 新增附件管理
export function addSysAttachment(data:object) {
  return request({
    url: '/api/v1/system/sysAttachment/add',
    method: 'post',
    data: data
  })
}
// 修改附件管理
export function updateSysAttachment(data:object) {
  return request({
    url: '/api/v1/system/sysAttachment/edit',
    method: 'put',
    data: data
  })
}
// 删除附件管理
export function delSysAttachment(ids:number[]) {
  return request({
    url: '/api/v1/system/sysAttachment/delete',
    method: 'delete',
    data:{
      ids:ids
    }
  })
}
// 附件管理状态修改
export function changeSysAttachmentStatus(id:number,status:boolean) {
  const data = {
    id,
    status
  }
  return request({
    url: '/api/v1/system/sysAttachment/changeStatus',
    method: 'put',
    data:data
  })
}

// 检查文件分片
export function checkMultipart(data:Object) {
  return request({
    url: '/api/v1/system/upload/checkMultipart',
    method: 'post',
    data:data
  })
}

type Callback = (result: AxiosProgressEvent) => void;
//上传分片
export function uploadPart(data:Object,progress:Callback,cancelToken:CancelTokenSource) {
  /*return axios.post(baseURL+'/api/v1/system/upload/uploadPart', data, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })*/
  return axios({
    url: baseURL+'api/v1/system/upload/uploadPart',
    method: 'post',
    data:data,
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    cancelToken:cancelToken.token,
    onUploadProgress: function (e) {
      progress(e)
    }
  })
}
