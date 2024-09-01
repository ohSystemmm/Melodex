package com.melodex.user;

import java.util.Objects;

public class User {
    private final String username;
    private final String clientID;
    private final String clientSecret;
    private final String redirectURI;

    public User(String username, String clientID, String clientSecret, String redirectURI) {
        this.username = Objects.requireNonNull(username, "Username cannot be empty");
        this.clientID = Objects.requireNonNull(clientID, "Client-ID cannot be empty");
        this.clientSecret = Objects.requireNonNull(clientSecret, "Client-Secret cannot be empty");
        this.redirectURI = Objects.requireNonNull(redirectURI, "Redirect-URI cannot be empty");
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
        return "{\n" +
                "  \"Username\": \"" + username + "\",\n" +
                "  \"Client-ID\": \"" + clientID + "\",\n" +
                "  \"Client-Secret\": \"" + clientSecret + "\",\n" +
                "  \"Redirect-URI\": \"" + redirectURI + "\"\n" +
                "}";
    }
}
