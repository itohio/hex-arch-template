import { gql, useMutation } from '@apollo/client';
import { HelloWorld, HelloWorldVariables } from './generated/HelloWorld';


export const HELLO_WORLD = gql`
    mutation HelloWorld($input: Input!) {
        helloWorld(input: $input)
    }
`;

export const useHelloWorld = () => {
    return useMutation<HelloWorld, HelloWorldVariables>(HELLO_WORLD);
}