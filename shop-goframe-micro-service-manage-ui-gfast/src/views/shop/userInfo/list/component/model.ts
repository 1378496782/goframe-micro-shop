export interface UserInfoTableColumns {    
    id:number;  //    
    name:string;  // 用户名    
    avatar:string;  // 头像    
    sex:number;  // 1男 2女    
    status:number;  // 1正常 2拉黑冻结    
    sign:string;  // 个性签名    
    createdAt:string;  //    
}


export interface UserInfoInfoData {    
    id:number|undefined;        //    
    name:string|undefined; // 用户名    
    avatar:string|undefined; // 头像    
    password:string|undefined; //    
    userSalt:string|undefined; // 加密盐 生成密码用    
    sex:number|undefined; // 1男 2女    
    status:number|undefined; // 1正常 2拉黑冻结    
    sign:string|undefined; // 个性签名    
    secretAnswer:string|undefined; // 密保问题的答案    
    createdAt:string|undefined; //    
    updatedAt:string|undefined; //    
    deletedAt:string|undefined; //    
}


export interface UserInfoTableDataState {
    ids:any[];
    tableData: {
        data: Array<UserInfoTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            id: number|undefined;            
            name: string|undefined;            
            avatar: string|undefined;            
            sex: number|undefined;            
            status: number|undefined;            
            sign: string|undefined;            
            createdAt: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface UserInfoEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:UserInfoInfoData;
    rules: object;
}