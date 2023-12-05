import React from "react";
import { Layout, Space } from 'antd';
import { headerStyle } from "./styles";
import Logo from "./logo";
import MenuComponent from "./menu";
import SignUpButton from "./signup";
import LoginButton from "./login";
import User from "./user";
import {UserType} from "../../types";

const { Header } = Layout;

interface HeaderProps {
  user: {user: UserType | undefined, loggedIn: boolean, setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>}
}

const HeaderComponent: React.FC<HeaderProps> = ( { user} ) => {
  return (
    <Header style={headerStyle}>
      <Space>
        <Logo />
      </Space>
      <Space>
        <MenuComponent />
        {user.loggedIn ? <User user={user.user} setLoggedIn={user.setLoggedIn} /> : <><SignUpButton setLoggedIn={user.setLoggedIn} /><LoginButton setLoggedIn={user.setLoggedIn} /></>}
      </Space>
    </Header>
  );
}

export default HeaderComponent;