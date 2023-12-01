import {Button, Checkbox, Form, Input, Space} from "antd";
import {usernameRules, passwordRules, confirmPasswordRules, emailRules, captchaRules, agreementRules} from "./rules";
import {MouseEventHandler, useEffect, useState} from "react";

export const UsernameItem = () => {
  return (
    <Form.Item
      name="username"
      label="Username"
      rules={usernameRules}
    >
      <Input/>
    </Form.Item>
  );
}

export const PasswordItem = () => {
  return (
    <Form.Item
      name="password"
      label="Password"
      rules={passwordRules}
      hasFeedback
    >
      <Input.Password/>
    </Form.Item>
  );
}

export const ConfirmPasswordItem = () => {
  return (
    <Form.Item
      name="confirm"
      label="Confirm Password"
      dependencies={['password']}
      hasFeedback
      rules={confirmPasswordRules}
    >
      <Input.Password/>
    </Form.Item>
  );
}

export const EmailItem = () => {
  return (
    <Form.Item
      name="email"
      label="E-mail"
      rules={emailRules}
    >
      <Input/>
    </Form.Item>
  );
}

interface CaptchaProps {
  handleGetCaptcha: MouseEventHandler<HTMLElement>;
  isCounting: boolean;
  countdown: number;
}

export const CaptchaItem = ({ handleGetCaptcha, isCounting, countdown }: CaptchaProps) => {
  return (
    <Form.Item label="Captcha">
      <Space>
        <Form.Item
          name="captcha"
          noStyle
          rules={captchaRules}
        >
          <Input />
        </Form.Item>
        <Button onClick={handleGetCaptcha} disabled={isCounting} style={{ width: '120px' }}>
          {isCounting ? `${countdown} s` : 'Get Captcha'}
        </Button>
      </Space>
    </Form.Item>
  );
}

export const AgreementItem = () => {
  return (
    <Form.Item
      name="agreement"
      valuePropName="checked"
      rules={agreementRules}
    >
      <Checkbox>
        I have read the <a>agreement</a>
      </Checkbox>
    </Form.Item>
  );
}