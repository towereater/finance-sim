import sqlite3

def get_connection(db_file):
    conn = None

    try:
        conn = sqlite3.connect(db_file)
        return conn
    except sqlite3.Error as e:
        print(f"An error has occurred while establishing a connection: {e}")
    
    return conn

def select_accounts(conn):
    comm = '''SELECT id, username, password FROM accounts'''

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

    accounts = select_accounts(connection)
    if accounts is None or len(accounts) <= 0:
        print("No accounts present in DB")
        return

    f = open(dump_path, "a")

    for acc in accounts:
        f.write("{0:9}{1:30}{2:20}\n".format(acc[0], acc[1], acc[2]))

    f.close()