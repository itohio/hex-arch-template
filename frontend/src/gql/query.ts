import { gql, QueryResult, useQuery } from '@apollo/client';
import { GetGreetings } from './generated/GetGreetings';


export const GET_GREETINGS = gql`
  query GetGreetings {
    greetings
  }
`;

export const useGreetings = (): QueryResult<GetGreetings> => {
  return useQuery<GetGreetings>(GET_GREETINGS)
}
