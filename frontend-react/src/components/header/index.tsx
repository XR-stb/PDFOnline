import React from "react";
import { Layout, Space } from 'antd';

import { headerStyle } from "./styles";
import Logo from "./logo";
import MenuComponent from "./menu";
import SignUpButton from "./signup";
import LoginButton from "./login";
import useUser from "../../auth/user";
import User from "./user";

const { Header } = Layout;

const HeaderComponent: React.FC = () => {
  const { loggedIn,user, setLoggedIn } = useUser()

  return (
    <Header style={headerStyle}>
      <Space>
        <Logo />
      </Space>
      <Space>
        <MenuComponent />
        {loggedIn ? <User user={user} setLoggedIn={setLoggedIn} /> : <><SignUpButton setLoggedIn={setLoggedIn} /><LoginButton setLoggedIn={setLoggedIn} /></>}
      </Space>
    </Header>
  );
}

export default HeaderComponent;