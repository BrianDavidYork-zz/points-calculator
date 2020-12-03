### SETUP INSTRUCTIONS ###

1) Download Golang -- Follow the instructions found here: https://golang.org/doc/install

2) Go will install in the home directory in a directory called Go/
- Inside Go/ you will find another directory called src/
- Download this project repository into src/ into a directory called PointsCalculator 
   
3) Use the cli to go to Go/src/PointsCalculator and execute "go install"

4) After the installation finishes, execute "go run main.go" 
   - If the api starts, you will see the following log message appear: "Starting Points api on port 8000" 
   
The PointsCalculator Api is now running on your localhost.

Next, choose your preferred method of api interaction. I recommend PostMan or cURL.

The API has 3 endpoints:

POST    /points/add

POST    /points/deduct

GET     /points/balance


The /add endoint takes a JSON body with 3 fields:

{
  "payer": string,
  "points": int,
  "date": string
}

The date field must be in the following format:  "mm/dd/yyyy hhPM"
All days, months, and hours must be two digits.  Febraury 5th 2020, 8 am would look like this: "02/05/2020 05AM"
An incorrectly formatted date will send back an error response from the server.

The /deduct endpoint takes a JSON body with 1 field:

{
  "points": int
}

#AUTHOR:
BRIAN YORK
