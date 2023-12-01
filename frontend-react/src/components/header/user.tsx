import {Avatar, Button, Popover, Space} from "antd";
import { UserOutlined } from '@ant-design/icons';
import {userPopoverStyle} from "./styles";
import {logout} from "../../api/pdfonline/user";

interface UserProps {
  user: any
  setMyUser: (user_id: string) => void;
}

const User = ({user, setMyUser}: UserProps) => {
  return (
    <Popover placement="bottomRight" title={user.username} content={<PopoverContent user={user} setMyUser={setMyUser} />} style={userPopoverStyle}>
      <Space>
        <Avatar icon={<UserOutlined />} />
        {user.username}
      </Space>
    </Popover>
  )
}

const PopoverContent = ({user, setMyUser}: UserProps) =>{
  const onClick = () => {
    logout()
    setMyUser('')
  }
  return (
    <Button type="primary" danger onClick={onClick}>Logout</Button>
  )
}

export default User