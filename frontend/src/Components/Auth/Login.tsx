import { useAuth0 } from "@auth0/auth0-react";

function LoginButton() {
  const { loginWithRedirect } = useAuth0();

  return <a 
      href="#0"
      className="link dib pa2 pv2 color-inherit"
      onClick={() => loginWithRedirect()}>
        Log In
    </a>;
};

export default LoginButton;