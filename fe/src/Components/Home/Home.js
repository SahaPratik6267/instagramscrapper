import React from 'react';
import Header from '../Header/Header';
import SearchUser from '../Services/SearchUser';
import './Home.css'
const Home = () => {
    return (
        <div className="home-content">
            <Header></Header>
            <SearchUser></SearchUser>
        </div>
    );
};

export default Home;