import React, { useState } from "react";
import { Modal } from "antd";

import SignUpForm from "../form/signup";

interface SignupModalProps {
  setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>
}

const useSignUpModal = ({ setLoggedIn }: SignupModalProps) => {
  const [isSignUpModalOpen, setSignUpModalOpen] = useState(false)
  const showSignUpModal = () => setSignUpModalOpen(true)
  const hideSignUpModal = () => setSignUpModalOpen(false)

  const signUpModal = (
    <Modal title="Sign Up" footer={null} open={isSignUpModalOpen} onCancel={hideSignUpModal}>
      <SignUpForm setLoggedIn={setLoggedIn} />
    </Modal>
  )

  return { signUpModal, showSignUpModal };
}

export default useSignUpModal;