import request from '/@/utils/request';

// 获取列表
export function getBigFileList (query:Object) {
    return request({
        url: '/api/v1/system/bigFile/list',
        method: 'get',
        params:query
    })
}


// 获取一条数据
export function getBigFile(id:number) {
    return request({
        url: '/api/v1/system/bigFile/get',
        method: 'get',
        params:{id}
    })
}

// 添加文件
export function addBigFile(data:any) {
    return request({
        url: '/api/v1/system/bigFile/add',
        method: 'post',
        data:data
    })
}

// 编辑文件夹
export function editBigFile(data:any) {
    return request({
        url: '/api/v1/system/bigFile/edit',
        method: 'put',
        data:data
    })
}


// 删除文件
export function deleteBigFile(ids:number[]) {
    return request({
        url: '/api/v1/system/bigFile/delete',
        method: 'delete',
        data:{ids}
    })
}

