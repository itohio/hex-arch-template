import { Input, Button } from '@mui/material';
import { useState } from 'react';
import { useHelloWorld } from '../../gql/mutation';

const About = () => {
  const [name, setName] = useState("");
  const [helloWorld, {data}] = useHelloWorld();

  return (
      <div>
        <h2>About</h2>
        <h2>Auth0 domain: {process.env.REACT_APP_AUTH0_DOMAIN}</h2>
        <h2>Auth0 client id: {process.env.REACT_APP_AUTH0_CLIENT_ID}</h2>  
        <h2>app uri: {process.env.REACT_APP_API_URL}</h2>
        <h2>audience: {process.env.REACT_APP_AUTH0_AUDIENCE}</h2>
        <h2>{data?.helloWorld}</h2>
        <form onSubmit={async (e) => {
          e.preventDefault();
        }}>
          <Input 
            type="text" 
            placeholder="What is your name?"
            onChange={(e) => setName(e.target.value)}
          ></Input>
          <Button
            onClick={()=>helloWorld({variables: {input: {name: name}}})}
          >
            Greet
          </Button>
        </form>
      </div>
  );
}

export default About;