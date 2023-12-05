import { useState } from "react";
import { Button, Form, message, Modal } from "antd";
import { FormStyle } from "./styles";
import { PasswordItem, UsernameItem } from "./items";
import { login } from "../../api/pdfonline/user";


interface LoginFormProps {
  setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>
}

const LoginForm = ({setLoggedIn}: LoginFormProps) => {
  const [form] = Form.useForm();
  const onFinish = () =>  {
    Promise.resolve(login(form.getFieldsValue())).then(user_id => {
      setLoggedIn(true);
      message.success('Log in successfully!');
    }).catch((error) => {
      message.error(error.message);
    })
  };

  const [isModalOpen, setIsModalOpen] = useState(false);

  const handleGuest = () => {
    setIsModalOpen(true);
  };

  const handleOk = () => {
    form.setFieldValue('username', 'guest');
    form.setFieldValue('password', 'guest123');
    form.submit()
    setIsModalOpen(false);
  };

  const handleCancel = () => {
    setIsModalOpen(false);
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
        <p><>Don't have an account? </>
          <a>
            Sign Up
          </a>
          <> or </>
          <Modal title="Reconfirm" open={isModalOpen} onOk={handleOk} onCancel={handleCancel} okText={"Yes"}>
            Are you sure to log in as a guest?
          </Modal>
          <a onClick={handleGuest}>
            Guest
          </a>
        </p>
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