import React, { useState } from "react";
import { Modal } from "antd";

import LoginForm from "../form/login";

interface LoginModalProps {
  setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>
  showSignupModal: () => void
}

const LoginModal = ({ setLoggedIn, showSignupModal }: LoginModalProps) => {
  const [isLoginModalOpen, setLoginModalOpen] = useState(false)
  const showLoginModal = () => setLoginModalOpen(true)
  const hideLoginModal = () => setLoginModalOpen(false)

  const loginModal = (
    <Modal title="Sign Up" footer={null} open={isLoginModalOpen} onCancel={hideLoginModal}>
      <LoginForm setLoggedIn={setLoggedIn} hideLoginModal={hideLoginModal} showSignUpModal={showSignupModal} />
    </Modal>
  )

  return { loginModal, showLoginModal };
}

export default LoginModal;