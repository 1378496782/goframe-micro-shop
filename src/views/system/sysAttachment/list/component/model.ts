export interface SysAttachmentTableColumns {
    id:number;  // 文件ID
    appId:string;  // 应用ID
    drive:string;  // 上传驱动
    name:string;  // 文件原始名
    kind:string;  // 上传类型
    path:string;  // 本地路径
    size:number;  // 文件大小
    ext:string;  // 扩展名
    status:boolean;  // 状态
    createdAt:string;  // 创建时间
}


export interface SysAttachmentInfoData {
    id:number|undefined;        // 文件ID
    appId:string|undefined; // 应用ID
    drive:string|undefined; // 上传驱动
    name:string|undefined; // 文件原始名
    kind:string|undefined; // 上传类型
    mimeType:string|undefined; // 扩展类型
    path:string|undefined; // 本地路径
    size:number|undefined; // 文件大小
    ext:string|undefined; // 扩展名
    md5:string|undefined; // md5校验码
    createdBy:number|undefined; // 上传人ID
    status:boolean; // 状态
    createdAt:string|undefined; // 创建时间
    updatedAt:string|undefined; // 修改时间
}


export interface SysAttachmentTableDataState {
    ids:any[];
    tableData: {
        data: Array<SysAttachmentTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;
            appId: string|undefined;
            drive: string|undefined;
            name: string|undefined;
            kind: string|undefined;
            mimeType: string|undefined;
            status: boolean|undefined;
            createdAt: string[];
            dateRange: string[];
        };
    };
}


export interface SysAttachmentEditState{
    isShowDialog: boolean;
    images:any[];
    files:any[];
}
