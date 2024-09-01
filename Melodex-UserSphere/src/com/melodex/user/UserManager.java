package com.melodex.user;

import java.util.ArrayList;
import java.util.List;

public class UserManager {
    private final List<User> users = new ArrayList<>();

    public void createUser(User user) {
        if (user == null) {
            throw new IllegalArgumentException("User cannot be null");
        }
        users.add(user);
    }

    public void deleteUser(String username) {
        if (username == null) {
            throw new IllegalArgumentException("Username cannot be null");
        }
        users.removeIf(user -> user.getUsername().equals(username));
    }

    public void modifyUser(String username, User updatedUser) {
        if (username == null || updatedUser == null) {
            throw new IllegalArgumentException("Username and updatedUser cannot be null");
        }
        for (int i = 0; i < users.size(); i++) {
            if (users.get(i).getUsername().equals(username)) {
                users.set(i, updatedUser);
                return;
            }
        }
        throw new IllegalArgumentException("User not found");
    }

    public void listUsers() {
        if (users.isEmpty()) {
            System.out.println("No users available.");
        } else {
            for (User user : users) {
                System.out.println(user);
            }
        }
    }
}