# Ultimate Crypto Currency Rate Provider
### By Volodymyr Kravchuk

This project is a powerful and scalable solution for fetching real-time cryptocurrency 
rates and dispatching to subscribed emails. Implemented using a microservice architecture of 4 services 
and the gRPC for connection between them. All services are independent and run in separate Docker containers.



## Running the application
Open `services` directory in the terminal and run:

```docker compose up```

Application will run on localhost:8080 by default.

## Requests

```
GET  -> http://localhost:8080/rate
POST -> http://localhost:8080/subscribe             
POST -> http://localhost:8080/sendEmails
```
