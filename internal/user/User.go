package user

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var Directory = "../../config/users"

type User struct {
	Username     string `json:"username"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
}

func getFilePath(username string) string {
	return filepath.Join(Directory, username+".json")
}

func addUser(username, clientID, clientSecret, redirectURI string) string {
	if err := os.MkdirAll(Directory, 0755); err != nil {
		return fmt.Sprintf("Error creating directory: %v", err)
	}

	if _, err := os.Stat(getFilePath(username)); !os.IsNotExist(err) {
		return "User already exists."
	}

	newUser := User{
		Username:     username,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
	}

	if err := saveUser(newUser); err != nil {
		return fmt.Sprintf("Error saving user: %v", err)
	}

	return "User added successfully."
}

func removeUser(username string) bool {
	filePath := getFilePath(username)
	if err := os.Remove(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println("Error removing file:", err)
		return false
	}
	return true
}

func removeAllUsers() bool {
	files, err := filepath.Glob(filepath.Join(Directory, "*.json"))
	if err != nil {
		fmt.Println("Error listing files:", err)
		return false
	}

	for _, file := range files {
		if err := os.Remove(file); err != nil {
			fmt.Println("Error removing file:", err)
			return false
		}
	}
	return true
}

func showUser(username string) string {
	user, err := loadUser(username)
	if err != nil {
		return fmt.Sprintf("Error loading user: %v", err)
	}
	return fmt.Sprintf("Username: %s, ClientID: %s, RedirectURI: %s", user.Username, user.ClientID, user.RedirectURI)
}

func showAllUsers() string {
	files, err := filepath.Glob(filepath.Join(Directory, "*.json"))
	if err != nil {
		fmt.Println("Error listing files:", err)
		return "Error retrieving users."
	}

	var result string
	for _, file := range files {
		username := filepath.Base(file[:len(file)-len(".json")])
		user, err := loadUser(username)
		if err != nil {
			result += fmt.Sprintf("Error loading user %s: %v\n", username, err)
			continue
		}
		result += fmt.Sprintf("Username: %s, ClientID: %s, RedirectURI: %s\n", user.Username, user.ClientID, user.RedirectURI)
	}

	if result == "" {
		return "No users found."
	}
	return result
}

func loadUser(username string) (User, error) {
	var user User
	filePath := getFilePath(username)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return user, fmt.Errorf("error reading file: %w", err)
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		return user, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return user, nil
}

func saveUser(user User) error {
	filePath := getFilePath(user.Username)

	data, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
