import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import { useAuth0 } from '@auth0/auth0-react';
import ProtectedRoute from './Components/Auth/Protected';
import Header from './Components/Header';
import Loading from './Components/Loading';
import Home from './Components/Pages/Home';
import About from './Components/Pages/About';
import Contact from './Components/Pages/Contact'
import Profile from './Components/Auth/Profile'
import { Container } from '@mui/material';
import Footer from './Components/Footer';

const sections = [
  { title: 'Profile', url: '/profile' },
  { title: 'About', url: '/about' },
  { title: 'Contacts', url: '/contact' },
];

function App() {
  const { isLoading } = useAuth0();

  if (isLoading) {
    return <Loading/>;
  }

  return (
    <BrowserRouter>
      <Container>
        <Header title="Hex Architecture" sections={sections}/>
        <main>
          <Switch>
            <Route exact path='/' component={Home} />
            <ProtectedRoute path='/profile' component={Profile} />
            <Route path='/contact' component={Contact} />
            <Route path='/about' component={About} />
          </Switch>
        </main>
        <Footer
          title="Footer"
          description="Something here to give the footer a purpose!"
        />
      </Container>
    </BrowserRouter>
  );
}

export default App;
