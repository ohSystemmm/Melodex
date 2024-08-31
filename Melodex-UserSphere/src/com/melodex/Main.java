package com.melodex;

import com.melodex.user.AllUsers;
import com.melodex.user.User;

public class Main {
    public static void main(String[] args) {
        AllUsers allUsers = new AllUsers();

        User user1 = new User("Test1234", "clientID123", "clientSecret123", "http://localhost:8080/callback");
        User user2 = new User("Test5678", "clientID456", "clientSecret456", "http://localhost:8080/callback");

        allUsers.addUser(user1);
        allUsers.addUser(user2);

        System.out.println("All users:");
        allUsers.printAllUsers();

        System.out.println("Number of users: " + allUsers.getSize());

        allUsers.removeUser(user1);
        System.out.println("All users after removal:");
        allUsers.printAllUsers();
        System.out.println("Number of users: " + allUsers.getSize());
    }
}

