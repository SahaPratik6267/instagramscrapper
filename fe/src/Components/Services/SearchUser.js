import React, { useContext, useState } from 'react';
import { useForm } from "react-hook-form";
import Cookies from 'universal-cookie';
import { LoggedInContext } from "../../App";

const SearchUser = () => {
    const [loggedInUser, setLoggedInUser] = useContext(LoggedInContext);
    const { register, handleSubmit } = useForm();

    const [scrappedUser, setScrappedUser] = useState([]);
    //Send data to server
    const onSubmit = data => {
        if (loggedInUser.isSignedIn === true) {
            const eventData = {
                userName: data.user_name,
                token: sessionStorage.getItem("session_token")
                
            }
            const url = 'http://localhost:8000/twitter';
            fetch(url, {
                method: 'POST',
                headers: {
                    'content-type': 'application/json'
                },
                body: JSON.stringify(eventData)
            })
                .then(response => response.json())
                .then(data => {
                    console.log('Success:', data);
                    return setScrappedUser(data);
                })
                .catch((error) => {
                    window.alert("you need to login again. Go to home and click login")
                    sessionStorage.removeItem("session_token")
                    window.location.reload(true);
                    console.error('Error:', error);
                });
        } else {
            console.log("You need to SignIn!");
        }

    };
    return (
        <div className="container-fluid row">
            <div className="col-md-12">
                <div className="container mt-5">
                    <form onSubmit={handleSubmit(onSubmit)} className="text-center product-form">
                        <h2 className='text-info fw-bold'>Search User</h2>
                        <input {...register("user_name")} type="text" placeholder={loggedInUser.isSignedIn ? 'Search User' : 'Sign in to search'} className="form-control" disabled={loggedInUser.isSignedIn ? '' : 'true'} />
                        <br />
                        <input type="submit" className="btn btn-primary mt-4" disabled={loggedInUser.isSignedIn ? '' : 'true'} />
                    </form>
                </div>
            </div>
            <div className="col-md-12 text-center">
                <div className="container mt-5">
                    <div className="row justify-content-center mt-5">
                        <h2 className='text-success fw-bold'>User Details</h2>
                        <div className='col-md-6 d-flex justify-content-center'>
                            <div className="card" style={{ width: 18 + 'rem' }}>
                                <img className="card-img-top" src={scrappedUser.Avatar ? scrappedUser.Avatar : "https://www.w3schools.com/howto/img_avatar.png"} alt="img"></img>
                                <div className="card-body">
                                    <h5 className="card-title">Name: {loggedInUser.isSignedIn ? scrappedUser.Name : 'Jon Doe'}</h5>
                                    <p className="card-text">{loggedInUser.isSignedIn ? scrappedUser.biography : ''}</p>
                                </div>
                                <ul className="list-group list-group-flush">
                                    <li className="list-group-item">Followers: {loggedInUser.isSignedIn ? scrappedUser.FollowersCount : '0'}</li>
                                    <li className="list-group-item">Like: {loggedInUser.isSignedIn ? scrappedUser.LikesCount : '0'}</li>
                                    <li className="list-group-item">Tweets: {loggedInUser.isSignedIn ? scrappedUser.TweetsCount : '0'}</li>
                                </ul>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default SearchUser;