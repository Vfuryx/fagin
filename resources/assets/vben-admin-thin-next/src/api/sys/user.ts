import { defHttp } from '/@/utils/http/axios';
import {
  LoginParams,
  LoginResultModel,
  GetUserInfoModel,
  GetCaptchaModel,
} from './model/userModel';

import { ErrorMessageMode } from '/#/axios';

enum Api {
  Login = '/login',
  Logout = '/common/auth/logout',
  GetUserInfo = '/common/auth/info',
  GetPermCode = '/common/auth/permcode',
  GetCaptcha = '/captcha',
  // Login = '/login',
  // Logout = '/logout',
  // GetUserInfo = '/getUserInfo',
  // GetPermCode = '/getPermCode',
}

/**
 * @description: user login api
 */
export function loginApi(params: LoginParams, mode: ErrorMessageMode = 'modal') {
  return defHttp.post<LoginResultModel>(
    {
      url: Api.Login,
      params,
    },
    {
      errorMessageMode: mode,
    },
  );
}

/**
 * @description: getUserInfo
 */
export function getUserInfo() {
  return defHttp.get<GetUserInfoModel>({ url: Api.GetUserInfo }, { errorMessageMode: 'none' });
}

/**
 * @description: 获取权限码
 */
export function getPermCode() {
  return defHttp.get<string[]>({ url: Api.GetPermCode });
}

/**
 * @description: 退出
 */
export function doLogout() {
  return defHttp.post({ url: Api.Logout });
}

/**
 * 获取验证码
 */
export function getCaptcha() {
  return defHttp.post<GetCaptchaModel>({ url: Api.GetCaptcha });
}
