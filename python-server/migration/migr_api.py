import requests

def migr_api(dump_path):
    url = "https://randomuser.me/api/?results=498&nat=gb,us"

    response = requests.get(url)

    if response.status_code != 200:
        print(f"Error with the API call: {response.status_code}")
        return

    accounts = response.json()["results"]

    f = open(dump_path, "a")
    id = 1001

    for acc in accounts:
        username = acc["login"]["username"]
        password = acc["login"]["password"]
        name = acc["name"]["first"]
        surname = acc["name"]["last"]
        birth = acc["dob"]["date"][:10]

        f.write("{0:9}{1:30}{2:20}{3:20}{4:20}{5:10}\n"
            .format(id, username, password, name, surname, birth))

        id = id+1

    f.close()