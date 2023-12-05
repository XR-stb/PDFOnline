import { Avatar, Button, Popover, Space } from "antd";
import { UserOutlined } from '@ant-design/icons';
import { userPopoverStyle } from "./styles";
import { logout } from "../../api/pdfonline/user";
import { UserType } from "../../types";

interface UserProps {
  user: UserType | undefined
  setLoggedIn: React.Dispatch<React.SetStateAction<boolean>>
}

const User = ({user, setLoggedIn}: UserProps) => {
  return (
    <Popover placement="bottomRight" title={user?.username} content={<PopoverContent user={user} setLoggedIn={setLoggedIn} />} style={userPopoverStyle}>
      <Space>
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
  const onClick = () => {
    logout()
    setLoggedIn(false)
  }
  return (
    <>
      <p>Email: {user?.email}</p>
      <p>Role: {roleMap[user?.role as number]}</p>
      <Button type="primary" danger onClick={onClick}>Logout</Button>
    </>
  )
}

export default User