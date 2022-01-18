from enum import IntEnum
from socket import *
from cryptography.fernet import Fernet
import json
import os

from financelib.db_utils import insert_account_wallet, insert_wallet
from financelib.db_utils import select_wallets_by_account
from financelib.db_utils import insert_account, select_account_by_username

# Data content of a json packet
REQUEST_TOKEN = "requestToken"
RESPONSE_TOKEN = "responseToken"
ACCESS_TOKEN = "accessToken"
PAYLOAD = "payload"

# List of possible requests
class RequestToken(IntEnum):
    CreateAccount = 0
    Login = 1
    CreateWallet = 21
    GetWallet = 22
    DeleteWallet = 23

# Active connections of the form (key: accessToken, value: accountId)
connections = {}

# Database path
dir = os.path.dirname(__file__)
db_path = os.path.join(dir, "sqlite", "db", "bankdata.db")

# Account creation procedure
# Adds a new entry in the accounts table corresponding to a username-password
# pair if it does not exist.
def create_account(payload):
    # Gets the new account data from the payload
    user = payload["user"]
    password = payload["password"]

    # Username cannot be null
    if user is None:
        return {
            REQUEST_TOKEN: RequestToken.CreateAccount,
            RESPONSE_TOKEN: 0,
            PAYLOAD: {
                'message': "Username cannot be null."
            },
        }
    # Password cannot be null
    elif password is None:
        return {
            REQUEST_TOKEN: RequestToken.CreateAccount,
            RESPONSE_TOKEN: 0,
            PAYLOAD: {
                'message': "Password cannot be null."
            },
        }
    # Username already present in the db
    elif select_account_by_username(db_path, user) is not None:
        return {
            REQUEST_TOKEN: RequestToken.CreateAccount,
            RESPONSE_TOKEN: 0,
            PAYLOAD: {
                'message': "Username already exists."
            },
        }
    # Username-password pair is inserted in the db
    else:
        insert_account(db_path, user, password)

        return {
            REQUEST_TOKEN: RequestToken.CreateAccount,
            RESPONSE_TOKEN: 1,
            PAYLOAD: {
                'message': "Account successfully created."
            },
        }

# Login procedure
# Searches for the given username-password pair in the db. If it finds it,
# then it returns a new accessToken and saves it in the connections list.
def login(payload):
    # Gets the login data from the payload
    user = payload["user"]
    password = payload["password"]

    # Searches for the login data in the db
    accounts = select_account_by_username(db_path, user)
    # If the username has been found
    if accounts is not None:
        # If the password matches
        if accounts["password"] == password:
            # Creates a new accessToken and saves it in the connections list
            accessToken = Fernet.generate_key().decode()
            connections[accessToken] = accounts["id"]

            return {
                REQUEST_TOKEN: RequestToken.Login,
                RESPONSE_TOKEN: 1,
                PAYLOAD: {
                    'accessToken': accessToken,
                    'message': "Login successful."
                },
            }
        # If the password does not match
        else:
            return {
                REQUEST_TOKEN: RequestToken.Login,
                RESPONSE_TOKEN: 0,
                PAYLOAD: {
                    'message': "Password is not correct.",
                },
            }
    # If the username has not been found
    else:
        return {
            REQUEST_TOKEN: RequestToken.Login,
            RESPONSE_TOKEN: 0,
            PAYLOAD: {
                'message': "Username does not exists.",
            },
        }

# Get wallet procedure
# Returns all the wallets associated to the given user.
def get_wallet(payload, accessToken):
    account_id = connections[accessToken]

    return {
        REQUEST_TOKEN: RequestToken.GetWallet,
        RESPONSE_TOKEN: 1,
        PAYLOAD: {
            'wallets': [ {
                'iban': wallet["iban"],
                'cash': wallet["cash"],
            } for wallet in select_wallets_by_account(db_path, account_id) ],
            'message': "Wallets associated to given account."
        },
    }

# Create wallet procedure
# Adds a new entry in the wallets table corresponding to a iban-cash pair.
#TODO: FIX IBAN RANDOMIZATION
def create_wallet(payload, accessToken):
    account_id = connections[accessToken]
    iban = Fernet.generate_key().decode()
    cash = 0

    wallet = insert_wallet(db_path, iban, cash)
    insert_account_wallet(db_path, account_id, wallet["id"])

    return {
        REQUEST_TOKEN: RequestToken.CreateWallet,
        RESPONSE_TOKEN: 1,
        PAYLOAD: {
            'wallets': [ {
                'iban': iban,
                'cash': cash,
            } ],
            'message': "New wallet associated to given account created."
        },
    }

def handle_request(request):
    # Decomposes the json into parts
    requestToken = request[REQUEST_TOKEN]
    accessToken = request[ACCESS_TOKEN]
    payload = request[PAYLOAD]

    # CREATE_ACCOUNT requested
    if requestToken == int(RequestToken.CreateAccount):
        print("CREATE_ACCOUNT requested")
        return create_account(payload)
    # LOGIN requested
    elif requestToken == int(RequestToken.Login):
        print("LOGIN requested")
        return login(payload)
    # The other procedures require a valid accessToken, so its existence is checked
    elif accessToken is None or connections[accessToken] is None:
        return {
                REQUEST_TOKEN: requestToken,
                RESPONSE_TOKEN: 0,
                PAYLOAD: {
                    'message': "Access token is not valid."
                },
        }
    # Checking more possible requests since accessToken is valid
    else:
        # GET_WALLET requested
        if requestToken == int(RequestToken.GetWallet):
            print("GET_WALLET requested")
            return get_wallet(payload, accessToken)
        # CREATE_WALLET requested
        elif requestToken == int(RequestToken.CreateWallet):
            print("CREATE_WALLET requested")
            return create_wallet(payload, accessToken)
        # UNKNOWN request
        else:
            print("UNKNOWN request")
            return {
                REQUEST_TOKEN: requestToken,
                RESPONSE_TOKEN: 0,
                PAYLOAD: {
                    'message': "Invalid request. The server cannot correctly handle it.",
                },
            }

def run_server(server_port):
    # Creates a TCP socket with a maximum of 5 active connections
    server_socket = socket(AF_INET, SOCK_STREAM)
    server_socket.bind(('', server_port))
    server_socket.listen(5)
    print ("The server is ready to receive")

    while 1:
        # Server waits for incoming connections on accept()
        # For incoming requests, a new socket is created on return
        connection_socket, addr = server_socket.accept()
        print("Connected to a client, waiting for requests")

        # Receives some data on the newly established connectionSocket
        data = connection_socket.recv(1024).decode("utf-8")

        # Analyses the request made by the client and sets up the response
        print("Request received, elaborating it")
        request = json.loads(data)
        response = handle_request(request)
        respJson = json.dumps(response)

        # Sends back the modified string to the client
        print("Sending the response to the client")
        connection_socket.send(respJson.encode("utf-8"))

        # Closes the connection with the client
        print("Closing the connection with the client")
        connection_socket.close()
    
def main():
    run_server(11000)

if __name__ == "__main__":
    main()