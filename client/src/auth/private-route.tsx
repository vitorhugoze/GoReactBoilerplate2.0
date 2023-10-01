import { useContext, useEffect } from 'react';
import axios from 'axios';
import { Outlet, redirect } from 'react-router-dom';
import { authContext } from './auth-context';

const PrivateRoute = () => {
  const context = useContext(authContext);

  const { userAuth, setAuth, userData, setUser } = context!;

  /*  useEffect(() => {
    axios
      .get('/auth')
      .then((res) => {
        if (res.status == 200) {
          setAuth(true);
        } else {
          setAuth(false);
          redirect('/');
        }
      })
      .catch(() => {
        setAuth(false);
        redirect('/');
      });
  }, []); */

  return <Outlet />;
};

export default PrivateRoute;
