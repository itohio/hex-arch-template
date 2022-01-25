import {
    ApolloProvider,
    ApolloClient,
    InMemoryCache,
    HttpLink,
    NormalizedCacheObject,
} from '@apollo/client';
import { setContext } from '@apollo/client/link/context';
import { useAuth0 } from '@auth0/auth0-react';
import React from 'react';

const ApolloProviderWithAuth0 = ({ children }: {children: any}) => {
    const { isAuthenticated, getAccessTokenSilently } = useAuth0();

    const httpLink = new HttpLink({
        uri: process.env.REACT_APP_API_URL,
        // Uncomment for CORS
        // credentials: 'same-origin',
        credentials: 'include',
    });

    const authLink = setContext(async (_, { headers, fetchOptions, ...rest }) => {
        if (!isAuthenticated) return { headers, fetchOptions, ...rest };

        let token;
        try {
            token = await getAccessTokenSilently();
        } catch (error) {
            console.log(error);
        }

        if (!token) return { headers, fetchOptions, ...rest };

        return {
            ...rest,
            headers: {
                ...headers,
                authorization: `Bearer ${token}`,
            },
            fetchOptions,
        };
    });

    const client = React.useRef<ApolloClient<NormalizedCacheObject>>();

    if (!client.current) {
      client.current = new ApolloClient({
        link: authLink.concat(httpLink),
        cache: new InMemoryCache(),
        credentials: 'include',
        headers: {
            'Access-Control-Allow-Origin': '*',
            'Access-Control-Allow-Headers': '*',
        }
      });
    }

    return (
        <ApolloProvider client={client.current}>
            {children}
        </ApolloProvider>
    );
};

export default ApolloProviderWithAuth0;