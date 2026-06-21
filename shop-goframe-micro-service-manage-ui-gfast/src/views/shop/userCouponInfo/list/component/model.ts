export interface UserCouponInfoTableColumns {    
    id:number;  //    
    userId:number;  // 用户id    
    couponId:number;  // 优惠券id    
    linkedCouponId?:LinkedUserCouponInfoCouponInfo; // 优惠券id    
    status:number;  // 状态    
    amount:number;  // 优惠金额（元）    
    createdAt:string;  // 创建时间    
    linkedUserCouponInfoCouponInfo:LinkedUserCouponInfoCouponInfo;    
}


export interface UserCouponInfoInfoData {    
    id:number|undefined;        //    
    userId:number|undefined; // 用户id    
    couponId:number|undefined; // 优惠券id    
    linkedCouponId?:LinkedUserCouponInfoCouponInfo; // 优惠券id    
    status:number|undefined; // 状态    
    amount:number|undefined; // 优惠金额（元）    
    createdAt:string|undefined; // 创建时间    
    updatedAt:string|undefined; // 更新时间    
    deletedAt:string|undefined; // 删除时间（软删除）    
    linkedUserCouponInfoCouponInfo?:LinkedUserCouponInfoCouponInfo;    
}


export interface LinkedUserCouponInfoCouponInfo {	
    id:number|undefined;    //	
    name:string|undefined;    // 优惠券名称	
}


export interface UserCouponInfoTableDataState {
    ids:any[];
    tableData: {
        data: Array<UserCouponInfoTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            id: number|undefined;            
            userId: number|undefined;            
            couponId: number|undefined;            
            status: number|undefined;            
            amount: number|undefined;            
            createdAt: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface UserCouponInfoEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:UserCouponInfoInfoData;
    rules: object;
}