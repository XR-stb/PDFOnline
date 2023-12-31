import React, { useEffect, useState } from "react";
import { Button, Form, message } from "antd";

import { AgreementItem, CaptchaItem, ConfirmPasswordItem, EmailItem, PasswordItem, UsernameItem } from "./items";
import { formStyle, signUpButtonStyle } from "./styles";
import { getRegisterCaptcha, register } from "../../api/pdfonline/user";

interface SignUpFormProps {
  setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>
}

const SignUpForm = ({ setLoggedIn }: SignUpFormProps) => {
  const [ form ] = Form.useForm();
  const onFinish = async () =>  {
    register(form.getFieldsValue()).then(user_id => {
      setLoggedIn(true);
      message.success('Sign Up successfully!');
    }).catch((error) => {
      message.error(error.message);
    })
  };

  const [isCounting, setIsCounting] = useState(false);
  const [countdown, setCountdown] = useState(60);
  const HandleGetCaptcha = async () => {
    try {
      await form.validateFields(['email']);

      getRegisterCaptcha(form.getFieldValue('email')).then(() => {
        setIsCounting(true);
      }).catch((error) => {
        message.error(error.message);
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
    <Form form={form} name="signup" onFinish={onFinish} style={formStyle}>
      <UsernameItem />
      <PasswordItem />
      <ConfirmPasswordItem />
      <EmailItem />
      <CaptchaItem handleGetCaptcha={HandleGetCaptcha} isCounting={isCounting} countdown={countdown} />
      <AgreementItem />
      <Button type="primary" htmlType="submit" style={signUpButtonStyle}>
        Sign Up
      </Button>
    </Form>
  )
}

export default SignUpForm;