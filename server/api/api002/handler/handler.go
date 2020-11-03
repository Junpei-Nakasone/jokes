package handler

import (
	"net/http"
	"nuxt-dadjokes/api/api002/domain"
	"nuxt-dadjokes/environment/db"

	"github.com/labstack/echo"
)

// FetchJokes is made for sample.
func FetchJokes(c echo.Context) error {

	db := db.CreateDBConnection()
	defer db.Close()

	result := []domain.Joke{}

	err := db.Find(&result).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
