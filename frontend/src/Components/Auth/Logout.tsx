import { useAuth0 } from "@auth0/auth0-react";

function LogoutButton() {
  const { logout } = useAuth0();

  return (
    <a href="#0"
      className="link dib pa2 pv2 color-inherit"
      onClick={() => logout({ returnTo: window.location.origin })}>
      Log Out
    </a>
  );
};

export default LogoutButton;