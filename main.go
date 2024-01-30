package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Load templates
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to find working directory: %s", err)
	}
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob(wd + "/html/*.html")),
	}
	e.Renderer = renderer

	// Route => handler
	e.GET("/login", loginPage)
	e.POST("/login", loginAttempt)

	// Start server
	port := 8080
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

// Handler
func loginPage(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func loginAttempt(c echo.Context) error {
	// Retrieve form values
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Perform authentication
	if authenticate(username, password) {
		// Successful login
		return c.String(http.StatusOK, fmt.Sprintf("Welcome, %s!", username))
	} else {
		// Invalid credentials
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
	}
}

func authenticate(username, password string) bool {
	// Implement your authentication logic here.
	// For demonstration purposes, always return true in this example.
	return true
}
