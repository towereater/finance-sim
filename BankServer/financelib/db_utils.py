import sqlite3

# Gets the connection with a given db file and returns it
def get_connection(db_file):
    conn = None

    try:
        conn = sqlite3.connect(db_file)
    except sqlite3.Error as e:
        print(f"An error has occurred while connecting to the db: {e}")
    
    return conn

# Executes a command on a given db file
def execute_command(db_file, command, params):
    # Opens the connection with the db file
    conn = get_connection(db_file)

    try:
        # Runs the command and fetches the response in case of a query
        cur = conn.cursor()
        cur.execute(command, params)
        queryResponse = cur.fetchall()

        # Closes the connection after the commitment
        conn.commit()
        conn.close()

        return queryResponse
    # Generic db error
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

        return None

# Account selection using username value
def select_account_by_username(db_file, username):
    comm = '''
        SELECT *
        FROM accounts
        WHERE username = ?
    '''
    return execute_command(db_file, comm, (username,))

# Wallet selection using account id
def select_wallets_by_account(db_file, account_id):
    comm = '''
        SELECT W.iban, W.cash
        FROM account_wallets AW
        JOIN wallets W
        ON W.id = AW.wallet_id
        WHERE account_id = ?
    '''
    return execute_command(db_file, comm, (account_id,))

# Account insertion
def insert_account(db_file, username, password):
    comm = '''
        INSERT INTO accounts (username, password)
        VALUES (?, ?)
    '''
    return execute_command(db_file, comm, (username, password))

# Wallet insertion
def insert_wallet(db_file, iban, cash=0):
    comm = '''
        INSERT INTO wallets (iban, cash)
        VALUES (?, ?)
    '''
    return execute_command(db_file, comm, (iban, cash))

# Account-wallet pair insertion
def insert_account_wallet(db_file, account_id, wallet_id):
    comm = '''
        INSERT INTO account_wallets (account_id, wallet_id)
        VALUES (?, ?)
    '''
    return execute_command(db_file, comm, (account_id, wallet_id))