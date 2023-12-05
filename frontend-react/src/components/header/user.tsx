import {Avatar, Button, Popover, Space} from "antd";
import { UserOutlined } from '@ant-design/icons';
import {userPopoverStyle} from "./styles";
import {logout} from "../../api/pdfonline/user";

interface UserProps {
  user: any
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

const PopoverContent = ({user, setLoggedIn}: UserProps) =>{
  const onClick = () => {
    logout()
    setLoggedIn(false)
  }
  return (
    <Button type="primary" danger onClick={onClick}>Logout</Button>
  )
}

export default User