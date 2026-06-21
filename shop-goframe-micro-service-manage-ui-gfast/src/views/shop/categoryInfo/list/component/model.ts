export interface CategoryInfoTableColumns {    
    id:number;  // ID    
    parentId:number;  // 父级id    
    linkedParentId?:LinkedCategoryInfoCategoryInfo; // 父级id    
    name:string;  // 名称    
    picUrl:string;  // 图片    
    level:number;  // 等级    
    sort:number;  // 排序    
    createdAt:string;  // 创建时间    
    linkedCategoryInfoCategoryInfo:LinkedCategoryInfoCategoryInfo;    
}


export interface CategoryInfoInfoData {    
    id:number|undefined;        // ID    
    parentId:number|undefined; // 父级id    
    linkedParentId?:LinkedCategoryInfoCategoryInfo; // 父级id    
    name:string|undefined; // 名称    
    picUrl:string|undefined; // 图片    
    level:number|undefined; // 等级    
    sort:number|undefined; // 排序    
    createdAt:string|undefined; // 创建时间    
    updatedAt:string|undefined; //    
    deletedAt:string|undefined; //    
    linkedCategoryInfoCategoryInfo?:LinkedCategoryInfoCategoryInfo;    
}


export interface LinkedCategoryInfoCategoryInfo {	
    id:number|undefined;    //	
    name:string|undefined;    //	
}


export interface CategoryInfoTableDataState {
    ids:any[];
    tableData: {
        data: Array<CategoryInfoTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            id: number|undefined;            
            parentId: number|undefined;            
            name: string|undefined;            
            picUrl: string|undefined;            
            level: number|undefined;            
            sort: number|undefined;            
            createdAt: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface CategoryInfoEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:CategoryInfoInfoData;
    rules: object;
}