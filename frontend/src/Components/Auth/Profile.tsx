import React, {useEffect, useState } from "react";
import { useAuth0 } from "@auth0/auth0-react";
import { useGreetings } from '../../gql/query'


const Profile = () => {
  const { user, getAccessTokenSilently } = useAuth0();
  const [userMetadata, setUserMetadata] = useState("");
  const { data, loading } = useGreetings();

  useEffect(() => {
    const getUserMetadata = async () => {
      try {
        const accessToken = await getAccessTokenSilently({
          audience: process.env.REACT_APP_AUTH0_AUDIENCE,
          scope: "read:all",
        });
  
        setUserMetadata(accessToken);
      } catch (e) {
        console.log(e);
      }
    };
  
    getUserMetadata();
  }, [getAccessTokenSilently, user?.sub]);

  return (
    <div>
      <div className="row align-items-center profile-header">
        <div className="col-md-2 mb-3">
          <img
            src={user?.picture}
            alt="Profile"
            className="rounded-circle img-fluid profile-picture mb-3 mb-md-0"
          />
        </div>
        <div className="col-md text-center text-md-left">
          <h2>{user?.name}</h2>
          <p className="lead text-muted">{user?.email}</p>
        </div>
      </div>
      <div className="row">
        <pre className="col-12 text-light bg-dark p-4">
          {JSON.stringify(user, null, 2)}
        </pre>
      </div>
      <div className="row">
        <pre className="col-12 text-light bg-dark p-4">
          {JSON.stringify(userMetadata)}
        </pre>
      </div>
      {loading?(
        <div className="row">
        <pre className="col-12 text-light bg-dark p-4">
          {data?.greetings}
        </pre>
      </div>
      ): (<div></div>)}
    </div>
  );
};

export default Profile;