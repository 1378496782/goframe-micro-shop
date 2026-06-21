export interface SysJobTableColumns {
    jobId:number;  // 任务ID
    jobName:string;  // 任务名称
    jobGroup:string;  // 任务组名
    invokeTarget:string;  // 任务方法
    cronExpression:string;  // cron执行表达式
    misfirePolicy:number;  // 计划执行策略
    status:number;  // 状态
}


export interface SysJobInfoData {
    jobId:number|undefined;        // 任务ID
    jobName:string|undefined; // 任务名称
    jobParams:string|undefined; // 参数
    jobGroup:string|undefined; // 任务组名
    invokeTarget:string|undefined; // 任务方法
    cronExpression:string|undefined; // cron执行表达式
    misfirePolicy:boolean; // 计划执行策略
    concurrent:number|undefined; // 是否并发执行
    status:boolean; // 状态
    createdBy:number|undefined; // 创建者
    updatedBy:number|undefined; // 更新者
    remark:string|undefined; // 备注信息
    createdAt:string|undefined; // 创建时间
    updatedAt:string|undefined; // 更新时间
    createdUser:{userNickname:string};
    updatedUser:{userNickname:string};
}


export interface SysJobTableDataState {
    jobIds:any[];
    tableData: {
        data: Array<SysJobTableColumns>;
        total: number;
        loading: boolean;
        param: {
            pageNum: number;
            pageSize: number;
            jobName: string|undefined;
            jobGroup: string|undefined;
            status: number|undefined;
        };
    };
}


export interface SysJobEditState{
    loading:boolean;
    isShowDialog: boolean;
    formData:SysJobInfoData;
    rules: object;
}

export interface SysJobLog{
    logIds:Array<number>;
    data:Array<SysJobLogData>;
    total: number;
    loading: boolean;
    param: {
        pageNum: number;
        pageSize: number;
        targetName: string|undefined;
    };
}

export interface SysJobLogData{
    id: number;
    targetName: string;
    createdAt: string;
    result: string
}
