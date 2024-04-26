import sys
from migr_db import migr_db
from migr_api import migr_api

def main():
    db_conn = sys.argv[1]
    dump_path = sys.argv[2]

    #Logging
    print("Running with DB connection {0} and dump path {1}", db_conn, dump_path)

    #Deleting file content
    f = open(dump_path, "w")
    f.close()

    #Dumping the user DB
    migr_db(db_conn, dump_path)

    #Dumping other random users
    migr_api(dump_path)

# EXECUTE MAIN FUNCTION
if __name__ == "__main__":
    main()