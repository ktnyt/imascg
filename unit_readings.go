package main

import (
	"net/http"

	"github.com/asdine/storm"
	"github.com/ktnyt/go.uuid"
	"github.com/labstack/echo"
)

func init() {
	g := e.Group("/units/readings")
	Register(g, UnitReadingHandlers{db: db})
}

// UnitReading is the Storm model for unit readings.
type UnitReading struct {
	ID           string `json:"uuid" storm:"id"`
	ReadingTuple `storm:"unique"`
}

// UnitReadingHandlers defines the REST handlers for the UnitReading model.
type UnitReadingHandlers struct {
	db *storm.DB
}

// Browse handler for the UnitReadings model.
func (h UnitReadingHandlers) Browse(c echo.Context) (err error) {
	list := make([]UnitReading, 0)

	if err = h.db.All(&list); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, list)
}

// Read handler for the UnitReadings model.
func (h UnitReadingHandlers) Read(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := UnitReading{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{Message: "Not found."})
	}

	return c.JSON(http.StatusOK, item)
}

// Edit handler for the UnitReadings model.
func (h UnitReadingHandlers) Edit(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := UnitReading{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{Message: "Not found."})
	}

	if err = h.db.Update(&item); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, item)
}

// Add handler for the UnitReadings model.
func (h UnitReadingHandlers) Add(c echo.Context) (err error) {
	item := UnitReading{
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

// Destroy handler for the UnitReadings model.
func (h UnitReadingHandlers) Destroy(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := UnitReading{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{Message: "Not found."})
	}

	if err = h.db.DeleteStruct(&item); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

// Wipe handler for the UnitReadings model.
func (h UnitReadingHandlers) Wipe(c echo.Context) (err error) {
	list := make([]UnitReading, 0)

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
