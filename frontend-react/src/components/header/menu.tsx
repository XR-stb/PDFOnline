import React from "react";
import { Menu } from "antd";
import {useNavigate} from "react-router-dom";

const item = (label: string, key: string) => {
  return {
    label,
    key,
  }
}

const Items= [
  item('HOME', '/home'),
  item('PDF', '/pdf'),
  item('ABOUT', '/about'),
];

const MenuComponent: React.FC = () => {
  const navigate = useNavigate()

  const onClick = (e: any) => {
    navigate(e.key, {replace: true})
  }

  return (
    <>
      <Menu
        theme="light"
        mode="horizontal"
        items={Items}
        onClick={onClick}
        selectable={false}
      />
    </>
  )
}

export default MenuComponent;