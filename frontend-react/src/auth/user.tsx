import { useState, useEffect } from "react";
import { message } from "antd";

import { UserType } from "../types";
import { getMe } from "../api/pdfonline/user";

const useUser = () => {
  const [ user, setUser ] = useState<UserType | undefined>(undefined);
  const [ loggedIn, setLoggedIn ] = useState(false);

  const fetchUser = (ignoreError: boolean = false) => {
    if ( user === undefined ) {
      Promise.resolve(getMe()).then((user) => {
        setUser(user);
        setLoggedIn(true);
      }).catch((error) => {
        if (!ignoreError) {
          message.error(error.message);
        }
      })
    }
  };

  useEffect(() => {
    fetchUser(true);
  }, []);

  useEffect(() => {
    loggedIn ? fetchUser() : setUser(undefined);
  }, [loggedIn]);

  return { user, loggedIn, setLoggedIn };
};

export default useUser;