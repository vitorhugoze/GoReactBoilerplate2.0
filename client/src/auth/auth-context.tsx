import { ReactElement, createContext, useState } from 'react';
import { UserContext } from '../models/user-models';

export type AuthContent = {
  userAuth: boolean;
  setAuth(val: boolean): void;
  userData: UserContext;
  setUser(val: UserContext): void;
};

export const authContext = createContext<AuthContent | null>(null);

export const AuthProvider = (content: { children: ReactElement }) => {
  const [userAuth, setUserAuth] = useState(false);
  const [userData, setUserData] = useState<UserContext>({
    user_name: '',
    user_mail: '',
  });

  const setAuth = function (val: boolean) {
    setUserAuth(val);
  };

  const setUser = function (val: UserContext) {
    setUserData(val);
  };

  return <authContext.Provider value={{ userAuth, setAuth, userData, setUser }}>{content.children}</authContext.Provider>;
};
