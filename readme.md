# City Parking 

WIP  - it's still Work in progress 

This is just a concept of backend for the city parking application.
The idea was to create an app with using native Golang packages and less external packages

## System Design
Miro Board https://miro.com/app/board/uXjVOZWLF4I=/?invite_link_id=106276965135

![](.idea/help/Global-Parking-System-Design.jpg)
This repo is all about Backend of the application
in ideal world, the design should be cloud agnostic, but we are all humans and can make mistakes. My first mistake is to design this with AWS :)

I'd like to host the Backend at AWS. it will be set of lumbda functions with AWS Api Gateway in front of this.  
### Benefits:
 - Scaling independently. We don't want to scale up Administrators API endpoints to scale up with main users api. 
 - Geo distribution
 - Api Gateway works like load balancer and manage user's authentication.
 - Our service needs to be available, and we are ok if the data won't be fresh. 
Example: Inspector API checks if the parking spot is paid or not. If it' paid it will be paid at least for an hour. this information can be changed letter (another person paid/extend the booking of the parking spot) but anyway this hour is already paid. we need this information available for inspector, then we can refresh it with new updated hours
 - Separation of payment methods to a separate lumbda function helps with scalability and performance. If the payment system is down, we don't want to slow down or main application. Let's implement circuit breaker 
### Database:
 - NoSQL DB would work well for this project at this stage, but if we want to add user stats and manage roles, relation DB would be better option. 
 - Database scheme and relation is well represented at `entity` folder

 
## LOCAL
The app do not require yet any DB connection 

### run on local: 
- If you are using GoLand or IntelliJ, you can just start the app from the `main.go`
![](.idea/help/go-build.png)
- Running terminal command  `go run main.go` would also make it work
### test on local:
see `HTTPrequest` folder for examples of requests
![](.idea/help/httprequests.png)

### Potential problems
if something wrong, run `go mod tidy` It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules that don’t provide any relevant packages