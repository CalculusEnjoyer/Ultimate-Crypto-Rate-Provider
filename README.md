# Ultimate Crypto Currency Rate Provider
### By Volodymyr Kravchuk

This project is a powerful and scalable solution for fetching real-time cryptocurrency 
rates and dispatching to subscribed emails. Implemented using a microservice architecture of 4 services 
and the gRPC for connection between them. All services are independent and run in separate Docker containers.

<br />
<img src="https://github.com/CalculusEnjoyer/Ultimate-Crypto-Rate-Provider/blob/main/micro.png">
<br />

## Running the application
Open `services` directory in the terminal and run:

```docker compose up```

Application will run on localhost:8080 by default.
## Setting up sender email
Test credentials for quick testing are already set, so you do not have to do anything 
in order to set up it. But if you want to change sender email, you can do this in `email` service
in `.env` file (path: `services/email/.env`)
## Requests

```
GET  -> http://localhost:8080/rate
POST -> http://localhost:8080/subscribe             
POST -> http://localhost:8080/sendEmails
```
