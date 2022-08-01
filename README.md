# twitterScrapper

This is a simple Scrapper app that scrapes information from twitter:

SETUP: The project is completed using docer. 

Step1: make sure you have docker installed in your system
step2: clone the project and then run the command inside the directory of docker-compose.yaml 
--docker compose up
PS: Sometimes you might get error that database connection unsuccessful for the first time running just run the command again and it should be fine.
//TODO: Make sure the dependency in docker compose is fine so that the above error is fixed

Project details:

Frontend: I have used React for frontend which you can access from localhost:3000

Since I am not a pro frontend person you might find some code like a noob but its doing the job. 

Flow: once you go to homepage you will see the data scrapper field is disabled. You need to login first.
=> Click on login button and create a user first (you can use login with google option too which is working)
=> Once you logged in your session will be active for 2 mins. Withing this time you can put any twitter username in the field and see scrapped values.
=> Once the session is finish you will be logged out.


BACKEND: You can find the API docs of Swagger at localhost:8000/docs
You will see the login,register and scrapping url in the docs.
TODO: add more description in the docs. Right now its just basic information

NOTE: I have implemented the login with google and Facebook in backend too. Just go to this link localhost:8000/google/login and localhost:8000/facebook/login
then follow the process and you will see the name associated with the profile is showing. Its not handled in swagger because of time shortage. which is why same
functionality was implemented in CLient side with react.

Database: I used postgresql db to store signup user data.

Things I need to change: 
Right now the authentication is not handled via header and since i am not very good with react, I need to have some time to work with authentication/cookie handling.  
So i need to implement it later.    

The base url and other environment values are stored in code but it should be handled via environment file.  

Need to change authentication token to jwt token.
