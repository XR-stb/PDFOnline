import React from "react";
import { Layout } from 'antd';

const { Footer } = Layout;

function FooterComponent() {
  return (
    <Footer
      style={{
        textAlign: 'center',
      }}
    >
      PDF Online ©2023
    </Footer>
  );
}

export default FooterComponent;