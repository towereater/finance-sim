# FinanceSim
Finance simulator software

## Description of the repository
...

## Architecture of the project
The project is divided in backend, BFF (backend-for-frontend) and frontend. The choice of create the extra bff was due in order to increase the project complexity.

### Backend
The backend is written in Go and uses MongoDB DB. It's made by the users and accounts microservices which also interact among themselves. The microservices have their own configuration files which can be edited, if needed.

In order to run the users microservice, navigate to the corresponding folder (server/mainframe/user) and run the commands:
```bash
go build -o ./bin/user .

./bin/user
```

In order to run the accounts microservice, navigate to the corresponding folder (server/mainframe/account) and run the commands:
```bash
go build -o ./bin/account .

./bin/account
```

### BFF
The BFF is written in Go and communicates with the backend microservices. The BFF have its own configuration files which can be edited, if needed.

In order to run the bff, navigate to the corresponding folder (server/bff) and run the commands:
```bash
go build -o ./bin/bff .

./bin/bff
```

### Frontend
The frontend is written in Flutter. The frontend have its own configuration files which can be edited, if needed.

In order to run the frontend, navigate to the corresponding folder (bank_app) and run the command:
```bash
flutter run
```

## Next steps
Stay tuned.
