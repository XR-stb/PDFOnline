import React, { ReactNode } from 'react';
import { Layout } from 'antd';
import { contentStyle } from "./styles";
import HeaderComponent from "../../components/header";
import {UserType} from "../../types";
// import FooterComponent from "../../components/footer";

const { Content } = Layout;

interface ContainerProps {
  children: ReactNode;
  user:  {user: UserType | undefined, loggedIn: boolean, setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>}
}

const BasicContainer: React.FC<ContainerProps> = ( { user, children } ) => {
  return (
    <>
      <HeaderComponent user={user} />
      <Content style={contentStyle}>
        {children}
      </Content>
      {/*<FooterComponent />*/}
    </>
  );
};

export default BasicContainer;