import React, { ReactNode } from 'react';
import { Layout } from 'antd';
import {contentStyle} from "./styles";
import HeaderComponent from "../../components/header";
// import FooterComponent from "../../components/footer";

const { Content } = Layout;


interface ContainerProps {
  children: ReactNode;
}

const BasicContainer: React.FC<ContainerProps> = ( props ) => {
  return (
    <>
      <HeaderComponent />
      <Content style={contentStyle}>
        {props.children}
      </Content>
      {/*<FooterComponent />*/}
    </>
  );
};

export default BasicContainer;