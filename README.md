# FinanceSim - Legacy Version
Finance simulator software

## Description of the repository
This was an old project written in .NET (frontend, CLI) and Python (backend) with MySQL DB support.

The goal of the new project was trying to update the stack of the application and change it to a modern one made by Flutter (frontend) and Go (backend) with MongoDB DB support. In particular, the legacy modernization had to be made in steps in order to simulate a large operation of the same category.

The steps of the project were the following.

Wave 1 - Users migration and cross management:
- Implement users management on the new backend
- Redirect old frontend users API calls to the new backend
- Migrate users credentials from old to new DB
- Update users id in the old DB in order to be able to make cross server joins

Wave 2 - Account migration and old server dismission:
- Implement account management on the new backend
- Redirect old fronent accounts API calls to the new backend
- Migrate accounts from old to new DB
- Dismiss old backend

Wave 3 - Frontend update:
- Implement user management on the new frontend
- Implement account management on the new frontend
- Implement CORS security
- Dismiss old frontend

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
This project won't be updated anymore. The project will be cloned in another repository in order to get more functions and upgrades. Moreover, the new version will have the old application components deleted.
