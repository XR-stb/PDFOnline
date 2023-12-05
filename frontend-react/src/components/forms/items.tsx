import {Button, Checkbox, Form, Input, Space, UploadProps, Upload, message, FormInstance} from "antd";
import { InboxOutlined } from '@ant-design/icons';
import {
  usernameRules,
  passwordRules,
  confirmPasswordRules,
  emailRules,
  captchaRules,
  agreementRules,
  titleRules, descriptionRules
} from "./rules";

import { MouseEventHandler} from "react";

export const UsernameItem = () => {
  return (
    <Form.Item
      name="username"
      label="Username"
      rules={usernameRules}
    >
      <Input placeholder={"Username"} />
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
      <Input.Password placeholder="Password" />
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
      <Input.Password placeholder="Confirm Password" />
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
      <Input placeholder="example@mail.com" />
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
          <Input placeholder="6-digit code" />
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

export const TitleItem = () => {
  return (
    <Form.Item
      name="title"
      label="Title"
      rules={titleRules}
    >
      <Input showCount maxLength={32} placeholder="title" />
    </Form.Item>
  );
}

interface DescriptionProps {
  onChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
}

export const DescriptionItem = ({ onChange }: DescriptionProps) => {
  return (
    <Form.Item
      name="description"
      label="Description"
      rules={descriptionRules}
    >
      <Input.TextArea
        showCount
        maxLength={100}
        placeholder="description"
        style={{ height: 120, resize: 'none' }}
        onChange={onChange}
      />
    </Form.Item>
  );
}

export const UploadDraggerItem = () => {
  return (
    <Form.Item
      name="file"
      rules={[{ required: true, message: 'Please upload your PDF!' }]}
    >
      <Upload.Dragger
        name={'file'}
        maxCount={1}
      >
        <p className="ant-upload-drag-icon">
          <InboxOutlined />
        </p>
        <p className="ant-upload-text">Click or drag file to this area to upload</p>
        <p className="ant-upload-hint">
          Support for PDF only.
        </p>
      </Upload.Dragger>
    </Form.Item>
  );
}