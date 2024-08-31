package com.melodex.user;

public class User {
    private final String username;
    private final String clientID;
    private final String clientSecret;
    private final String redirectURI;

    public User(String username, String clientID, String clientSecret, String redirectURI) {
        this.username = username;
        this.clientID = clientID;
        this.clientSecret = clientSecret;
        this.redirectURI = redirectURI;
    }

    public String getUsername() {
        return username;
    }

    public String getClientID() {
        return clientID;
    }

    public String getClientSecret() {
        return clientSecret;
    }

    public String getRedirectURI() {
        return redirectURI;
    }

    @Override
    public String toString() {
        return "User{" +
                "username='" + username + '\'' +
                ", clientID='" + clientID + '\'' +
                ", clientSecret='" + clientSecret + '\'' +
                ", redirectURI='" + redirectURI + '\'' +
                '}';
    }
}
