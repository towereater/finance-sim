import sqlite3

def get_connection(db_file):
    conn = None

    try:
        conn = sqlite3.connect(db_file)
        return conn
    except sqlite3.Error as e:
        print(f"An error has occurred while establishing a connection: {e}")
    
    return conn

def select_wallets(conn):
    comm = '''SELECT w.id, w.iban, aw.account_id, w.cash FROM wallets w
        JOIN account_wallets_new aw ON aw.wallet_id = w.id
        ORDER BY w.id ASC'''

    try:
        # Runs the command and fetches the response in case of a query
        cur = conn.cursor()
        cur.execute(comm)
        queryResponse = cur.fetchall()

        # Closes the connection after the commitment
        conn.close()

        return queryResponse
    # Generic db error
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

        return None

def migr_db(conn, dump_path):
    connection = get_connection(conn)
    if connection is None:
        print("Could not connect to DB")
        return

    wallets = select_wallets(connection)
    if wallets is None or len(wallets) <= 0:
        print("No wallets present in DB")
        return

    f = open(dump_path, "a")

    for w in wallets:
        f.write("{0:9}{1:30}{2:24}{3:20}\n".format(w[0], w[1], w[2], w[3]))

    f.close()