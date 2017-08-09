package main

import (
	"github.com/asdine/storm"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
)

func init() {
	g := e.Group("/characters/readings")
	Register(g, CharacterReadingHandlers{db: db})
}

// CharacterReading is the Storm model for character readings.
type CharacterReading struct {
	ID           string `json:"uuid" storm:"id"`
	ReadingTuple `storm:"inline,unique"`
}

// CharacterReadingHandlers defines the REST handlers for the CharacterReading model.
type CharacterReadingHandlers struct {
	db *storm.DB
}

// Browse handler for the CharacterReadings model.
func (h CharacterReadingHandlers) Browse(c echo.Context) (err error) {
	list := make([]CharacterReading, 0)

	if err = h.db.All(&list); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, list)
}

// Read handler for the CharacterReadings model.
func (h CharacterReadingHandlers) Read(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := CharacterReading{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{Message: "Not found."})
	}

	return c.JSON(http.StatusOK, item)
}

// Edit handler for the CharacterReadings model.
func (h CharacterReadingHandlers) Edit(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := CharacterReading{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{Message: "Not found."})
	}

	if err = c.Bind(&item); err != nil {
		return err
	}

	if err = h.db.Update(&item); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, item)
}

// Add handler for the CharacterReadings model.
func (h CharacterReadingHandlers) Add(c echo.Context) (err error) {
	item := CharacterReading{
		ID: uuid.NewV4().String(),
	}

	if err = c.Bind(&item); err != nil {
		return err
	}

	if err = h.db.Save(&item); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, item)
}

// Destroy handler for the CallTable model.
func (h CharacterReadingHandlers) Destroy(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := CharacterReading{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{Message: "Not found."})
	}

	if err = h.db.DeleteStruct(&item); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

// Wipe handler for the CallTable model.
func (h CharacterReadingHandlers) Wipe(c echo.Context) (err error) {
	list := make([]CharacterReading, 0)

	if err = h.db.All(&list); err != nil {
		return err
	}

	for _, item := range list {
		if err = h.db.DeleteStruct(&item); err != nil {
			return err
		}
	}

	return c.NoContent(http.StatusNoContent)
}
