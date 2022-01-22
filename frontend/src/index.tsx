import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import reportWebVitals from './reportWebVitals';
import './index.css';
import Auth0ProviderWithHistory from './Components/Auth/Provider';

ReactDOM.render(
  <React.StrictMode>
    <Auth0ProviderWithHistory>
      <App />
    </Auth0ProviderWithHistory>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
