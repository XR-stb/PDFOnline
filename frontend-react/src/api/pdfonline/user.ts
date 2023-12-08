import { PDFOnlineClient } from "./client";
import { UserType } from "../../types";

export const getRegisterCaptcha = (email: string) =>
  PDFOnlineClient({
    url: "users/captcha/register",
    method: "post",
    data: { email },
  })

interface RegisterOpts {
  username: string;
  password: string;
  email: string;
  captcha: string;
}

export const register = (data: RegisterOpts) =>
  PDFOnlineClient<{user: UserType}>({
    url: "users",
    method: "post",
    data,
  }).then((data) => data.user)

interface LoginOpts {
  username: string;
  password: string;
}

export const login = (data: LoginOpts) =>
  PDFOnlineClient<{user: UserType}>({
    url: "users/login",
    method: "post",
    data,
  }).then((data) => data.user)

export const loginGuest = () =>
  PDFOnlineClient<{user: UserType}>({
    url: "users/guest/login",
    method: "post",
  }).then((data) => data.user)

export const logout = () =>
  PDFOnlineClient({
    url: "users/logout",
    method: "post",
  })

export const getUserInfo = (user_id: string) =>
  PDFOnlineClient<{user: UserType}>({
    url: `users/${user_id}`,
    method: "get",
  }).then((data) => data.user)

export const getMe = () =>
  PDFOnlineClient<{user: UserType}>({
    url: `users`,
    method: "get",
    // withCredentials: true,
  }).then((data) => data.user)

interface GetResetPasswordCaptchaOpts {
  username: string;
  email: string;
}

export const getResetPasswordCaptcha = (data: GetResetPasswordCaptchaOpts) =>
  PDFOnlineClient({
    url: "users/captcha/password/reset",
    method: "post",
    data,
  })

interface ResetPasswordOpts {
  username: string;
  captcha: string;
  password: string;
}

export const resetPassword = (data: ResetPasswordOpts) =>
  PDFOnlineClient({
    url: `users/password/reset`,
    method: "put",
    data,
  })