import { Link } from 'react-router-dom';

function MainNav() {
  return (
    <nav id="main-nav" className="dn db-l">
      <Link to={'/profile'} className="link db dib-l pa2 pv2 color-inherit">Profile</Link>
      <Link to={'/about'} className="link db dib-l pa2 pv2 color-inherit">About</Link>
      <Link to={'/contact'} className="link db dib-l pa2 pv2 color-inherit">Contact</Link>
    </nav>
  );
}

export default MainNav;