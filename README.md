
About project:

This project is all about accessing the videos by the authorised user.
Here I have used JWT for the authentication of the user.

HLD included with name "HLD.txt"


Requirements:

1. Install latest GO version into your system.

2. Initialise the project by : "go mod init skid_go"

3. Install dependencies by: "go mod tidy"

4. To run this project: "go run main.go"


Tips to access the videos:

1. Import the sql file into your local DBMS.
2. There is a table named "users", it contains "username" and "password".
3. So the users registered into the users table are authorised to access the videos and its content.
4. For videos I have took the random url form the youtube videos.