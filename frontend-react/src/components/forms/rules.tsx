import { FormRule } from "antd";

export const usernameRules: FormRule[] = [
  {required: true, message: 'Please input your username!', whitespace: true},
  {max: 32, message: 'Must be less than 32 characters!'},
  {min: 2, message: 'Must be more than 2 characters!'}
]

export const passwordRules: FormRule[] = [
  {required: true, message: 'Please input your password!'},
  {max: 32, message: 'Must be less than 32 characters!'},
  {min: 6, message: 'Must be more than 6 characters!'}
]

export const confirmPasswordRules: FormRule[] = [
  {required: true, message: 'Please confirm your password!'},
  {max: 32, message: 'Must be less than 32 characters!'},
  {min: 6, message: 'Must be more than 6 characters!'},
  ({ getFieldValue }) => ({
    validator(_: any, value: any) {
      if (!value || getFieldValue('password') === value) {
        return Promise.resolve();
      }
      return Promise.reject(new Error('The new password that you entered do not match!'));
    },
  }),
]

export const emailRules: FormRule[] = [
  {type: 'email', message: 'The input is not valid E-mail!'},
  {required: true, message: 'Please input your E-mail!'},
]

export const captchaRules = [
  {required: true, message: 'Please input the captcha you got!',},
  {len: 6, message: 'The captcha must be 6 characters!'}
]

export const agreementRules: FormRule[] = [
  {
    validator: (_: any, value: any) =>
      value ? Promise.resolve() : Promise.reject(new Error('Should accept agreement')),
  },
]

export const titleRules: FormRule[] = [
  {required: true, message: 'Please input the title!', whitespace: true},
  {max: 32, message: 'Must be less than 32 characters!'},
]

export const descriptionRules: FormRule[] = [
  {required: true, message: 'Please input the description!', whitespace: true},
  {max: 100, message: 'Must be less than 100 characters!'},
]