//go:generate go run .
//go:generate go run generate_wasm.go
//go:generate npx tailwindcss-cli@latest build -i main.css -o style.css

package main

import (
	"frontend/src/components"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &components.Hello{})

	app.RunWhenOnBrowser()

	h := &app.Handler{
		Title: "Hello",
		Styles: []string{
			"style.css?direct",
			"app.css?direct",
		},
		Description: "A progressive web app written in Go.",
	}

	app.GenerateStaticWebsite(".", h)
}
