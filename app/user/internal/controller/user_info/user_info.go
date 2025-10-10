package user_info

import (
	"context"
	"errors"
	"fmt"
	v1 "shop-goframe-micro-service-refacotor/app/user/api/user_info/v1"
	"shop-goframe-micro-service-refacotor/app/user/internal/dao"
	"shop-goframe-micro-service-refacotor/app/user/internal/logic/user_info"
	"shop-goframe-micro-service-refacotor/app/user/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility"
	"shop-goframe-micro-service-refacotor/utility/consts"
	"shop-goframe-micro-service-refacotor/utility/rabbitmq"
	"strings"
	"time"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"

	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

type Controller struct {
	v1.UnimplementedUserInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserInfoServer(s.Server, &Controller{})
}

func (*Controller) Login(ctx context.Context, req *v1.UserInfoLoginReq) (res *v1.UserInfoLoginRes, err error) {
	// 调用logic层
	token, expireIn, userInfo, err := user_info.Login(ctx, req.Name, req.Password)
	// 错误类型
	infoError := consts.InfoError(consts.UserInfo, consts.LoginFail)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 计算过期时间
	expireTime := time.Now().Add(time.Duration(expireIn) * time.Second)
	expireProto := timestamppb.New(expireTime)
	if err := expireProto.CheckValid(); err != nil {
		return nil, err
	}

	// 返回响应
	return &v1.UserInfoLoginRes{
		Type:     "Bearer",
		Token:    token,
		ExpireIn: uint32(expireIn),
		UserInfo: &v1.UserInfoBase{
			Id:     uint32(userInfo.Id),
			Name:   userInfo.Name,
			Avatar: userInfo.Avatar,
			Sex:    uint32(userInfo.Sex),
			Sign:   userInfo.Sign,
			Status: uint32(userInfo.Status),
		},
	}, nil
}

func (c *Controller) Register(ctx context.Context, req *v1.UserInfoRegisterReq) (*v1.UserInfoRegisterRes, error) {
	var registerData *entity.UserInfo
	// 将请求参数req转换为实体对象UserInfo
	if err := gconv.Struct(req, &registerData); err != nil {
		return nil, err
	}
	// 错误类型
	infoError := consts.InfoError(consts.UserInfo, consts.RegisterFail)
	// 调用logic层注册
	userInfo, err := user_info.Register(ctx, registerData)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	go rabbitmq.PublishUserRegisteredEvent(userInfo.Id)

	// 返回响应
	return &v1.UserInfoRegisterRes{
		Id: uint32(userInfo.Id),
	}, nil
}

func (c *Controller) UpdatePassword(ctx context.Context, req *v1.UserInfoUpdatePasswordReq) (*v1.UserInfoUpdatePasswordRes, error) {
	// 调用logic层修改密码
	err := user_info.UpdatePassword(ctx, int(req.Id), req.Password, req.SecretAnswer)
	// 错误类型
	infoError := consts.InfoError(consts.UserInfo, consts.UpdatePasswordFail)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回响应
	return &v1.UserInfoUpdatePasswordRes{
		Id: req.Id,
	}, nil
}

func (*Controller) GetUserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	// 调用logic层获取用户信息
	userInfo, err := user_info.GetUserInfo(ctx, int(req.Id))
	// 错误类型
	infoError := consts.InfoError(consts.UserInfo, consts.GetUserInfoFail)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回响应
	return &v1.UserInfoRes{
		UserInfo: &v1.UserInfoBase{
			Id:     uint32(userInfo.Id),
			Name:   userInfo.Name,
			Avatar: userInfo.Avatar,
			Sex:    uint32(userInfo.Sex),
			Sign:   userInfo.Sign,
			Status: uint32(userInfo.Status),
		},
	}, nil
}

func (*Controller) WxMiniLogin(ctx context.Context, req *v1.WxMiniLoginReq) (res *v1.WxMiniLoginRes, err error) {
	// 发起授权
	miniprogram := wechat.NewWechat().GetMiniProgram(&miniConfig.Config{
		AppID:     g.Cfg().MustGet(nil, "wxMiniConf.appId").String(),
		AppSecret: g.Cfg().MustGet(nil, "wxMiniConf.secret").String(),
		Cache:     cache.NewMemory(),
	})
	authResult, err := miniprogram.GetAuth().Code2Session(req.Code)
	if err != nil || authResult.ErrCode != 0 || authResult.OpenID == "" {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "发起授权请求失败")
	}

	// 解析微信返回数据
	userData, err := miniprogram.GetEncryptor().Decrypt(authResult.SessionKey, req.EncryptedData, req.IV)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "解析数据失败")
	}

	if len(req.Nickname) == 0 {
		nicknameArr := []string{userData.NickName, utility.Krand(6, utility.KC_RAND_KIND_NUM)}
		req.Nickname = strings.Join(nicknameArr, "")
	}
	if len(req.Avatar) == 0 {
		req.Avatar = userData.AvatarURL
	}
	if len(req.Phone) == 0 {
		req.Phone = userData.PhoneNumber
	}

	// 绑定用户或登录
	token, expireIn, userInfo, isNewUser, err := user_info.WxMiniLogin(ctx, authResult.OpenID, req)
	// 错误类型
	infoError := consts.InfoError(consts.UserInfo, consts.WxMiniLoginFail)
	if err != nil {
		// 记录错误日志
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	if isNewUser {
		go rabbitmq.PublishUserRegisteredEvent(userInfo.Id)
	}

	// 计算过期时间
	expireTime := time.Now().Add(time.Duration(expireIn) * time.Second)
	expireProto := timestamppb.New(expireTime)
	if err := expireProto.CheckValid(); err != nil {
		return nil, err
	}

	// 返回响应
	return &v1.WxMiniLoginRes{
		Type:          "Bearer",
		Token:         token,
		ExpireIn:      uint32(expireIn),
		OpenId:        authResult.OpenID,
		NeedPhoneAuth: isNewUser,
		UserInfo: &v1.UserInfoBase{
			Id:     uint32(userInfo.Id),
			Name:   userInfo.Name,
			Avatar: userInfo.Avatar,
			Sex:    uint32(userInfo.Sex),
			Sign:   userInfo.Sign,
			Status: uint32(userInfo.Status),
		},
	}, nil
}

func (*Controller) UpdateInfo(ctx context.Context, req *v1.UserInfoUpdateReq) (res *v1.UserInfoUpdateRes, err error) {
	infoError := consts.InfoError(consts.UserInfo, consts.UpdateFail)
	// 根据ID更新数据库中的信息
	_, err = dao.UserInfo.Ctx(ctx).Where("id", req.Id).Update(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回更新成功响应，包含被更新ID
	return &v1.UserInfoUpdateRes{Id: req.Id}, nil
}

func (*Controller) FillPhone(ctx context.Context, req *v1.FillPhoneReq) (res *v1.FillPhoneRes, err error) {
	miniprogram := wechat.NewWechat().GetMiniProgram(&miniConfig.Config{
		AppID:     g.Cfg().MustGet(nil, "wxMiniConf.appId").String(),
		AppSecret: g.Cfg().MustGet(nil, "wxMiniConf.secret").String(),
		Cache:     cache.NewMemory(),
	})
	authResult, err := miniprogram.GetAuth().Code2Session(req.Code)
	if err != nil || authResult.ErrCode != 0 || authResult.OpenID == "" {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "发起授权请求失败")
	}

	// 解析微信返回数据
	userData, err := miniprogram.GetEncryptor().Decrypt(authResult.SessionKey, req.EncryptedData, req.IV)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "解析数据失败")
	}

	phone := userData.PhoneNumber
	if phone == "" {
		return nil, gerror.WrapCode(gcode.CodeInternalError, errors.New("手机号不能为空"))
	}
	fmt.Println("phone", phone)

	infoError := consts.InfoError(consts.UserInfo, consts.FillPhoneFail)
	// 填充数据库中的手机号
	_, err = dao.UserInfo.Ctx(ctx).Where("id", req.Id).Data(g.Map{
		"phone": phone,
	}).Update()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	// 返回更新成功响应，包含被更新ID
	return &v1.FillPhoneRes{Id: req.Id}, nil
}
