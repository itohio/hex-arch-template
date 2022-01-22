// src/auth/protected-route.js

import React from 'react';
import { Route, RouteProps } from 'react-router-dom';
import { withAuthenticationRequired } from '@auth0/auth0-react';
import Loading from '../Loading';

interface ProtectedRouteProps extends RouteProps {
    component: React.ComponentType;
}

const ProtectedRoute = ({ component, ...args }: ProtectedRouteProps) => (
  <Route
    component={withAuthenticationRequired(component, {
      onRedirecting: () => <Loading />,
    })}
    {...args}
  />
);

export default ProtectedRoute;