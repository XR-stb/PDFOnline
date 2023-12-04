import {useState} from "react";
import {Button, Modal} from "antd";
import SignUpForm from "./forms/signup";

interface SignUpButtonProps {
  setMyUser: (user_id: string) => void;
}

const SignUpButton = ({setMyUser}: SignUpButtonProps) => {
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
        type={'primary'}
        onClick={showModal}
      >
        Sign Up
      </Button>
      <Modal title="Sign Up" footer={null} open={isModalOpen} onCancel={handleCancel}>
        <SignUpForm setMyUser={setMyUser} />
      </Modal>
    </>
  );
}

export default SignUpButton;