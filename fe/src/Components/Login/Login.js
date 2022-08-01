import React, { useContext, useState } from 'react';
import { initializeApp } from "firebase/app";
import { getAuth, signInWithPopup, GoogleAuthProvider} from "firebase/auth";
import firebaseConfig from './firebase.config';
import firebase from "firebase/compat/app";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faFacebook, faGoogle } from '@fortawesome/free-brands-svg-icons'
import { LoggedInContext } from '../../App';
import { useNavigate, useLocation } from 'react-router-dom';
import './Login.css'
import Header from '../Header/Header';
import Cookies from 'universal-cookie';


// Initialize Firebase
const app = initializeApp(firebaseConfig);
// Initialize Firebase Authentication and get a reference to the service
const auth = getAuth(app);
const customAuth = getAuth();

const Login = () => {
    const [loggedInUser, setLoggedInUser] = useContext(LoggedInContext);
    const navigate = useNavigate();
    const location = useLocation();
    let { from } = location.state || { from: { pathname: "/" } };

    //Toggle signIn/signup
    const [newUser, setNewUser] = useState(false)
    const accountToggle = () => {
        if (newUser === false) {
            setNewUser(true);
        }
        else {
            setNewUser(false);
        }
    }

    const [user, setUser] = useState({
        isSignedIn: false,
        name: '',
        email: '',
        password: '',
        confirmPassword: '',
        error: '',
        success: false,
        token:''
    });


    //Google signin
    const provider = new GoogleAuthProvider();
    const handleGoogleSignIn = () => {
        signInWithPopup(auth, provider)
            .then((res) => {
                console.log(res.user)
                const { displayName, email } = res.user
                const newUserInfo = { ...user };
                newUserInfo.error = '';
                newUserInfo.success = true;
                newUserInfo.isSignedIn = true;
                newUserInfo.name = displayName;
                newUserInfo.email = email;
                setUser(newUserInfo);
                setLoggedInUser(newUserInfo);
                navigate(from);
            }).catch((error) => {
                console.log(error);
            });
    }
    //Facebook signIn
    const handleFbSignIn = () => {
        console.log(user)
        const fb = ({
            signinMethod: "Facebook",
        })
            const url = 'http://localhost:8000/facebook/login';
            fetch(url, {
                method: 'POST',
                headers: {
                    'content-type': 'application/json'
                },
                body: JSON.stringify(fb)
            })
                .then(response => response.json())
                .then(data => {
                    console.log('Success:', data);
                    const { name,token } = data
                    const newUserInfo = { ...user };
                    newUserInfo.error = '';
                    newUserInfo.success = true;
                    newUserInfo.isSignedIn = true;
                    newUserInfo.name = name;
                    newUserInfo.token= token;
                    setUser(newUserInfo);
                    setLoggedInUser(newUserInfo);
                    navigate(from);
                })
                .catch((error) => {
                    const newUserInfo = { ...user };
                    newUserInfo.error = error.message;
                    newUserInfo.success = false;
                    setUser(newUserInfo);
                });
    }
    //get value from form field
    const handleBlur = (e) => {
        let isFormValid;
        if (e.target.name === 'name') {
            isFormValid = e.target.value.length > 5;
            const newUserInfo = { ...user };
            newUserInfo.error = 'Name length is short';
            setUser(newUserInfo);
        }
        if (e.target.name === 'email') {
            isFormValid = /\S+@\S+\.\S+/.test(e.target.value);
            const newUserInfo = { ...user };
            newUserInfo.error = 'email not valid';
            setUser(newUserInfo);
        }
        if (e.target.name === 'password') {
            const isPasswordValid = e.target.value.length > 5;
            const hasNumber = /\d{1}/.test(e.target.value);
            isFormValid = isPasswordValid;
            const newUserInfo = { ...user };
            newUserInfo.error = 'Password length is short';
            setUser(newUserInfo);
        }
        if (e.target.name === 'confirmPassword') {
            const isPasswordValid = e.target.value.length > 5;
            isFormValid = isPasswordValid;
            const newUserInfo = { ...user };
            if (newUserInfo.password === e.target.value) {
                isFormValid = true;
                newUserInfo.error = 'Password don"t match';
                setUser(newUserInfo);
            }

        }
        if (isFormValid) {
            const newUserInfo = { ...user };
            newUserInfo.error = '';
            newUserInfo[e.target.name] = e.target.value;
            setUser(newUserInfo);
        }
    }

    const handleSubmit = (e) => {
        //Error message if the form is not complete
        if (newUser) {
            if (user.email === '' || user.password === '' || user.name === '' || user.confirmPassword === '') {
                const newUserInfo = { ...user };
                newUserInfo.error = 'Please complete the form';
                setUser(newUserInfo);
            }
            if (user.password !== user.confirmPassword) {
                const newUserInfo = { ...user };
                newUserInfo.error = 'Password don"t match';
                setUser(newUserInfo);
            }
        }
        else {
            if (user.email === '' || user.password === '') {
                const newUserInfo = { ...user };
                newUserInfo.error = 'Please complete the form';
                setUser(newUserInfo);
            }
        }

        //Sign up 
        if (newUser && user.email && user.password && user.confirmPassword) {
            if (user.password === user.confirmPassword) {
                console.log(user);
                const url = 'http://localhost:8000/register';
                fetch(url, {
                    method: 'POST',
                    headers: {
                        'content-type': 'application/json'
                    },
                    body: JSON.stringify(user)
                })
                    .then(response => response.json())
                    .then(data => {
                        console.log('Success:', data);
                        const { name } = data
                        const newUserInfo = { ...user };
                        newUserInfo.error = '';
                        newUserInfo.success = true;
                        newUserInfo.isSignedIn = true;
                        newUserInfo.name = name;
                        setUser(newUserInfo);
                        setLoggedInUser(newUserInfo);
                        navigate(from);
                    })
                    .catch((error) => {
                        const newUserInfo = { ...user };
                        newUserInfo.error = error.message;
                        newUserInfo.success = false;
                        setUser(newUserInfo);
                    });
            }
        }

        //Sign In
        if (!newUser && user.email && user.password) {
            console.log(user);
            const url = 'http://localhost:8000/login';
            fetch(url, {
                method: 'POST',
                headers: {                   
                    'content-type': 'application/json'
                },
                body: JSON.stringify(user)
            })
                .then(response => response.json())
                .then(data => {
                    console.log('Success:', data);
                    const { name,token } = data
                    const newUserInfo = { ...user };
                    newUserInfo.error = '';
                    newUserInfo.success = true;
                    newUserInfo.isSignedIn = true;
                    newUserInfo.name = name;
                    sessionStorage.setItem("session_token", token);
//                     const cookies = new Cookies();
// cookies.set('session_token', token);
                    setUser(newUserInfo);
                    setLoggedInUser(newUserInfo);
                    navigate(from);
                })
                .catch((error) => {
                    const newUserInfo = { ...user };
                    newUserInfo.error = error.message;
                    newUserInfo.success = false;
                    setUser(newUserInfo);
                });
        }
        e.preventDefault();
    }

    return (
        <div>
            <Header></Header>
            <div className="container">
                <div className="row text-center justify-content-center">
                    <div className="col-md-8">
                        <div className="form-content">
                            {
                                newUser ? <p>Already have an account? <span onClick={accountToggle}>Sign In</span></p>
                                    : <p>Don't have an account? <span onClick={accountToggle}>Create account</span></p>
                            }

                            <form>
                                <div className="mb-3">
                                    {newUser && <input type="text" name="name" className="form-control" placeholder="Name" required onBlur={handleBlur} />}
                                </div>
                                <div className="mb-3">
                                    <input type="email" name="email" className="form-control" placeholder="Email" required onBlur={handleBlur} />
                                </div>
                                <div className="mb-3">
                                    <input type="password" name="password" className="form-control" placeholder="Password" required onBlur={handleBlur} />
                                </div>
                                {
                                    newUser ? <div className="mb-3">
                                        <input type="password" name="confirmPassword" className="form-control" placeholder="Confirm Password" required onBlur={handleBlur} />
                                    </div> : ''
                                }

                                <button className="btn btn-primary" onClick={handleSubmit}>{newUser ? 'Sign Up' : 'Sign In'}</button>
                            </form>
                            <p className="text-danger">{user.error}</p>
                            {user.success && <p className="text-success">User {newUser ? 'Created' : 'Logged In'} successfully</p>}
                        </div>
                    </div>
                    <div className="col-md-6">
                        <div className="login-form text-center">
                            <h4>Don't have an account? Sign in with Google/Facebook</h4>
                            <button className="btn btn-outline-danger" onClick={handleGoogleSignIn}>
                                <FontAwesomeIcon icon={faGoogle} /> Continue with Google
                            </button>
                            <br />
                            <button className="btn btn-outline-info mt-3" onClick={handleFbSignIn}>
                                <FontAwesomeIcon icon={faFacebook} /> Continue with Facebook
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Login;