import React from "react";
import { Button, ConfigProvider } from 'antd';

const Logo: React.FC = () => {
  return (
    <ConfigProvider
      theme={{
        components: {
          Button: {
            textHoverBg: '#fff',
          },
        },
      }}
    >
      <Button
        href={`/`}
        type="text"
        size={'large'}
      >
        PDF Online
      </Button>
    </ConfigProvider>
  );
}

export default Logo;