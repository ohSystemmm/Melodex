package com.melodex;

import com.melodex.exception.CustomException;
import com.melodex.validator.Validator;

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        System.out.println("Welcome to Melodex!\n");
    }

    private static void createUser() {
    }

    private static void deleteUser() {
    }

    private static void modifyUser() {
    }

    private static void listUsers() {
    }

    private static void switchUser() {
    }

    private static String input(String message) {
        Scanner scanner = new Scanner(System.in);
        System.out.print(message);
        return scanner.nextLine();
    }

    private static boolean isInputValid(String regex, String content) {
        Validator validator = new Validator(regex);
        try {
            return validator.checkRegex(content);
        } catch (Exception e) {
            System.out.println("Error validating input: " + e.getMessage());
            return false;
        }
    }
}
