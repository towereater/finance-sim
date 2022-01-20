# FinanceSim
Finance simulator software

## Server
In order to run the server, navigate to the corresponding folder and run the command:
```bash
python3 main.py
```

The application will then start to listen to incoming connections on port 11000 and will address their requests depending on the particular flag in the JSON files exchanged with them.

## Client
In order to run the client, navigate to the corresponding folder and run the command:
```bash
dotnet run
```

The one developed is a console app that allow simple requests to the server Not all the ones displayed are available to the user and working as expected.

### Server TODO features
- Fix "AddWallet" function
- Async server update
- Add initial config file
- Delete old connections from the active connections list during server activity to save memory

### Client TODO features
- Fix client features
- General refactoring
- UI update
- Visual app replacing console one