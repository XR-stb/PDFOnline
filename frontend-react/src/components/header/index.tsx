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
  const { loggedIn,user, setMyUser } = useUser()

  return (
    <Header style={headerStyle}>
      <Space>
        <Logo />
      </Space>
      <Space>
        <MenuComponent />
        {loggedIn ? <User user={user} setMyUser={setMyUser} /> : <><SignUpButton setMyUser={setMyUser} /><LoginButton setMyUser={setMyUser} /></>}
      </Space>
    </Header>
  );
}

export default HeaderComponent;