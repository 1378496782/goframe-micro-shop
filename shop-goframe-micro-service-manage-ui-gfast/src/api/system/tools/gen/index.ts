import request from '/@/utils/request';

export function getDBList() {
    return request({
        url: '/api/v1/system/tools/gen/dbList',
        method: 'get',
    })
}

export function getTableList(query:Object) {
    return request({
        url: '/api/v1/system/tools/gen/tableList',
        method: 'get',
        params:query
    })
}

export function getImportTableList(query:Object) {
    return request({
        url: '/api/v1/system/tools/gen/dataList',
        method: 'get',
        params:query
    })
}


export function importTable(tables:string[], group:string) {
    return request({
        url: '/api/v1/system/tools/gen/importTableSave',
        method: 'post',
        data:{tables:tables, group:group}
    })
}

export function deleteTables(ids:number[]) {
    return request({
        url: '/api/v1/system/tools/gen/tableDelete',
        method: 'delete',
        data:{ids}
    })
}

export function getGenTable(tableId:any){
    return request({
        url: '/api/v1/system/tools/gen/columnList',
        method: 'get',
        params:{tableId}
    })
}

export function getRelationTable(query:Object){
    return request({
        url: '/api/v1/system/tools/gen/relationTable',
        method: 'get',
        params:query
    })
}


export function saveEdit(data:any) {
    return request({
        url: '/api/v1/system/tools/gen/editSave',
        method: 'post',
        data:data
    })
}

export function showPreview(tableId:number){
    return request({
        url: '/api/v1/system/tools/gen/preview',
        method: 'get',
        params: {tableId}
    })
}

export function batchGenCode(ids:number[]){
    return request({
        url: '/api/v1/system/tools/gen/batchGenCode',
        method: 'post',
        data: {ids}
    })
}


export function syncTable(tableId:number){
    return request({
        url: '/api/v1/system/tools/gen/syncTable',
        method: 'post',
        data: {tableId}
    })
}
