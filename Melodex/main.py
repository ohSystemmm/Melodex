from Melodex.User.UserService import UserService

# Testing
User1 = UserService()
User1.addUser("Test0", "1", "2", "https")
User1.addUser("Test1", "1", "2", "https")
User1.showUser("Test0")
User1.editUser("Test1", new_uri="unknown")
User1.showUser("Test1")
User1.showAllUsers()
User1.removeUser("Test1")
User1.showAllUsers()
User1.selectUser("Test0")
# User1.removeAllUsers()