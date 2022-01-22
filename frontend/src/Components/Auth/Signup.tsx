import { useAuth0 } from '@auth0/auth0-react';

const SignupButton = () => {
  const { loginWithRedirect } = useAuth0();
  return (
    <a
      href="#0"
      className="link dib pa2 pv2 color-inherit ba br2"
      onClick={() =>
        loginWithRedirect({
          screen_hint: 'signup',
        })
      }
    >
      Sign Up
    </a>
  );
};

export default SignupButton;