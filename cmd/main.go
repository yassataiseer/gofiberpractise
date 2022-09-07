package main

import (
	"fmt"
	"log"

	"example.com/newmodel"
	"example.com/routes"
	"github.com/gofiber/fiber/v2"
)


type POSTSTRUCT struct {
	Name string
	Data string
	Age int16	
}

func getPosts() []POSTSTRUCT {
	NoteData  := []POSTSTRUCT{

	} 
	FirstNote := POSTSTRUCT{		Name: "Yassa Taiseer",
	Data: "July 29,2022",
	Age: 16,
}
	NoteData = append(NoteData, FirstNote)
	return NoteData
}

func Process(c *fiber.Ctx) error {
	posts := getPosts() // your logic
	if len(posts) == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"error":   "There are no posts!",
		})
	}
	return c.JSON(&fiber.Map{
		"success": true,
		"posts":   posts,
	})
}
func main() {
    app := fiber.New()
	routes.Register(app)
	string1:=newmodel.Note()
	fmt.Println(string1)
	app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
	app.Get("/api/posts", Process)
    log.Fatal(app.Listen(":3000"))
}
