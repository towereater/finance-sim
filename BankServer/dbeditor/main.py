from cryptography.fernet import Fernet
from random import randint
import sqlite3

def get_connection(db_file):
    conn = None

    try:
        conn = sqlite3.connect(db_file)
        return conn
    except sqlite3.Error as e:
        print(f"An error has occurred while establishing a connection: {e}")
    
    return conn

def create_wallets(conn):
    comm = '''CREATE TABLE IF NOT EXISTS wallets (
            id INTEGER PRIMARY KEY,
            iban CHAR(30) UNIQUE,
            cash FLOAT(15, 4) DEFAULT(0)
        )'''

    try:
        cur = conn.cursor()
        cur.execute(comm)
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

def insert_wallet(conn, params):
    comm = '''INSERT INTO wallets VALUES (?, ?)'''

    try:
        cur = conn.cursor()
        cur.execute(comm, params)
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

def select_wallets(conn):
    comm = '''SELECT * FROM wallets ORDER BY cash'''

    try:
        cur = conn.cursor()
        cur.execute(comm)
        return cur.fetchall()
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

def generate_wallets(conn):
    create_wallets(conn)

    for i in range(5):
        iban = Fernet.generate_key().decode()
        cash = randint(0, 10000)
        insert_wallet(conn, (iban, cash))

    for row in select_wallets(conn):
        print(row)

def create_accounts(conn):
    comm = '''CREATE TABLE IF NOT EXISTS accounts (
            id INTEGER PRIMARY KEY,
            username VARCHAR(20) NOT NULL UNIQUE,
            password VARCHAR(20) NOT NULL
        )'''

    try:
        cur = conn.cursor()
        cur.execute(comm)
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

def insert_account(conn, username, password):
    comm = '''INSERT INTO accounts (username, password) VALUES (?, ?)'''

    try:
        cur = conn.cursor()
        cur.execute(comm, (username, password))
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

def select_accounts(conn):
    comm = '''SELECT * FROM accounts'''

    try:
        cur = conn.cursor()
        cur.execute(comm)
        return cur.fetchall()
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

def generate_accounts(conn):
    create_accounts(conn)

    params = [
        ("andnic", "pass"),
        ("frau", "xdxd"),
        ("admin", "admin")
    ]

    for i in range(3):
        insert_account(conn, params[i])

    for row in select_accounts(conn):
        print(row)
    
def create_account_wallets(conn):
    comm = '''CREATE TABLE IF NOT EXISTS account_wallets (
            id INTEGER PRIMARY KEY,
            account_id INTEGER NOT NULL,
            wallet_id INTEGER NOT NULL,
            FOREIGN KEY (account_id) REFERENCES accounts (id),
            FOREIGN KEY (wallet_id) REFERENCES wallets (id)
        )'''

    try:
        cur = conn.cursor()
        cur.execute(comm)
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

def execute_command(conn, comm):
    try:
        cur = conn.cursor()
        cur.execute(comm)
        return cur.fetchall()
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

def main():
    db_path = "/Users/andnic/Documents/VSCode/FinanceSim/BankServer/sqlite/db/bankdata.db"
    db_mem = ":memory:"
    conn = get_connection(db_path)

    #generate_wallets(conn)
    #generate_accounts(conn)

    create_accounts(conn)
    create_wallets(conn)
    create_account_wallets(conn)

    #insert_account(conn, "andnic", "pass")
    for acc in select_accounts(conn):
        print(acc)

    conn.commit()
    conn.close()

# EXECUTE MAIN FUNCTION
if __name__ == "__main__":
    main()