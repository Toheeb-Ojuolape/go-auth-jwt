## Go-authentication using gin with JWT, ENV and PostgreSQL set up
You can easily get started with authenticating your users using signup, login and fetch user data using this Github project. Just fork it to get started. 


## Setting up Hot Reload in Go
You can use go daemon to setup hot reload so you don't need to restart your project everytime you push an update:

-[x] 1) In your command line, run : go install github.com/joho/godotenv
-[x] 2) then run compiledaemon --command="./go-gin-auth"  (replace go-gin-auth with the name of your own project)



## Available endpoints
[x] /signup - takes in email, password, username, firstName, lastName, phone (you can always add more)
[x] /login - takes in email and password and returns a jwt token and user data
-[x] /user - fetches all the user's data from the database
-[x] /forgot-password - sends an otp for user to reset their password (coming soon)
-[x] /reset-password - enables a user to change their password (coming soon)
-[x] /google-auth - enables a user to login with google (coming soon)


## Database setup
Project uses postgreSQL but you can use any database of your choice, but just make sure to follow best practice by specifying your connection string in your .env. Sample DB string for Postgres:
"host=localhost user=postgres password=Password dbname=go-gin-auth port=5432 sslmode=disable"
