import { PDFOnlineClient } from "./client";

export const getCaptcha = (email: string) =>
  PDFOnlineClient({
    url: "users/captcha",
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
  PDFOnlineClient<{user_id:string}>({
    url: "users/register",
    method: "post",
    data,
  }).then((data) => data.user_id)

interface LoginOpts {
  username: string;
  password: string;
}

export const login = (data: LoginOpts) =>
  PDFOnlineClient<{user_id:string}>({
    url: "users/login",
    method: "post",
    data,
  }).then((data) => data.user_id)

export const logout = () =>
  PDFOnlineClient({
    url: "users/logout",
    method: "post",
  })

export interface User {
  id: string;
  username: string;
  email: string;
}

export const getUserInfo = (user_id: string) =>
  PDFOnlineClient<{user: User}>({
    url: `users/${user_id}`,
    method: "get",
  }).then((data) => data.user)