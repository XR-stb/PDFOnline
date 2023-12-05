import { useState, useEffect } from "react";
import { getMe, User } from "../api/pdfonline/user";
import { message } from "antd";

export type UserContextType = User | undefined;

const useUser = () => {
  const [user, setUser] = useState<UserContextType>(undefined);
  const [loggedIn, setLoggedIn] = useState(false);

  const fetchUser = () => {
    Promise.resolve(getMe()).then((user) => {
      setUser(user);
      setLoggedIn(true);
    }).catch((error) => {
      message.error(error.message);
    });
  };

  useEffect(fetchUser, []);
  useEffect(() => {
    loggedIn ? fetchUser() : setUser(undefined);
  }, [loggedIn]);

  return { user, loggedIn, setLoggedIn };
};

export default useUser;