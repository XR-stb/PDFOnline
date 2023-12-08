import React, { useEffect, useState } from "react";
import { Button, Form, message } from "antd";

import { CaptchaItem, ConfirmPasswordItem, EmailItem, PasswordItem, UsernameItem } from "./items";
import { formStyle, signUpButtonStyle } from "./styles";
import { getResetPasswordCaptcha,  resetPassword} from "../../api/pdfonline/user";

interface ResetPasswordFormProps {
  hideResetPasswordModal: () => void;
}

const ResetPasswordForm = (props: ResetPasswordFormProps) => {
  const [ form ] = Form.useForm();
  const onFinish = async () =>  {
    Promise.resolve(resetPassword(form.getFieldsValue())).then(user_id => {
      message.success('Reset password successfully!');
      props.hideResetPasswordModal();
    }).catch((error) => {
      message.error(error.message);
    })
  };

  const [isCounting, setIsCounting] = useState(false);
  const [countdown, setCountdown] = useState(60);
  const HandleGetCaptcha = async () => {
    try {
      await form.validateFields(['email']);

      let data = {
        username: form.getFieldValue('username'),
        email: form.getFieldValue('email'),
      }

      getResetPasswordCaptcha(data).then(() => {
        setIsCounting(true);
      }).catch((error) => {
        let errmsg = error.message;
        if (error.status === 404) {
          errmsg = 'username not exist!';
          form.setFields([{
            name: 'username',
            errors: [errmsg],
          }]);
        }

        if (error.status === 400) {
          errmsg = 'email not match!';
          form.setFields([{
            name: 'email',
            errors: [errmsg],
          }]);
        }

        message.error(errmsg);
      });
    } catch (error) {
      return
    }
  };
  useEffect(() => {
    let countdownInterval: NodeJS.Timer;

    if (isCounting) {
      countdownInterval = setInterval(() => {
        setCountdown((prevCountdown) => {
          if (prevCountdown === 0) {
            clearInterval(countdownInterval);
            setIsCounting(false);
            return 60; // 重置为初始状态
          }
          return prevCountdown - 1;
        });
      }, 1000);
    }

    return () => {
      clearInterval(countdownInterval);
    };
  }, [isCounting]);

  return (
    <Form form={form} name="resetPassword" onFinish={onFinish} style={formStyle}>
      <UsernameItem />
      <EmailItem />
      <CaptchaItem handleGetCaptcha={HandleGetCaptcha} isCounting={isCounting} countdown={countdown} />
      <PasswordItem />
      <ConfirmPasswordItem />
      <Button type="primary" htmlType="submit" style={signUpButtonStyle}>
        Reset Password
      </Button>
    </Form>
  )
}

export default ResetPasswordForm;