import request from '/@/utils/request';


// 根据ids 获取用户列表信息
export function getUserByIds(query:Object) {
    return request({
        url: '/api/v1/system/user/getUserByIds',
        method: 'get',
        params:query
    })
}


export function getUserList(query:Object) {
    return request({
        url: '/api/v1/system/user/list',
        method: 'get',
        params:query
    })
}

export function getUsersById(roleId:number) {
    return request({
        url: '/api/v1/system/user/getByRoleId',
        method: 'get',
        params:{roleId}
    })
}


export function getUserSelector(query:Object) {
    return request({
        url: '/api/v1/system/user/selector',
        method: 'get',
        params:query
    })
}

export function getDeptTree(showOwner?:boolean) {
    return request({
        url: '/api/v1/system/dept/treeSelect',
        method: 'get',
        params:{showOwner}
    })
}

export function getParams() {
    return request({
        url: '/api/v1/system/user/params',
        method: 'get'
    })
}

export function getEditUser(id:number) {
    return request({
        url: '/api/v1/system/user/getEdit',
        method: 'get',
        params:{id}
    })
}

export function addUser(data:object) {
    return request({
        url: '/api/v1/system/user/add',
        method: 'post',
        data:data
    })
}


export function editUser(data:object) {
    return request({
        url: '/api/v1/system/user/edit',
        method: 'put',
        data:data
    })
}

export function resetUserPwd(userId:number, password:string) {
    return request({
        url: '/api/v1/system/user/resetPwd',
        method: 'put',
        data:{userId,password}
    })
}

export function changeUserStatus(userId:number, status:number) {
    return request({
        url: '/api/v1/system/user/setStatus',
        method: 'put',
        data:{userId,status}
    })
}


export function deleteUser(ids:number[]) {
    return request({
        url: '/api/v1/system/user/delete',
        method: 'delete',
        data:{ids}
    })
}


