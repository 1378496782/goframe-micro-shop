package admin_info

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"shop-goframe-micro-service-refacotor/utility"

	"shop-goframe-micro-service-refacotor/app/admin/internal/dao"
	"shop-goframe-micro-service-refacotor/app/admin/internal/model/entity"
	"time"
)

func Login(ctx context.Context, name, password string) (token string, expire time.Time, err error) {
	// 1. 参数校验
	if name == "" || password == "" {
		return "", time.Time{}, errors.New("账号密码不能为空")
	}

	// 2. 查询用户
	adminRecord, err := dao.AdminInfo.Ctx(ctx).Where("name", name).One()
	if err != nil {
		g.Log().Errorf(ctx, "查询用户失败: %v", err)
		return "", time.Time{}, errors.New("系统错误")
	}
	if adminRecord.IsEmpty() {
		return "", time.Time{}, errors.New("用户不存在")
	}

	// 3. 转换为实体
	var admin entity.AdminInfo
	if err = adminRecord.Struct(&admin); err != nil {
		g.Log().Errorf(ctx, "用户数据解析失败: %v", err)
		return "", time.Time{}, errors.New("系统错误")
	}

	// 4. 验证密码
	encryptedInput := utility.EncryptPassword(password, admin.UserSalt)
	if encryptedInput != admin.Password {
		return "", time.Time{}, errors.New("密码错误")
	}

	// 5. 生成JWT Token
	return utility.GenerateToken(uint32(admin.Id))
}

// Register 管理员注册
func Register(ctx context.Context, name, password string) (*entity.AdminInfo, error) {
	// 1. 参数校验
	if name == "" {
		return nil, errors.New("用户名不能为空")
	}
	if len(password) < 6 {
		return nil, errors.New("密码长度至少6位")
	}

	// 2. 检查用户名是否已存在
	count, err := dao.AdminInfo.Ctx(ctx).Where("name", name).Count()
	if err != nil {
		return nil, errors.New("检查用户名失败")
	}
	if count > 0 {
		return nil, errors.New("用户名已存在")
	}

	// 3. 生成随机盐值 (10位)
	salt := utility.GenerateSalt(10)

	// 4. 使用双重MD5加密密码
	encryptedPassword := utility.EncryptPassword(password, salt)

	// 5. 创建管理员记录
	now := gtime.Now()
	admin := &entity.AdminInfo{
		Name:      name,
		Password:  encryptedPassword,
		UserSalt:  salt,
		RoleIds:   "2", // 默认角色ID
		IsAdmin:   0,   // 默认超级管理员
		CreatedAt: now,
		UpdatedAt: now,
	}

	// 6. 保存到数据库
	id, err := dao.AdminInfo.Ctx(ctx).InsertAndGetId(admin)
	if err != nil {
		g.Log().Errorf(ctx, "创建用户失败: %v", err)
		return nil, errors.New("创建用户失败")
	}

	// 7. 设置ID并返回
	admin.Id = int(id)
	return admin, nil
}
