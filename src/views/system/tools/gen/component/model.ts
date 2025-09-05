// 定义接口来定义对象的类型

export interface TableData {
    tableId:number;
    group:string;
    tableName:string;
    tableComment:string;
    className:string;
    createTime:string;
    updateTime:string;
}

export interface DictOpt{
    dictName :string;
    dictType:string;
}

export interface TableColumns {
    columnId: number;
    tableId: number;
    tableName:string;
    columnName: string;
    columnComment: string;
    columnType: string;
    goType: string;
    tsType:string;
    goField: string;
    htmlField: string;
    isPk: boolean;
    isIncrement: boolean;
    isRequired: boolean;
    isEdit: boolean;
    isList: boolean;
    isDetail: boolean;
    isQuery: boolean;
    sortOrderEdit: number;
    sortOrderList: number;
    sortOrderDetail: number;
    sortOrderQuery: number;
    queryType: string;
    htmlType: string;
    dictType: string;
    linkTableName: string;
    linkTableClass: string;
    linkTableModuleName: string;
    linkTableBusinessName: string;
    linkTablePackage: string;
    linkLabelId: string;
    linkLabelName: string;
    colSpan: number;
    rowSpan: number;
    isRowStart: boolean;
    minWidth: number;
    isFixed: boolean;
    isOverflowTooltip: boolean;
    isCascade: boolean;
    parentColumnName: string;
    cascadeColumnName: string;
}

export interface TableDataInfo extends TableData{
    treeCode:string;
    treeName:string;
    treeParentCode:string;
    functionAuthor:string;
    sortColumn:string;
    sortType:string;
    remark:string;
    overwrite:boolean;
    useSnowId:boolean;
    showDetail:boolean;
    excelPort:boolean;
    excelImp:boolean;
    tplCategory:string;
    packageName:string;
    moduleName:string;
    businessName:string;
    functionName:string;
    linkTableName:string;
    useVirtual:boolean;
    columns:TableColumns[];
    overwriteInfo:Array<any>;
    menuPid:number;
}
export interface TableDataState {
    ids:number[];
    // 非单个禁用
    single: boolean,
    // 非多个禁用
    multiple: boolean,
    // 选中表数组
    tableNames: string[],
    tableData: {
        data: Array<TableData>;
        total: number;
        loading: boolean;
        param: {
            tableName:string;
            tableComment:string;
            pageNum: number;
            pageSize: number;
            dateRange: string[];
        };
    };
}

export interface ImportTableDataState{
    tableData:{
        data:Array<TableData>;
        total:number;
        loading:boolean;
        param:{
            group: string
            pageNum: number;
            pageSize: number;
            tableName: string;
            tableComment: string;
        }
    }
    groups: Array<string>
}

export interface previewState{
    flush: boolean;
    fullscreen: boolean;
    open: boolean;
    title: string;
    data: object;
    activeName: string;
}
