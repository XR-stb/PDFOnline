import { Button } from "antd";

import useLoginModal from "../modal/login";

interface LoginButtonProps {
  setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>
  showSignUpModal: () => void
}

const LoginButton = ({ setLoggedIn, showSignUpModal }: LoginButtonProps) => {
  const { loginModal, showLoginModal } = useLoginModal({ setLoggedIn, showSignupModal: showSignUpModal });
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