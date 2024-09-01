package com.melodex;

import com.melodex.user.User;
import com.melodex.user.UserManager;
import com.melodex.validator.Validator;
import com.melodex.exception.CustomException;

public class Main {
    public static void main(String[] args) {
        System.out.println("Welcome to Melodex!\n");

        UserManager userManager = new UserManager();
        Validator usernameValidator = new Validator("^[a-zA-Z0-9_]{5,15}$");
        Validator clientIDValidator = new Validator("^[a-zA-Z0-9]{5,10}$"); // Example regex
        Validator redirectURIValidator = new Validator("^(http|https)://.*$"); // Example regex

        try {
            User user = new User("john_doe", "12345", "secret", "http://example.com");

            if (usernameValidator.checkRegex(user.getUsername()) &&
                    clientIDValidator.checkRegex(user.getClientID()) &&
                    redirectURIValidator.checkRegex(user.getRedirectURI())) {

                userManager.createUser(user);
                System.out.println("User created successfully.");

                // List users to verify creation
                System.out.println("\nCurrent Users:");
                userManager.listUsers();

            } else {
                throw new CustomException("Invalid input format.");
            }

        } catch (CustomException e) {
            System.err.println("Error: " + e.getMessage());
        }
    }
}
