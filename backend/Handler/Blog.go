package Handler

import (
	"basic-blog/Database"

	"github.com/gofiber/fiber/v2"
)

func BlogAdd(c *fiber.Ctx) error {
	var blog BlogForm

	if err := c.BodyParser(&blog); err != nil {
		return err
	}

	if blog.Title != "" && blog.Description != "" {
		err, msg := Database.BlogAdd(blog.Title, blog.Description)
		if err != true {
			c.JSON(fiber.Map{
				"status": 502,
				"error":  msg,
			})
		} else {
			c.JSON(fiber.Map{
				"status":  "OK",
				"message": "Başarılı Eklendi",
			})
		}
	} else {
		c.JSON(fiber.Map{
			"status": 502,
			"error":  "Boş description yada title",
		})
	}
	return nil
}

func BlogViews(c *fiber.Ctx) error {
	var blog BlogInfo
	if err := c.BodyParser(&blog); err != nil {
		return err
	}

	blogs, errs := Database.BlogView(blog.Title)
	if errs != nil {
		return c.JSON(fiber.Map{
			"status": 404,
		})
	}

	c.JSON(fiber.Map{
		"status": "OK",
		"data":   blogs,
	})

	return nil
}


func BlogList(c *fiber.Ctx) error {
	blogs, errs := Database.BlogList()
	if errs != nil {
		return c.JSON(fiber.Map{
			"status": 404,
		})
	}

	c.JSON(fiber.Map{
		"status": "OK",
		"data":   blogs,
	})

	return nil
}


func BlogDel(c *fiber.Ctx) error {
	var blog BlogDels

	if err := c.BodyParser(&blog); err != nil {
		return err
	}

	if blog.Token != "" {
		_, err := Database.BlogDel(blog.Token)
		if err != true {
			c.JSON(fiber.Map{
				"status": 502,
				"error":  err,
			})
		} else {
			c.JSON(fiber.Map{
				"status":  "OK",
				"message": "Başarılı Silindi",
			})
		}
	} else {
		c.JSON(fiber.Map{
			"status": 502,
			"error":  "Boş token",
		})
	}
	return nil
}