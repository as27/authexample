package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"

	"github.com/markbates/goth/gothic"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			Env:         ENV,
			SessionName: "_authexample_session",
			Host:        "http://localhost:3000",
		})

		app.ServeFiles("/assets", assetsPath())
		auth := app.Group("/auth")
		auth.GET("/{provider}", buffalo.WrapHandlerFunc(gothic.BeginAuthHandler))
		auth.GET("/{provider}/callback", AuthCallback)
		app.GET("/login",
			func(c buffalo.Context) error {
				return c.Render(200, r.HTML("login/index.html"))
			})
		secure := app.Group("/secure")
		secure.Use(CheckAuth)
		secure.GET("/",
			func(c buffalo.Context) error {
				return c.Render(200, r.HTML("secure/index.html"))
			})
		secure.DELETE("/logout",
			func(c buffalo.Context) error {
				session := c.Session()
				session.Delete("userID")
				session.Save()
				return c.Redirect(301, "/login")
			})
		app.Redirect(301, "/", "/login")
	}

	return app
}
