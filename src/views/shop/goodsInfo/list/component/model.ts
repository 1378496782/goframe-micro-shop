export interface GoodsInfoTableColumns {    
    id:number;  // ID    
    name:string;  // 名称    
    picUrl:string;  // 封面图    
    price:number;  // 价格(分)    
    level1CategoryId:number;  // 一级分类    
    linkedLevel1CategoryId?:LinkedGoodsInfoCategoryInfo; // 一级分类    
    level2CategoryId:number;  // 二级分类    
    linkedLevel2CategoryId?:LinkedGoodsInfoCategoryInfo; // 二级分类    
    level3CategoryId:number;  // 三级分类    
    linkedLevel3CategoryId?:LinkedGoodsInfoCategoryInfo; // 三级分类    
    brand:string;  // 品牌    
    stock:number;  // 库存    
    sale:number;  // 销量    
    tags:string;  // 标签    
    createdAt:string;  //    
    sort:number;  // 排序 倒叙    
    linkedGoodsInfoCategoryInfo:LinkedGoodsInfoCategoryInfo;    
}


export interface GoodsInfoInfoData {    
    id:number|undefined;        // ID    
    name:string|undefined; // 名称    
    images:any[]; // 多图    
    picUrl:string|undefined; // 封面图    
    price:number|undefined; // 价格(分)    
    level1CategoryId:number|undefined; // 一级分类    
    linkedLevel1CategoryId?:LinkedGoodsInfoCategoryInfo; // 一级分类    
    level2CategoryId:number|undefined; // 二级分类    
    linkedLevel2CategoryId?:LinkedGoodsInfoCategoryInfo; // 二级分类    
    level3CategoryId:number|undefined; // 三级分类    
    linkedLevel3CategoryId?:LinkedGoodsInfoCategoryInfo; // 三级分类    
    brand:string|undefined; // 品牌    
    stock:number|undefined; // 库存    
    sale:number|undefined; // 销量    
    tags:string|undefined; // 标签    
    detailInfo:string|undefined; // 商品详情    
    createdAt:string|undefined; //    
    sort:number|undefined; // 排序 倒叙    
    updatedAt:string|undefined; //    
    deletedAt:string|undefined; //    
    linkedGoodsInfoCategoryInfo?:LinkedGoodsInfoCategoryInfo;    
}


export interface LinkedGoodsInfoCategoryInfo {	
    id:number|undefined;    //	
    name:string|undefined;    //	
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
            price: number|undefined;            
            level1CategoryId: number|undefined;            
            level2CategoryId: number|undefined;            
            level3CategoryId: number|undefined;            
            brand: string|undefined;            
            stock: number|undefined;            
            sale: number|undefined;            
            tags: string|undefined;            
            createdAt: string|undefined;            
            sort: number|undefined;            
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