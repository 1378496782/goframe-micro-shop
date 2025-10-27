package collection_info

import (
	"context"
	"database/sql"
	"fmt"
	v1 "shop-goframe-micro-service-refacotor/app/interaction/api/collection_info/v1"
	"shop-goframe-micro-service-refacotor/app/interaction/api/pbentity"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/dao"
	"shop-goframe-micro-service-refacotor/app/interaction/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility"
)

// GetList 列表
func GetList(ctx context.Context, req *v1.CollectionInfoGetListReq) ([]*pbentity.CollectionInfo, int, error) {

	// 查询总数
	total, err := dao.CollectionInfo.Ctx(ctx).
		Where(dao.CollectionInfo.Columns().Type, req.Type).
		Where(dao.CollectionInfo.Columns().UserId, req.UserId).
		Count()
	if err != nil {
		return nil, total, fmt.Errorf("统计收藏列表失败: %v", err)
	}
	if total == 0 {
		return nil, 0, nil
	}

	// 查询当前页数据
	collectionList := make([]*entity.CollectionInfo, total)
	err = dao.CollectionInfo.Ctx(ctx).
		Where(dao.CollectionInfo.Columns().Type, req.Type).
		Where(dao.CollectionInfo.Columns().UserId, req.UserId).
		Page(int(req.Page), int(req.Size)).
		Scan(&collectionList)
	if err != nil {
		return nil, total, fmt.Errorf("查询收藏列表失败: %v", err)
	}

	// 数据转换
	// 在循环中替换手动赋值
	list := make([]*pbentity.CollectionInfo, total)
	for i, v := range collectionList {
		list[i] = &pbentity.CollectionInfo{}
		list[i].Id = int32(v.Id)
		list[i].ObjectId = int32(v.ObjectId)
		list[i].UserId = int32(v.UserId)
		list[i].Type = int32(v.Type)
		list[i].CreatedAt = utility.SafeConvertTime(v.CreatedAt)
		list[i].UpdatedAt = utility.SafeConvertTime(v.UpdatedAt)
	}

	return list, total, nil
}

// Create 创建
func Create(ctx context.Context, req *v1.CollectionInfoCreateReq) (int, error) {

	// 先查询。
	collectionInfo := entity.CollectionInfo{}
	err := dao.CollectionInfo.Ctx(ctx).Where(dao.CollectionInfo.Columns().UserId, req.UserId).Scan(&collectionInfo)
	if err != nil && err != sql.ErrNoRows {
		return 0, fmt.Errorf("查询收藏失败: %v", err)
	}
	// 有数据，直接返回了。
	if collectionInfo.UserId > 0 {
		return 0, nil
	}
	// 没数据进行插入
	// 向数据库中插入数据并获取自动生成的ID
	lastInsertId, err := dao.CollectionInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		return 0, fmt.Errorf("插入收藏失败: %v", err)
	}

	// 返回创建成功响应，包含新创建的ID
	return int(lastInsertId), nil
}

// Delete 删除
func Delete(ctx context.Context, req *v1.CollectionInfoDeleteReq) (int, error) {

	// 先查询。
	collectionInfo := entity.CollectionInfo{}
	err := dao.CollectionInfo.Ctx(ctx).
		Where(dao.CollectionInfo.Columns().Id, req.Id).
		Where(dao.CollectionInfo.Columns().UserId, req.UserId).
		Scan(&collectionInfo)
	if err != nil && err != sql.ErrNoRows {
		return 0, fmt.Errorf("查询收藏记录失败: %v", err)
	}
	// 没有数据，直接返回了。
	if collectionInfo.UserId == 0 {
		return 0, nil
	}

	// 有数据，就删除。
	// 根据ID从数据库中删除对应信息
	_, err = dao.CollectionInfo.Ctx(ctx).
		Where("id", req.Id).
		Where(dao.CollectionInfo.Columns().UserId, req.UserId).
		Delete()
	if err != nil {
		return 0, fmt.Errorf("删除收藏记录失败: %v", err)
	}

	// 返回删除成功的空响应
	return int(req.Id), nil // 返回空结构体
}
