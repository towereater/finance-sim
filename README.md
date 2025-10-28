# FinanceSim
Finance simulator software

## Description of the repository
This project aims to reproduce a fully functional banking and trading app, while letting me strengthen my knowledge on REST APIs, DevOps and general good development practices.

## Architecture of the project
This project is split in four main parts:
- Mainframe: a banking simulator composed by many microservices (despite the name);
- BFF: a backend-for-frontend used to support the mobile app;
- Bank app [discontinued]: a mobile app used to perform basic banking operations;
- Xchanger: a finance simulator written in monolithic style to manage finance exchanges.

### Mainframe
The mainframe is composed by many microservices, namely:
- Security: performs api key checks, manages enabled banks institutes (ABIs) and bank admin users.
- User: manages bank common users, like those wanting to access to the mobile app.
- Account: contains the list of joins between users and all their accounts created by other services.
- Checking account: manages both checking accounts and payments. Payment processing is asynchronous and performed by a dedicated processor.
- Dossier: manages all financial transactions including dossiers. This service actively communicates with XChanger sevices since it's not independent due to the need of market data.

All mainframe services are written using Go and use Mongo for DB and Kafka for queue management.

### BFF
The BFF communicates with mainframe services to offer custom endpoints to the frontend.

The BFF is written using Go and does not need access to DB.

### Frontend
The fronend offers basic access to banking and finance operations.

The frontend is written using Dart/Flutter.

The frontend is currently discontinued and it will not work with the current state of the project. Future releases may fix this problem.

### XChanger
XChanger offers some trading services like stocks exchange. Differently from mainframe, XChanger is a monolithic structure with integrated security, db, jobs and so on.

XChanger is written using Java/Spring and uses Mongo for DB and Kafka for queue management.

### Other components
Other components of the project include Mongo and Kafka.

## Run the project
This section contains info about running the project and performing initial setup operations.

### Running the containers
To run the entire project there are three steps to follow:
- Run the setup.sh in the main folder to allign all project dependencies using
```bash
./setup.sh
```
This step can be performed just once since all vendor folders in the Go project will then be set up correctly.
- Run the docker compose command using
```bash
docker compose up
```
This step will launch all backend services of the project.
- [discontinued] Run the frontend from its folder using
```bash
cd bank_app/

flutter run
```

### Setup of the services
The following steps aim to setup the security services:
- Create an admin bank, i.e. a technical ABI code, on mainframe security service.
- Create a client bank, i.e. another ABI code, on mainframe security service. This request should contain an XChanger api key for later use.
- Create a client bank user, i.e. an admin user, on mainframe security service. This request should contain an admin apy key identifying the user itself.
- Create an admin bank, i.e. a technical ABI code, on XChanger security service. This will generate an admin api key to use for further operations.
- Create a client bank, i.e. another ABI code, on XChanger security service. This request must be made using the admin api key. This request should contain the same XChanger api key used before as the api token and the user bank api token used before as the external api token.

Service setup is now complete. The used api token may be then used to perform next operations like creating a bank common user, creating one or more accounts for it and using them.

## Next steps
Next service features may include, without any precise order:
- Security roles and resources management;
- Checking account monitor written in Rust;
- Stocks periodic coupons;
- Bank admin accounts selling suit;
- Bank admin accounts monitoring suit
- Frontend update.

Next devops features may include, without any precise order:
- Async logging using Kafka;
- Container orchestration using K8;
- Service monitoring using Prometheus;
- Data visualization using Grafana;
- Infrastructure setup using Terraform;
- Automatic release and testing using Jenkins;
- Multiple release environments.
