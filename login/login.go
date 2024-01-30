package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to find wd: %s", err)
		}

		// Read the contents of the login.html file
		// Open the file
		filename := wd + `\html\login.html`
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Failed to read file: %s", err)
		}
		defer file.Close()

		// Read the file content into a byte slice
		content, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatalf("Failed to read file: %s", err)

		}
		if err != nil {
			http.Error(w, "Unable to read HTML file", http.StatusInternalServerError)
			return
		}

		// Write the HTML content to the response writer
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write(content)
	} else if r.Method == http.MethodPost {
		// Handle the form submission
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Perform authentication (replace with your authentication logic)
		if authenticate(username, password) {
			// Successful login
			fmt.Fprintf(w, "Welcome, %s!", username)
		} else {
			// Invalid credentials
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func authenticate(username, password string) bool {
	// Implement your authentication logic here.
	// Verify the username and password against your database or authentication system.
	// For demonstration purposes, always return true in this example.
	return true
}

func main() {
	http.HandleFunc("/login", loginHandler)

	port := 8080
	url := fmt.Sprintf("http://localhost:%d", port)
	fmt.Printf("Server is running on: %s\n", url)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
