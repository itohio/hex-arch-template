import {
    ApolloProvider,
    ApolloClient,
    InMemoryCache,
    HttpLink,
} from '@apollo/client';
import { setContext } from '@apollo/link-context';
import { useAuth0 } from '@auth0/auth0-react';

const ApolloProviderWithAuth0 = ({ children }: {children: any}) => {
    const { getAccessTokenSilently } = useAuth0();

    const httpLink = new HttpLink({
        uri: process.env.REACT_APP_API_URI,
    });

    const authLink = setContext(async (_, { headers, ...rest }) => {
        let token;
        try {
            token = await getAccessTokenSilently();
        } catch (error) {
            console.log(error);
        }

        if (!token) return { headers, ...rest };

        return {
            ...rest,
            headers: {
                ...headers,
                authorization: `Bearer ${token}`,
            },
        };
    });

    return new ApolloClient({
        link: authLink.concat(httpLink),
        cache: new InMemoryCache(),
        connectToDevTools: true
    });
};

export default ApolloProviderWithAuth0;