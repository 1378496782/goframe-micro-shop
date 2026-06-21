import axios from 'axios'
import {baseURL, getToken} from "/@/utils/gfast"
import { ElMessage,ElLoading } from 'element-plus';
const mimeMap = {
  xlsx: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
  zip: 'application/zip'
}

const baseUrl = import.meta.env.VITE_API_URL
export function downLoadZip(str:string, filename:string) {
  var url = baseUrl + str
  axios({
    method: 'get',
    url: url,
    responseType: 'blob',
    headers: { 'Authorization': 'Bearer ' + getToken() }
  }).then(res => {
    if(filename){
      res.headers['content-disposition'] = 'attachment; filename='+filename;
    }
    resolveBlob(res, mimeMap.zip)
  })
}

const getHost = ():string=>{
  return window.location.origin
}

// 下载插件
export function downLoadXml (url:string,params:any,rType='get') {
  if (baseUrl=='' ||baseURL=='/'){
    url = getHost() + url
  }else{
    url = baseUrl + url
  }
  const loading = ElLoading.service({
    lock: true,
    text: '正在生成导出数据...',
    spinner: 'el-icon-loading',
    background: 'rgba(0, 0, 0, 0.7)'
  })
  let option = {
    url: url,
    params:params,
    timeout:300000,
    method: 'get',
    responseType: 'blob',
    headers: { 'Authorization': 'Bearer ' + getToken() }
  }
  if(rType=='post'){
    option = {
      url: url,
      data:params,
      timeout:300000,
      method: 'post',
      responseType: 'blob',
      headers: { 'Authorization': 'Bearer ' + getToken() }
    }
  }
  axios(option).then(res => {
    loading.close()
    if (res.status !== 200) {
      ElMessage({
        message: '发生错误，导出失败。',
        type: 'error'
      })
      return
    }
    try {
      const aLink = document.createElement('a')
      var blob = new Blob([res.data], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' })
      var patt = new RegExp('filename=([^;]+\\.[^\\.;]+);*')
      var contentDisposition = decodeURI(res.headers['content-disposition'] || res.headers['Content-Disposition'])
      var result = patt.exec(contentDisposition)
      var fileName = result[1]
      fileName = fileName.replace(/"/g, '')
      aLink.href = URL.createObjectURL(blob)
      aLink.setAttribute('download', fileName) // 设置下载文件名称
      document.body.appendChild(aLink)
      aLink.click()
      document.body.appendChild(aLink)
    } catch (e) {
      console.log(e)
      ElMessage({
        message: '下载失败，可能您登录状态过期，请重新登录后下载。',
        type: 'error'
      })
    }
  }).catch(e=>{
    console.log(e)
    loading.close()
    ElMessage({
      message: '导出失败',
      type: 'error'
    })
  })
}

/**
 * 解析blob响应内容并下载
 * @param {*} res blob响应内容
 * @param {String} mimeType MIME类型
 */
export function resolveBlob(res: any, mimeType:any) {
  const aLink = document.createElement('a')
  var blob = new Blob([res.data], { type: mimeType })
  // //从response的headers中获取filename, 后端response.setHeader("Content-disposition", "attachment; filename=xxxx.docx") 设置的文件名;
  var patt = new RegExp('filename=([^;]+\\.[^\\.;]+);*')
  var contentDisposition = decodeURI(res.headers['content-disposition']||res.headers['Content-Disposition'])
  var result = patt.exec(contentDisposition)
  var fileName = result[1]
  fileName = fileName.replace(/\"/g, '')
  aLink.href = URL.createObjectURL(blob)
  aLink.setAttribute('download', fileName) // 设置下载文件名称
  document.body.appendChild(aLink)
  aLink.click()
  document.body.appendChild(aLink)
}
