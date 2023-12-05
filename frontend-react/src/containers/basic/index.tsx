import React, { ReactNode } from 'react';
import { Layout } from 'antd';

import { contentStyle } from "./styles";
import { UserType } from "../../types";
import HeaderComponent from "../../components/header";
// import FooterComponent from "../../components/footer";

interface ContainerProps {
  children: ReactNode;
  user:  { user: UserType | undefined, loggedIn: boolean, setLoggedIn: React.Dispatch<React.SetStateAction<boolean>> }
}

const BasicContainer: React.FC<ContainerProps> = ({ user, children }: ContainerProps) => {
  return (
    <>
      <HeaderComponent user={user} />
      <Layout.Content style={contentStyle}>
        {children}
      </Layout.Content>
      {/*<FooterComponent />*/}
    </>
  );
};

export default BasicContainer;