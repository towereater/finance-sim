from socket import *
import json

import financelib.useraccount as ua

# List of possible bank accounts
user_accounts = [
    ua.UserAccount("andnic", "pass"),
    ua.UserAccount("frau", "xdxd"),
]

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

    # Extracts the user account data
    logJson = json.loads(data)
    usr = logJson["usr"]
    pwd = logJson["pwd"]

    # Verifies the validity of the received data
    valid_cred = False
    user_account = None
    for usr_acc in user_accounts:
        if usr_acc.match_credentials(usr, pwd):
            user_account = usr_acc
            valid_cred = True

    # Prepares the correct response depending on the result
    if valid_cred:
        response = {
            'valid': 1,
            'data': {
                'user': user_account.user,
                'iban': user_account.bankAcc.iban,
                'cash': user_account.bankAcc.cash,
            }
        }
    else:
        response = {
            'valid': 0
        }
    respJson = json.dumps(response)

    # send back modified string to client
    print("Sending a response to the client")
    connection_socket.send(respJson.encode("utf-8"))
    print("Closing the connection with the client")