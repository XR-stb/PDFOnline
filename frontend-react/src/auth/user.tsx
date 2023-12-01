import { useState } from "react";
import { getUserInfo, User } from "../api/pdfonline/user";
import { message } from "antd";

export type UserContextType = User | undefined;

const useUser = () => {
  const [loggedIn, setLoggedIn] = useState(false);
  const [user, setUser] = useState(undefined as UserContextType);

  const setMyUser = (user_id: string) => {
    if (user_id === "") {
      setLoggedIn(false);
      setUser(undefined);
      return;
    }

    getUserInfo(user_id).then(user => {
      setUser(user);
      setLoggedIn(true);
    }).catch(error => {
      message.error(`Failed to get user info: ${error.message}`);
    });
  };

  return { loggedIn, user, setMyUser };
};

export default useUser;
