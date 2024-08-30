import os
import json
import toml

USERFILEPATH = "../config/users/"
GENERALFILEPATH = "../config/general.toml"

class UserService:
    def __init__(self, user_filepath=USERFILEPATH, general_filepath=GENERALFILEPATH):
        self.user_filepath = user_filepath
        self.general_filepath = general_filepath

    def addUser(self, username, client, secret, uri):
        userdata = {
            "Username": username,
            "Client-ID": client,
            "Client-Secret": secret,
            "Redirect-Uri": uri
        }

        userdata_json = json.dumps(userdata, indent=4)
        userpath = os.path.join(self.user_filepath, f"{username}.json")
        os.makedirs(os.path.dirname(userpath), exist_ok=True)

        if os.path.exists(userpath):
            print(f"User {username} already exists")
        else:
            try:
                with open(userpath, 'w') as file:
                    file.write(userdata_json)
                print(f"User data for {username} has been successfully written to {userpath}")
            except IOError as e:
                print(f"An error occurred while writing to the file: {e}")


    def showUser(self, username):
        userpath = os.path.join(self.user_filepath, f"{username}.json")
        if os.path.exists(userpath):
            try:
                with open(userpath, 'r') as file:
                    userdata = json.load(file)
                    print(json.dumps(userdata, indent=4))
            except IOError as e:
                print(f"An error occurred while reading from the file: {e}")
        else:
            print(f"User {username} does not exist")

    def removeUser(self, username):
        userpath = os.path.join(self.user_filepath, f"{username}.json")
        if os.path.exists(userpath):
            try:
                os.remove(userpath)
                print(f"User {username} has been successfully removed")
            except IOError as e:
                print(f"An error occurred while removing from the file: {e}")
        else:
            print(f"User {username} does not exist")


    def removeAllUsers(self):
        try:
            files = os.listdir(self.user_filepath)
            json_files = [f for f in files if f.endswith('.json')]
            for filename in json_files:
                file_path = os.path.join(self.user_filepath, filename)
                os.remove(file_path)
            print(f"All users have been removed.")
        except FileNotFoundError as e:
            print(f"Directory not found: {e}")
        except IOError as e:
            print(f"An error occurred while removing files: {e}")
        except Exception as e:
            print(f"An unexpected error occurred: {e}")

    def editUser(self, username, new_client=None, new_secret=None, new_uri=None):
        userpath = os.path.join(self.user_filepath, f"{username}.json")
        if not os.path.isfile(userpath):
            print(f"User file for {username} does not exist.")
            return

        try:
            with open(userpath, 'r') as file:
                userdata = json.load(file)

            if new_client is not None:
                userdata["Client-ID"] = new_client
            if new_secret is not None:
                userdata["Client-Secret"] = new_secret
            if new_uri is not None:
                userdata["Redirect-Uri"] = new_uri

            with open(userpath, 'w') as file:
                json.dump(userdata, file, indent=4)
            print(f"User data for {username} has been updated.")
        except json.JSONDecodeError as e:
            print(f"Error decoding JSON from file {userpath}: {e}")
        except IOError as e:
            print(f"An error occurred while reading or writing the file: {e}")
        except Exception as e:
            print(f"An unexpected error occurred: {e}")


    def showAllUsers(self):
        try:
            files = os.listdir(self.user_filepath)
            json_files = [f for f in files if f.endswith('.json')]

            if not json_files:
                print("No user files found.")
                return

            users = []

            for filename in json_files:
                file_path = os.path.join(self.user_filepath, filename)
                try:
                    with open(file_path, 'r') as file:
                        user_data = json.load(file)
                        users.append(user_data)
                except json.JSONDecodeError as e:
                    print(f"Error decoding JSON from file {filename}: {e}")
                except IOError as e:
                    print(f"An error occurred while reading the file {filename}: {e}")

            if users:
                print("List of all users:")
                for user in users:
                    print(json.dumps(user, indent=4))
            else:
                print("No valid user data found.")
        except FileNotFoundError as e:
            print(f"Directory not found: {e}")
        except Exception as e:
            print(f"An unexpected error occurred: {e}")


    def selectUser(self, username):
        userpath = os.path.join(self.user_filepath, f"{username}.json")
        if not os.path.isfile(userpath):
            print(f"User file for {username} does not exist.")
            return

        user_json_path = os.path.abspath(userpath)
        toml_data = {
            "Active-User": user_json_path
        }

        try:
            with open(self.general_filepath, 'w') as file:
                toml.dump(toml_data, file)
            print(f"User path for {username} has been written to {self.general_filepath}.")
        except IOError as e:
            print(f"An error occurred while writing to the TOML file: {e}")
        except Exception as e:
            print(f"An unexpected error occurred: {e}")