export interface CouponInfoTableColumns {    
    id:number;  // ID    
    name:string;  // 名称    
    type:number;  // 类型    
    amount:number;  // 优惠金额（元）    
    deadline:string;  // 过期时间    
    createdAt:string;  // 创建时间    
}


export interface CouponInfoInfoData {    
    id:number|undefined;        // ID    
    goodsId:number|undefined; // 关联商品id（0表示全场通用）    
    name:string|undefined; // 名称    
    type:number|undefined; // 类型    
    amount:number|undefined; // 优惠金额（元）    
    deadline:string|undefined; // 过期时间    
    createdAt:string|undefined; // 创建时间    
    updatedAt:string|undefined; // 更新时间    
    deletedAt:string|undefined; // 删除时间（软删除）    
}


export interface CouponInfoTableDataState {
    ids:any[];
    tableData: {
        data: Array<CouponInfoTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            id: number|undefined;            
            name: string|undefined;            
            type: number|undefined;            
            amount: number|undefined;            
            deadline: string|undefined;            
            createdAt: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface CouponInfoEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:CouponInfoInfoData;
    rules: object;
}