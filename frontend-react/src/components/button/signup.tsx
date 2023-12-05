import React from "react";
import { Button } from "antd";

import SignUpModal from "../modal/signup";

interface SignUpButtonProps {
  setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>
}

const SignUpButton = ({ setLoggedIn }: SignUpButtonProps) => {
  const { signUpModal, showSignUpModal } = SignUpModal({ setLoggedIn });
  const signUpButton = (
    <>
      <Button size={'large'} type={'primary'} onClick={showSignUpModal}>
        Sign Up
      </Button>
      {signUpModal}
    </>
  )
  return { signUpButton, showSignUpModal };
}

export default SignUpButton;