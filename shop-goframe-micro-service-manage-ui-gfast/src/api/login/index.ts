import request from '/@/utils/request';

/**
 * 登录api接口集合
 * @method signIn 用户登录
 */
export function login(params: object){
	return request({
		url: '/api/v1/system/login',
		method: 'post',
		data: params,
	});
}

/**
 * 获取验证码
 */
export function captcha(){
	return request({
		url:"/api/v1/pub/captcha/get",
		method:"get"
	})
}

/**
 * 获取验证码V2
 */
export function captchaV2(){
	return request({
		url:"/api/v1/pub/captcha/v2",
		method:"get"
	})
}

/**
 * 检查v2验证码
 */
export function checkCaptchaV2(data: { dots: string, key: string }){
	return request({
		url:"/api/v1/pub/captcha/v2Check",
		data:data,
		method:"post"
	})
}

/**
 * 退出登录
 */
export function logout(){
	return request({
		url: '/api/v1/system/logout',
		method: 'get',
	});
}
