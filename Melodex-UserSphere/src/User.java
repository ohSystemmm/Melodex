import java.io.File;
import java.io.IOException;

public class User {
    public String username;
    public String clientID;
    public String clientSecret;
    public String redirectURI;
    public String clientPath = "../../Melodex-ConfigCenter/users/";
    public String generalPath = "../../Melodex-ConfigCenter/General.toml";

    public User () {

    }

    public void addUser(String username,String clientID, String clientSecret, String redirectURI) {
        File userdata = new File(clientPath + username + ".toml");
        try {
            if (!userdata.exists()) {
                if (userdata.createNewFile()) {
                    System.out.println("User " + username + " created");
                } else {
                    throw new CustomException("Failed to create file for user " + username);
                }
            } else {
                throw new CustomException("User " + username + " already exists");
            }
        } catch (IOException e) {
            System.err.println("IO Exception: " + e.getMessage());
        } catch (CustomException e) {
            System.err.println("Custom Exception: " + e.getMessage());
        }
    }

    public void removeUser(String username) {

    }

    public void removeAllUsers() {

    }

    public void editUser(String username) {

    }

    public void showUser(String username) {

    }

    public void listUsers() {

    }

    public void selectUser(String username) {

    }
}
