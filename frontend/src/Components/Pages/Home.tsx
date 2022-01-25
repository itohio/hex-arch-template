
const Home = () => {
    return (
        <div>
          <h2>{process.env.REACT_APP_AUTH0_DOMAIN}</h2>
          <h2>{process.env.REACT_APP_AUTH0_CLIENT_ID}</h2>          
        </div>
    );
}

export default Home;