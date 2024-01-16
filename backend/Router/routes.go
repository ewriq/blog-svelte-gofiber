package Router

import (
	Handler "basic-blog/Handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Initalize(app *fiber.App){
	v1 := app.Group("/api/v1")
	blog := app.Group("/api/v1/blog")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders:  "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	v1.Post("/user", Handler.User) 
    v1.Post("/login", Handler.Login)
    v1.Post("/register", Handler.Register)

	blog.Post("/add", Handler.BlogAdd)
    blog.Post("/view", Handler.BlogViews)
	blog.Post("/list", Handler.BlogList)
	blog.Post("/del", Handler.BlogDel)

	app.Get("/", Handler.Home)
	
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}