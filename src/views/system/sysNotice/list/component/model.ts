export interface SysNoticeTableColumns {
    id:number;  // ID
    title:string;  // 标题
    type:number;  // 类型
    tag:number;  // 标签
    content:string;  // 内容
    remark:string;  // 备注
    sort:number;  // 排序
    status:number;  // 状态
    createdBy:string;  // 发送人
    createdAt:string;  // 创建时间
}


export interface SysNoticeInfoData {
    id:number|undefined;        // ID
    receiver:number[]|undefined;//用户id
    title:string|undefined; // 标题
    type:number|undefined; // 类型
    tag:number|undefined; // 标签
    content:string|undefined; // 内容
    remark:string|undefined; // 备注
    sort:number|undefined; // 排序
    status:number; // 状态
    createdBy:number|undefined; // 发送人
    updatedBy:number|undefined; // 修改人
    createdAt:string|undefined; // 创建时间
    updatedAt:string|undefined; // 更新时间
    deletedAt:string|undefined; // 删除时间
}


export interface SysNoticeTableDataState {
    ids:any[];
    tableData: {
        data: Array<SysNoticeTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;
            id: number|undefined;
            title: string|undefined;
            type: number|undefined;
            tag: number|undefined;
            status: number|undefined;
            createdAt: string|undefined;
            dateRange: string[];
        };
    };
}


export interface SysNoticeEditState{
    loading:boolean;
    isShowDialog: boolean;
    userListOptions:object[];
    title:string;
    formData:SysNoticeInfoData;
    rules: object;
}
