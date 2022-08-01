import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import React, { createContext, useState } from "react";
import {
  BrowserRouter as Router,
  Routes,
  Switch,
  Route,
  Link
} from "react-router-dom";
import Home from './Components/Home/Home';
import Login from './Components/Login/Login';

export const LoggedInContext = createContext();
function App() {
  const [loggedInUser, setLoggedInUser] = useState({});
  return (
    <LoggedInContext.Provider value={[loggedInUser, setLoggedInUser]}>
        <Router>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<Login />} />
          </Routes>
        </Router>
    </LoggedInContext.Provider>
  );
}

export default App;
