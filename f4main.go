package main

import (
	"fmt"
	"html/template"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Create a new template engine instance
	engine := template.New("views")
	var Cont2 string

	// // Parse your template files
	// templateFile := `
	// <!DOCTYPE html>
	// <html>
	// <head>
	//     <title>{{.Title}}</title>
	// </head>
	// <body>
	//     <h1>{{.Content}}</h1>
	// 	<p1>{ "hello ,"}</p1>
	// 	<p2>{"This is my first document"}</p2>
	// 	<p3>{"Welcome to fiber"}</p3>
	// 	<p4>{{.Body}}</p4>

	// </body>
	// </html>

	// `

	// t, err := engine.Parse(templateFile)
	// if err != nil {
	// 	// Handle the error
	// 	return
	// }

	// // Define a route to render the template
	// app.Get("/web", func(c *fiber.Ctx) error {
	// 	data := struct {
	// 		Title   string
	// 		Content string
	// 		Body    string
	// 	}{
	// 		Title:   "My Page",
	// 		Content: "Welcome to my website!",
	// 		Body:    "this is fourth para",
	// 	}
	// 	return t.Execute(c.Response().BodyWriter(), data)
	// })
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	data := struct {
	// 		Title   string
	// 		Content string
	// 	}{
	// 		Title:   "My Page",
	// 		Content: "Welcome to my website!",
	// 	}

	// 	// Add a condition to display additional content
	// 	if err==nil {
	// 		data.Content = "You're a registered user!"
	// 	}

	// 	return t.Execute(c.Response().BodyWriter(), data)
	// })

	// Define a layout template

	layoutTemplate := `

<!DOCTYPE html>

<html>

<head>

    <title>{{.Title}}</title>

</head>

<body>

    <header>

        {{.Header}}

    </header>

    <main>

       <p1> {{.Content1}}<p1/>

		<p2>{{.Content2}}<p2/>

    </main>

    <footer>

        {{.Footer}}

    </footer>

</body>

</html>

`

	// Define partial templates

	// 	headerTemplate := `

	// <h1>{{.Title}}</h1>

	// `

	// 	contentTemplate := `

	// <p>{{.Content}}</p>

	// `

	// 	footerTemplate := `

	// <p>&copy; 2023 My Website</p>

	// `

	// Parse the layout and partial templates

	layout, err := engine.Parse(layoutTemplate)

	if err != nil {

		// Handle the error

		return

	}

	// header, err := engine.Parse(headerTemplate)

	// if err != nil {

	// 	// Handle the error

	// 	return

	// }

	// content, err := engine.Parse(contentTemplate)

	// if err != nil {

	// 	// Handle the error

	// 	return

	// }

	// footer, err := engine.Parse(footerTemplate)

	// if err != nil {

	// 	// Handle the error

	// 	return

	// }

	// Define a route to render the view with the layout and partials

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Print("enter the content to be Written")
		fmt.Scanln(&Cont2)

		data := struct {
			Title    string
			Header   string
			Content2 string
			Content1 string
			Footer   string
		}{

			Title:    "My Page",
			Header:   "My Website",
			Content1: "Welcome to my website!",
			Content2: Cont2,
			Footer:   "<p>&copy; 2023 My Website ",
		}

		// Execute the layout template and include partials

		layout.Execute(c.Response().BodyWriter(), data)

		return nil

	})

	app.Listen(":3010")
}
