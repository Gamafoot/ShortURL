package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/url"
	"shorturl/internal/model"
	"shorturl/internal/pkg/logging"
	stringUtil "shorturl/internal/utils/string"
	"time"
)

func Index(c *fiber.Ctx) error {
	err := c.Query("error", "")

	return c.Render("index", fiber.Map{
		"Error": err,
	})
}

func Redirect(c *fiber.Ctx) error {
	shortURL := c.Params("short_url", "")
	shortURL = stringUtil.Strip(shortURL)

	if len(shortURL) == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	urlObj, err := model.GetURLByShortURL(shortURL)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Redirect(urlObj.Origin)
}

func Shorter(c *fiber.Ctx) error {
	urlObj := new(model.URL)

	originURL := c.FormValue("origin_url", "")
	originURL = stringUtil.Strip(originURL)

	if len(originURL) == 0 {
		return c.Redirect("/?error=the url has 0 length")
	}

	_, err := url.ParseRequestURI(originURL)
	if err != nil {
		logging.Log.Errorf("wrong url %v\n", err)
		return c.Redirect("/?error=wrong url")
	}

	urlObj.Origin = originURL
	urlObj.CreatedAt = time.Now()
	urlObj.GenerateShortURL()

	err = model.CreateURL(urlObj)

	fullURL := fmt.Sprintf("http://%s/redirect/%s", c.Hostname(), urlObj.Short)

	return c.Render("detail", fiber.Map{
		"URL": fullURL,
	})
}
