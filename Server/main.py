from enum import Enum
from socket import *
from cryptography.fernet import Fernet
import json
from Server.financelib.wallet import Wallet

import financelib.useraccount as ua

# Data content of a json packet
REQUEST_TOKEN = "requestToken"
RESPONSE_TOKEN = "responseToken"
PAYLOAD = "payload"

# List of possible requests
class RequestToken(Enum):
    LOGIN = 0
    WALLET = 1

# List of user accounts
wallet1 = Wallet()
wallet2 = Wallet()
wallet3 = Wallet()
user1 = ua.UserAccount("andnic", "pass")
user1.wallets = [ wallet1, wallet2 ]
user2 = ua.UserAccount("frau", "xdxd")
user2.wallets = [ wallet2, wallet3 ]
user_accounts = [
    user1,
    user2,
]

# Active connections
access_tokens = {}

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
    print("Connected to a client, waiting for data")
    # Receives sentence on newly established connectionSocket
    data = connection_socket.recv(1024).decode("utf-8")
    print("Data received, elaborating it")

    # Prepares the object to analyse
    log_json = json.loads(data)
    requestToken = log_json[REQUEST_TOKEN]
    payload = log_json[PAYLOAD]

    # If login procedure was initiated
    if requestToken == RequestToken.LOGIN:
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

        # Prepares the correct response depending on the result
        if valid_cred:
            key = Fernet.generate_key().decode()
            access_tokens[key].append(user_index)
            response = {
                REQUEST_TOKEN: requestToken,
                RESPONSE_TOKEN: 1,
                PAYLOAD: {
                    'accessToken': key,
                },
            }
        else:
            response = {
                REQUEST_TOKEN: requestToken,
                RESPONSE_TOKEN: 0,
                PAYLOAD: {
                    'message': "The combination username and password was not correct.",
                },
            }
    elif requestToken == RequestToken.WALLET:
        print("WALLET requested")

        key = payload["accessToken"]
        index = access_tokens[key]

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