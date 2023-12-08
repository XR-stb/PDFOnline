import React from "react";
import { Button } from "antd";

import useSignUpModal from "../modal/signup";

interface SignUpButtonProps {
  setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>
}

const SignUpButton = ({ setLoggedIn }: SignUpButtonProps) => {
  const { signUpModal, showSignUpModal } = useSignUpModal({ setLoggedIn });
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