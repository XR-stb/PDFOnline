import { Button } from "antd";

import LoginModal from "../modal/login";

interface LoginButtonProps {
  setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>
  showSignUpModal: () => void
}

const LoginButton = ({ setLoggedIn, showSignUpModal }: LoginButtonProps) => {
  const { loginModal, showLoginModal } = LoginModal({ setLoggedIn, showSignupModal: showSignUpModal });
  return (
    <>
      <Button
        size={'large'}
        onClick={showLoginModal}
      >
        Log In
      </Button>
      {loginModal}
    </>
  );
}

export default LoginButton;