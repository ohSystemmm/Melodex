package com.melodex.user;

import com.melodex.list.LinkedList;

public class AllUsers {
    private LinkedList<User> users;

    public AllUsers() {
        this.users = new LinkedList<>();
    }

    public void addUser(User user) {
        users.add(user);
    }

    public boolean removeUser(User user) {
        return users.remove(user);
    }

    public void printAllUsers() {
        users.printList();
    }

    public int getSize() {
        return users.size();
    }
}
