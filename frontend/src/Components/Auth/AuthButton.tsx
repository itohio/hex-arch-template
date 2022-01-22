import React from 'react';

import LoginButton from './Login';
import LogoutButton from './Logout';
import SignupButton from './Signup';

import { useAuth0 } from '@auth0/auth0-react';

const AuthenticationButton = () => {
  const { isAuthenticated } = useAuth0();

  return isAuthenticated ? <LogoutButton /> : (
    <div className="link dib pa2 pv2 color-inherit">
    <LoginButton />
    <SignupButton />
    </div>
  );
};

export default AuthenticationButton;