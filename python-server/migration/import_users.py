import sys
import sqlite3

def get_connection(db_file):
    conn = None

    try:
        conn = sqlite3.connect(db_file)
        return conn
    except sqlite3.Error as e:
        print(f"An error has occurred while establishing a connection: {e}")
    
    return conn

def create_new_table(conn):
    comm = '''
        CREATE TABLE IF NOT EXISTS account_wallets_new (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            account_id TEXT NOT NULL,
            wallet_id INTEGER NOT NULL,
            FOREIGN KEY (wallet_id) REFERENCES wallets (id)
        );
    '''

    try:
        # Runs the command and fetches the response in case of a query
        cur = conn.cursor()
        cur.execute(comm)

        return 0
    # Generic db error
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

        return 1

def select_account_wallets(conn, account_id):
    comm = '''
        SELECT wallet_id
        FROM account_wallets
        WHERE account_id = ?
    '''

    try:
        # Runs the command and fetches the response in case of a query
        cur = conn.cursor()
        cur.execute(comm, (account_id,))
        return cur.fetchall()
    # Generic db error
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

def insert_account_wallets(conn, account_id, wallet_id):
    comm = '''
        INSERT INTO account_wallets_new (account_id, wallet_id)
        VALUES (?, ?)
    '''

    try:
        # Runs the command and fetches the response in case of a query
        cur = conn.cursor()
        cur.execute(comm, (account_id, wallet_id))
        conn.commit()

        return 0
    # Generic db error
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

        return 1

def main():
    db_conn = sys.argv[1]
    dump_path = sys.argv[2]

    #Logging
    print("Running with DB connection {0} and dump path {1}".format(db_conn, dump_path))

    #Get DB connection
    conn = get_connection(db_conn)

    #Update DB constraints
    res = create_new_table(conn)
    if res > 0:
        return

    f = open(dump_path, "r")
    n = 34

    #Get id associations
    row = f.read(n)
    while len(row) == n:
        old_id = ''.join(row[:9])
        new_id = ''.join(row[9:-1])

        account_wallets = select_account_wallets(conn, old_id)
        if account_wallets is None:
            continue

        for i in range(len(account_wallets)):
            res = insert_account_wallets(conn, new_id, account_wallets[i][0])
            if res > 0:
                return
        
        row = f.read(n)

    f.close()
    conn.close()

    if len(row) > 0:
        print("Something weird happened! Only {0} bytes were read".format(len(row)))
    
# EXECUTE MAIN FUNCTION
if __name__ == "__main__":
    main()