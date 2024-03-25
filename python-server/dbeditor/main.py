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

def insert_account(conn, username, password):
    comm = '''INSERT INTO accounts (username, password) VALUES (?, ?)'''

    try:
        cur = conn.cursor()
        cur.execute(comm, (username, password))
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

def select_accounts(conn, username):
    comm = '''SELECT username, password FROM accounts WHERE username = ?'''

    try:
        cur = conn.cursor()
        cur.execute(comm, (username,))
        return cur.fetchall()
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")
    
def create_account_wallets(conn):
    comm = '''CREATE TABLE IF NOT EXISTS account_wallets (
            id INTEGER PRIMARY KEY,
            account_id INTEGER NOT NULL,
            wallet_id INTEGER NOT NULL,
            FOREIGN KEY (wallet_id) REFERENCES wallets (id)
        )'''

    try:
        cur = conn.cursor()
        cur.execute(comm)
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

def insert_account_wallets(conn, username, password):
    comm = '''INSERT INTO account_wallets (account_id, wallet_id) VALUES (?, ?)'''

    try:
        cur = conn.cursor()
        cur.execute(comm, (username, password))
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

def update_account_wallets(conn, account_id_new, account_id_old):
    comm = '''UPDATE account_wallets SET account_id = ? WHERE account_id = ?'''

    try:
        cur = conn.cursor()
        cur.execute(comm, (account_id_new, account_id_old))
    except sqlite3.Error as e:
        print(f"An error has occurred while executing a command: {e}")

def select_account_wallets(conn, account_id):
    comm = '''SELECT wallet_id FROM account_wallets WHERE account_id = ?'''

    try:
        cur = conn.cursor()
        cur.execute(comm, (account_id,))
        return cur.fetchall()
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

    print(select_accounts(conn, "andnic") is None)
    for row in select_accounts(conn, "andnic"):
        print(row[0])
        print(row[1])

    conn.commit()
    conn.close()

# EXECUTE MAIN FUNCTION
if __name__ == "__main__":
    main()