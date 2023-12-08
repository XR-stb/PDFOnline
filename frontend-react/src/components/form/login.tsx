import { Button, Form, message, Popconfirm } from "antd";
import { UserOutlined } from "@ant-design/icons";

import { formStyle, guestLoginIconStyle, loginButtonStyle } from "./styles";
import { PasswordItem, RememberMeItem, UsernameItem } from "./items";
import { login, loginGuest } from "../../api/pdfonline/user";

interface LoginFormProps {
  setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>
  hideLoginModal: () => void
  showSignUpModal: () => void
}

const LoginForm = ({ setLoggedIn, hideLoginModal, showSignUpModal}: LoginFormProps) => {
  const [form] = Form.useForm();
  const onFinish = () =>  {
    Promise.resolve(login(form.getFieldsValue())).then(user_id => {
      setLoggedIn(true);
      message.success('Log in successfully!');
    }).catch((error) => {
      message.error(error.message);
    })
  };

  const handleOk = () => {
    Promise.resolve(loginGuest()).then(user_id => {
      setLoggedIn(true);
      message.success('Log in successfully!');
    }).catch((error) => {
      message.error(error.message);
    })
  };

  const handleSignUp = () => {
    hideLoginModal();
    showSignUpModal();
  }

  return (
    <Form
      form={form}
      name="login"
      onFinish={onFinish}
      style={formStyle}
      scrollToFirstError
    >
      <Items />
      <Button type="primary" htmlType="submit" style={loginButtonStyle}>
        Log In
      </Button>
      <p>
        <>Don't have an account? </>
        <a onClick={handleSignUp}>Sign Up</a>
        <> or </>
        <Popconfirm title={"Are you sure to log in as guest?"} placement={"bottom"} okText={"Yes"} cancelText={"No"} onConfirm={handleOk} icon={<UserOutlined style={guestLoginIconStyle} />}>
          <a>Guest</a>
        </Popconfirm>
      </p>
    </Form>
  )
}

const Items = () => {
  return (
    <>
      <UsernameItem />
      <PasswordItem />
      <RememberMeItem />
    </>
  )
}

export default LoginForm;