import React from 'react';
import { Link } from 'react-router-dom';
import MainNav from './Nav/Main'
import AuthNav from './Nav/Auth'
import logo from '../logo.svg';

function Header() {
  return (
    <header className="pv3 ph2 relative bb">
    <div id="container" className="flex justify-between items-center">
        <div id="logo" className="overflow-hidden">
        <Link to={'/'} className="link color-inherit">
            <img src={logo} alt="logo" />
            Logo
        </Link>
        </div>

        <MainNav/>
        <AuthNav/>
    </div>
    </header>
  );
}

export default Header;
