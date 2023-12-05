import React from "react";
import { Layout, Space } from 'antd';

import { headerStyle } from "./styles";
import { UserType } from "../../types";
import Logo from "./logo";
import MenuComponent from "./menu";
import SignUpButton from "../button/signup";
import LoginButton from "../button/login";
import UserDock from "./userDock";

interface HeaderProps {
  user: { user: UserType | undefined, loggedIn: boolean, setLoggedIn: React.Dispatch<React.SetStateAction<boolean>> }
}

const HeaderComponent: React.FC<HeaderProps> = ({ user }: HeaderProps) => {
  const { setLoggedIn } = user
  const { signUpButton, showSignUpModal } = SignUpButton({ setLoggedIn })
  return (
    <Layout.Header style={headerStyle}>
      <Space>
        <Logo />
      </Space>
      <Space>
        <MenuComponent />
        {
          user.loggedIn ?
          <UserDock user={user.user} setLoggedIn={user.setLoggedIn} /> :
          <>
            {signUpButton}
            <LoginButton setLoggedIn={user.setLoggedIn} showSignUpModal={showSignUpModal}/>
          </>
        }
      </Space>
    </Layout.Header>
  );
}

export default HeaderComponent;