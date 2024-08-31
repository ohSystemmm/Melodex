import java.util.Scanner;

public class Main {
    public static void main(String[] args) {

        System.out.println("Melodex");

        // Debug values for now
        String username = "Test1234";
        checkInput("^[A-Za-z][A-Za-z0-9]*$", username);

        String clientID = "7e4a9c13b77d46b9a77c2e5f9d6a1a8b";
        checkInput("^[a-fA-F0-9]{32}$", clientID);

        String clientSecret = "3f2b7d9c8a9e4b1c9f9d6e2a8c4d7e5a";
        checkInput("^[a-fA-F0-9]{32}$", clientSecret);

        String redirectURI = "http://localhost:1234/callback";
        checkInput("^http://localhost:\\d+/callback$", redirectURI);

        // Debug output
        System.out.print("\n");
        System.out.println("Username: " + username);
        System.out.println("Client ID: " + clientID);
        System.out.println("Client Secret: " + clientSecret);
        System.out.println("Redirect URI: " + redirectURI);

        Users users = new Users();
        users.addUser(username, clientID, clientSecret, redirectURI);
    }

    static String input(String message) {
        Scanner input = new Scanner(System.in);
        System.out.print(message);
        return input.nextLine();
    }

    static void checkInput(String regex, String content) {
        Validator validator = new Validator(regex);
        if (!validator.checkRegex(content)) {
            throw new CustomException("Invalid Input: " + content);
        }
    }
}
