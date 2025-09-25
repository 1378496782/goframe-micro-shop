export interface GoodsInfoTableColumns {    
    id:number;  // ID    
    name:string;  // 名字    
    picUrl:string;  // 主图    
    images:any[];  // 详情配图    
    price:number;  // 价格(分)    
    stock:number;  // 库存    
    sale:number;  // 销量    
    tags:string;  // 标签    
    sort:number;  // 排序 倒叙    
    enableBargain:number;  // 允许砍价    
    createdAt:string;  //    
}


export interface GoodsInfoInfoData {    
    id:number|undefined;        // ID    
    name:string|undefined; // 名字    
    picUrl:string|undefined; // 主图    
    images:any[]; // 详情配图    
    price:number|undefined; // 价格(分)    
    level1CategoryId:number|undefined; // 1级分类id    
    level2CategoryId:number|undefined; // 2级分类id    
    level3CategoryId:number|undefined; // 3级分类id    
    brand:string|undefined; // 品牌    
    stock:number|undefined; // 库存    
    sale:number|undefined; // 销量    
    tags:string|undefined; // 标签    
    sort:number|undefined; // 排序 倒叙    
    detailInfo:string|undefined; // 商品详情    
    enableBargain:number|undefined; // 允许砍价    
    createdAt:string|undefined; //    
    updatedAt:string|undefined; //    
    deletedAt:string|undefined; //    
}


export interface GoodsInfoTableDataState {
    ids:any[];
    tableData: {
        data: Array<GoodsInfoTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;            
            id: number|undefined;            
            name: string|undefined;            
            picUrl: string|undefined;            
            images: string|undefined;            
            price: number|undefined;            
            stock: number|undefined;            
            sale: number|undefined;            
            tags: string|undefined;            
            sort: number|undefined;            
            enableBargain: number|undefined;            
            createdAt: string|undefined;            
            dateRange: string[];
        };
    };
}


export interface GoodsInfoEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:GoodsInfoInfoData;
    rules: object;
}