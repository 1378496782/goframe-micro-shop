export interface RotationInfoTableColumns {    
    id:number;  // ID    
    picUrl:string;  // 图片    
    link:string;  // 跳转链接    
    sort:number;  // 排序字段    
    createdAt:string;  // 创建时间    
}


export interface RotationInfoInfoData {    
    id:number|undefined;        // ID    
    picUrl:string|undefined; // 图片    
    link:string|undefined; // 跳转链接    
    sort:number|undefined; // 排序字段    
    createdAt:string|undefined; // 创建时间    
    updatedAt:string|undefined; //    
    deletedAt:string|undefined; //    
}


export interface RotationInfoTableDataState {
    ids:any[];
    tableData: {
        data: Array<RotationInfoTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            id: number|undefined;            
            picUrl: string|undefined;            
            link: string|undefined;            
            sort: number|undefined;            
            createdAt: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface RotationInfoEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:RotationInfoInfoData;
    rules: object;
}