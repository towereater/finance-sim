from enum import IntEnum
from socket import *
from cryptography.fernet import Fernet
import json

from financelib.wallet import Wallet
from financelib.useraccount import UserAccount

# Data content of a json packet
REQUEST_TOKEN = "requestToken"
RESPONSE_TOKEN = "responseToken"
ACCESS_TOKEN = "accessToken"
PAYLOAD = "payload"

# List of possible requests
class RequestToken(IntEnum):
    Login = 0
    Wallet = 1

# List of user accounts
wallet1 = Wallet()
wallet2 = Wallet()
wallet3 = Wallet()
user1 = UserAccount("andnic", "pass")
user1.wallets = [ wallet1, wallet2 ]
user2 = UserAccount("frau", "xdxd")
user2.wallets = [ wallet2, wallet3 ]
user_accounts = [
    user1,
    user2,
]

# Active connections of the form (key: accessToken, value: userAccountIndex)
connections = {}

# Server port number
server_port = 11000
# Creates TCP socket
server_socket = socket(AF_INET, SOCK_STREAM)
# Binds socket to the chosen local port
server_socket.bind(('', server_port))
# Puts socket in passive mode
server_socket.listen(1)
print ("The server is ready to receive")

while 1:
    # Server waits for incoming connections on accept()
    # For incoming requests, a new socket is created on return
    connection_socket, addr = server_socket.accept()
    print("Connected to a client, waiting for requests")
    # Receives sentence on newly established connectionSocket
    data = connection_socket.recv(1024).decode("utf-8")
    print("Request received, elaborating it")

    # Prepares the object to analyse
    reqJson = json.loads(data)

    print(reqJson)

    requestToken = reqJson[REQUEST_TOKEN]
    accessToken = reqJson[ACCESS_TOKEN]
    payload = reqJson[PAYLOAD]

    # If login procedure was initiated
    if requestToken == int(RequestToken.Login):
        print("LOGIN requested")

        user = payload["user"]
        password = payload["password"]

        # Verifies the validity of the received data
        #TODO: DATABASE SET
        valid_cred = False
        user_index = -1
        for idx,item in enumerate(user_accounts):
            if item.match_credentials(user, password):
                user_index = idx
                valid_cred = True
                break

        # The user-pass pair has been found
        if valid_cred:
            key = Fernet.generate_key().decode()
            connections[key] = user_index
            response = {
                REQUEST_TOKEN: requestToken,
                RESPONSE_TOKEN: 1,
                PAYLOAD: {
                    'accessToken': key,
                },
            }
        # The user-pass pair has not been found
        else:
            response = {
                REQUEST_TOKEN: requestToken,
                RESPONSE_TOKEN: 0,
                PAYLOAD: {
                    'message': "The combination username and password was not correct.",
                },
            }
    elif requestToken == int(RequestToken.Wallet):
        print("WALLET requested")

        index = connections[accessToken]

        # Checking the access token
        if index == None:
            response = {
                REQUEST_TOKEN: requestToken,
                RESPONSE_TOKEN: 0,
                PAYLOAD: {
                    'message': "Invalid access token has been provided",
                },
            }
        else:
            response = {
                REQUEST_TOKEN: requestToken,
                RESPONSE_TOKEN: 1,
                PAYLOAD: {
                    'wallets': [ wallet.to_json() for wallet in user_accounts[index].wallets ],
                },
            }
    else:
        print("UNKNOWN request was made")

        response = {
            REQUEST_TOKEN: requestToken,
            RESPONSE_TOKEN: 0,
            PAYLOAD: {
                'message': "Invalid request token. The server cannot correctly handle it.",
            },
        }

    respJson = json.dumps(response)

    # send back modified string to client
    print("Sending a response to the client")
    connection_socket.send(respJson.encode("utf-8"))
    print("Closing the connection with the client")