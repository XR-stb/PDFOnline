import {Avatar, Button, Divider, Popconfirm, Popover, Space} from "antd";
import { UserOutlined, LogoutOutlined } from '@ant-design/icons';

import {logoutIconStyle, popoverTitleStyle, userPopoverStyle} from "./styles";
import { logout } from "../../api/pdfonline/user";
import { UserType } from "../../types";

interface UserProps {
  user: UserType | undefined
  setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>
}

const UserDock = ({user, setLoggedIn}: UserProps) => {
  return (
    <Popover placement={"bottomRight"} title={<p style={popoverTitleStyle}>{user?.username}</p>} content={<PopoverContent user={user} setLoggedIn={setLoggedIn} />} style={userPopoverStyle} trigger={"click"}>
      <Space style={{cursor: 'pointer'}}>
        <Avatar icon={<UserOutlined />} />
        {user?.username}
      </Space>
    </Popover>
  )
}

const roleMap:{[key: number]: string} = {
  0: "Guest",
  1: "User",
  2: "Admin"
}

const PopoverContent = ({user, setLoggedIn}: UserProps) =>{
  const handleLogout = () => {
    logout()
    setLoggedIn(false)
  }

  return (
    <>
      <p><strong>Email: </strong>{user?.email}</p>
      <p><strong>Role: </strong>{roleMap[user?.role as number]}</p>
      <Popconfirm title={"Are you sure to logout?"} onConfirm={handleLogout} okType={'danger'} okText={'Yes'} cancelText={'No'} placement={'bottom'} icon={<LogoutOutlined style={logoutIconStyle} />}>
        <Button type="primary" danger>Logout</Button>
      </Popconfirm>
    </>
  )
}

export default UserDock