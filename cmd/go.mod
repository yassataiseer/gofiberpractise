module example.com/cmd

go 1.15

require (
	example.com/routes v0.0.0-00010101000000-000000000000
	github.com/gofiber/fiber/v2 v2.37.0
)

replace example.com/routes => ../routes
