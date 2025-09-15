package user_info

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"shop-goframe-micro-service-refacotor/app/user/internal/dao"
	"shop-goframe-micro-service-refacotor/app/user/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility"
	"time"
)

func Login(ctx context.Context, name, password string) (token string, expireIn int, userInfo *entity.UserInfo, err error) {
	// 1. 参数校验
	if name == "" || password == "" {
		return "", 0, nil, errors.New("账号密码不能为空")
	}

	// 2. 查询用户
	userRecord, err := dao.UserInfo.Ctx(ctx).Where("name", name).One()
	if err != nil {
		g.Log().Errorf(ctx, "查询用户失败: %v", err)
		return "", 0, nil, errors.New("系统错误")
	}
	if userRecord.IsEmpty() {
		return "", 0, nil, errors.New("用户不存在")
	}

	// 3. 转换为实体
	var user entity.UserInfo
	if err = userRecord.Struct(&user); err != nil {
		g.Log().Errorf(ctx, "用户数据解析失败: %v", err)
		return "", 0, nil, errors.New("系统错误")
	}

	// 4. 验证密码
	encryptedInput := utility.EncryptPassword(password, user.UserSalt)
	if encryptedInput != user.Password {
		return "", 0, nil, errors.New("密码错误")
	}

	// 5. 生成JWT Token
	token, expireTime, err := utility.GenerateToken(uint32(user.Id))
	if err != nil {
		return "", 0, nil, errors.New("生成token错误")
	}
	expireIn = int(expireTime.Sub(time.Now()).Seconds())
	return token, expireIn, &user, nil
}

func Register(ctx context.Context, req *entity.UserInfo) (*entity.UserInfo, error) {
	// 1. 参数校验
	if req.Name == "" {
		return nil, errors.New("用户名不能为空")
	}
	if len(req.Password) < 6 {
		return nil, errors.New("密码长度至少6位")
	}

	// 2. 检查用户名是否已存在
	count, err := dao.UserInfo.Ctx(ctx).Where("name", req.Name).Count()
	if err != nil {
		return nil, errors.New("检查用户名失败")
	}
	if count > 0 {
		return nil, errors.New("用户名已存在")
	}

	// 3. 生成随机盐值 (10位)
	req.UserSalt = utility.GenerateSalt(10)

	// 4. 使用双重MD5加密密码
	req.Password = utility.EncryptPassword(req.Password, req.UserSalt)

	// 5. 设置默认值
	now := gtime.Now()
	req.Status = 1 // 默认状态正常
	req.CreatedAt = now
	req.UpdatedAt = now

	// 6. 保存到数据库
	id, err := dao.UserInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "创建用户失败: %v", err)
		return nil, errors.New("创建用户失败")
	}

	// 7. 设置ID并返回
	req.Id = int(id)
	return req, nil
}

func UpdatePassword(ctx context.Context, userId int, newPassword, secretAnswer string) error {
	// 1. 查询用户
	userRecord, err := dao.UserInfo.Ctx(ctx).Where("id", userId).One()
	if err != nil {
		g.Log().Errorf(ctx, "查询用户失败: %v", err)
		return errors.New("系统错误")
	}
	if userRecord.IsEmpty() {
		return errors.New("用户不存在")
	}

	// 2. 转换为实体
	var user entity.UserInfo
	if err = userRecord.Struct(&user); err != nil {
		g.Log().Errorf(ctx, "用户数据解析失败: %v", err)
		return errors.New("系统错误")
	}
	// 3. 验证密保答案
	if user.SecretAnswer != secretAnswer {
		return errors.New("密保答案错误")
	}

	// 4. 生成新密码
	newEncryptedPassword := utility.EncryptPassword(newPassword, user.UserSalt)

	// 5. 更新密码
	_, err = dao.UserInfo.Ctx(ctx).Where("id", userId).Update(g.Map{
		"password":   newEncryptedPassword,
		"updated_at": gtime.Now(),
	})
	if err != nil {
		g.Log().Errorf(ctx, "更新密码失败: %v", err)
		return errors.New("系统错误")
	}

	return nil
}

func GetUserInfo(ctx context.Context, userId int) (*entity.UserInfo, error) {
	// 1. 查询用户
	userRecord, err := dao.UserInfo.Ctx(ctx).Where("id", userId).One()
	if err != nil {
		g.Log().Errorf(ctx, "查询用户失败: %v", err)
		return nil, errors.New("系统错误")
	}
	if userRecord.IsEmpty() {
		return nil, errors.New("用户不存在")
	}

	// 2. 转换为实体
	var user entity.UserInfo
	if err = userRecord.Struct(&user); err != nil {
		g.Log().Errorf(ctx, "用户数据解析失败: %v", err)
		return nil, errors.New("系统错误")
	}

	return &user, nil
}

func WxMiniLogin(ctx context.Context, openId, nickName, avatar string) (token string, expireIn int, userInfo *entity.UserInfo, err error) {
	// 1. 参数校验
	if openId == "" {
		return "", 0, nil, errors.New("用登录凭证不能为空")
	}

	// 2. 查询用户
	userRecord, err := dao.UserInfo.Ctx(ctx).Where(" open_id", openId).One()
	if err != nil {
		g.Log().Errorf(ctx, "查询用户失败: %v", err)
		return "", 0, nil, errors.New("系统错误")
	}

	// 3. 注册 or 登录
	now := gtime.Now()
	var userId int
	var user entity.UserInfo
	if userRecord.IsEmpty() {
		// 3-1. 用户不存在，注册
		user = entity.UserInfo{
			Name:      nickName,
			Avatar:    avatar,
			OpenId:    openId,
			CreatedAt: now,
			UpdatedAt: now,
			Status:    1,
		}
		// 3-2. 保存到数据库
		Id, err := dao.UserInfo.Ctx(ctx).InsertAndGetId(&user)
		if err != nil {
			g.Log().Errorf(ctx, "创建用户失败: %v", err)
			return "", 0, nil, errors.New("创建用户失败")
		}
		userId = int(Id)
	} else {
		// 3-1. 用户存在，获取 userId
		if err = userRecord.Struct(&user); err != nil {
			g.Log().Errorf(ctx, "用户数据解析失败: %v", err)
			return "", 0, nil, errors.New("系统错误")
		}
		userId = user.Id
	}

	// 5. 生成JWT Token
	token, expireTime, err := utility.GenerateToken(uint32(userId))
	if err != nil {
		return "", 0, nil, errors.New("生成token错误")
	}
	expireIn = int(expireTime.Sub(time.Now()).Seconds())
	return token, expireIn, &user, nil
}
