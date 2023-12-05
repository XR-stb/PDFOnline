import {Button, Form, message} from "antd";
import {PasswordItem, UsernameItem} from "./items";
import {FormStyle} from "./styles";
import {login} from "../../../api/pdfonline/user";


interface LoginFormProps {
  setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>
}

const LoginForm = ({setLoggedIn}: LoginFormProps) => {
  const [form] = Form.useForm();
  const onFinish = async () =>  {
    login(form.getFieldsValue()).then(user_id => {
      setLoggedIn(true);
      message.success('Log in successfully!');
    }).catch((error) => {
      message.error(error.message);
    })
  };

  return (
    <Form
      form={form}
      name="login"
      onFinish={onFinish}
      style={FormStyle}
      scrollToFirstError
    >
      <Items />
      <Button type="primary" htmlType="submit" style={{width: '100%'}}>
        Log In
      </Button>
      <div>
        <p>Don't have an account? <a>Sign Up</a></p>
      </div>
    </Form>
  )
}

const Items = () => {
  return (
    <>
      <UsernameItem />
      <PasswordItem />
    </>
  )
}

export default LoginForm;