import React from 'react';
import { useHistory } from 'react-router-dom';
import { Auth0Provider } from '@auth0/auth0-react';

const Auth0ProviderWithHistory = ({ children }: {children: any}) => {
  const history = useHistory();
  const domain = `${process.env.REACT_APP_AUTH0_DOMAIN}`;
  const clientId = `${process.env.REACT_APP_AUTH0_CLIENT_ID}`;
  const audience = `${process.env.REACT_APP_AUTH0_AUDIENCE}`;

  const onRedirectCallback = (appState: any) => {
    history.push(appState?.returnTo || window.location.pathname);
  };

  return (
    <Auth0Provider
      domain={domain}
      clientId={clientId}
      audience={audience}
      scope="read:all"
      redirectUri={window.location.origin}
      onRedirectCallback={onRedirectCallback}
    >
      {children}
    </Auth0Provider>
  );
};

export default Auth0ProviderWithHistory;