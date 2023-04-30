package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/howeyc/gopass"
	"golang.org/x/crypto/bcrypt"
)

// Config struct to store password hashes
type Config struct {
	Passwords map[string][]byte
}

// ReadStdIn reads a line from standard input and returns it as a string
func ReadStdIn() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return string(input[:len(input)-1])
}

// ReadPassword reads a password from standard input without echoing it
func ReadPassword() string {
	passwordBytes, err := gopass.GetPasswd()
	if err != nil {
		log.Fatal(err)
	}
	return string(passwordBytes)
}

func main() {
	// Parse command-line flags
	add := flag.Bool("add", false, "Add a new password to the config.")
	configFile := flag.String("config", "config.json", "Config file with password hashes.")
	flag.Parse()

	// Initialize the configuration
	cfg := Config{
		Passwords: map[string][]byte{},
	}
	// Function to save the configuration to a file
	save := func() {
		cfgBytes, err := json.MarshalIndent(&cfg, "", "  ")
		if err != nil {
			log.Fatalf("Error creating empty config: %v", err)
		}
		if err := ioutil.WriteFile(*configFile, append(cfgBytes, []byte("\n")...), 0600); err != nil {
			log.Fatalf("Error writing empty config to file: %v", err)
		}
		log.Printf("Config saved to %s.", *configFile)
	}
	// Create an empty config file if it does not exist
	if _, err := os.Stat(*configFile); os.IsNotExist(err) {
		log.Println("File does not exist. Creating an empty config file.")
		save()
	}
	// Read the config file
	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	// Parse the config file
	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	// Add a new password to the config
	if *add {
		fmt.Print("Enter new account name: ")
		accountName := ReadStdIn()

		fmt.Print("Enter your password: ")
		password0 := ReadPassword()
		fmt.Print("Confirm password: ")
		password1 := ReadPassword()
		if password0 != password1 {
			log.Fatal("Passwords do not match.")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password0), 16)
		if err != nil {
			log.Fatal(err)
		}
		cfg.Passwords[accountName] = hashedPassword
		save()
		return
	}

	// Practice passwords
	for accountName, hashedPassword := range cfg.Passwords {
		fmt.Printf("Account: \"%s\"\n", accountName)
		for {
			fmt.Print("Password: ")
			password := ReadPassword()
			err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
			if err == nil {
				fmt.Printf("Password is correct\n\n")
				break
			}
			fmt.Println("Incorrect password")
		}
	}
}
