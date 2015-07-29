package main

import "github.com/whobrokethebuild/goingup"

func main() {
	app := goingup.NewApp()
	app.Options.Port = 8080

	app.AddPage("/", "Home", "page")
	app.AddPage("/about", "About", "page")
	app.AddPage("/contact", "Contact", "page")

	app.Options.LoginAction = "/login"
	app.AddPage("/login", "Login", "login")
	app.Options.RegisterAction = "/register"
	app.AddPage("/register", "Register", "register")

	app.Options.Menus["Main"] = []goingup.MenuItem{
		{URL: "/", Text: "Home"},
		{URL: "/about", Text: "About"},
		{URL: "/contact", Text: "Contact"},
	}

	app.Run()
}
