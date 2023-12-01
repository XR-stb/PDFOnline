import {Button, Modal} from "antd";
import {useState} from "react";
import LoginForm from "./forms/login";

interface LoginButtonProps {
  setMyUser: (user_id: string) => void;
}

const LoginButton = ({setMyUser}: LoginButtonProps) => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const showModal = () => {
    setIsModalOpen(true);
  };
  const handleCancel = () => {
    setIsModalOpen(false);
  };

  return (
    <>
      <Button
        size={'large'}
        onClick={showModal}
      >
        Log In
      </Button>
      <Modal title="Log In" footer={null} open={isModalOpen} onCancel={handleCancel}>
        <LoginForm setMyUser={setMyUser} />
      </Modal>
    </>
  );
}

export default LoginButton;