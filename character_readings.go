package main

import (
	"net/http"
	"github.com/satori/go.uuid"
	"github.com/asdine/storm"
	"github.com/labstack/echo"
)

type CharacterReading struct {
  ID string    `json:"uuid" storm:"id"`
  ReadingTuple `storm:"inline,unique"`
}

type CharacterReadingHandlers struct {
	db *storm.DB
}

func (h *CharacterReadingHandlers) Browse(c echo.Context) (err error) {
	list := make([]CharacterReading, 0)

	if err = h.db.All(&list); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, list)
}

func (h *CharacterReadingHandlers) Read(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := CharacterReading{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{ Message: "Not found." })
	}

	return c.JSON(http.StatusOK, item)
}

func (h *CharacterReadingHandlers) Edit(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := CharacterReading{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{ Message: "Not found." })
	}

	if err = c.Bind(&item); err != nil {
		return err
	}

	if err = h.db.Update(&item); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, item)
}

func (h *CharacterReadingHandlers) Add(c echo.Context) (err error) {
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

func (h *CharacterReadingHandlers) Destroy(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := CharacterReading{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{ Message: "Not found." })
	}

	if err = h.db.DeleteStruct(&item); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h * CharacterReadingHandlers) Wipe(c echo.Context) (err error) {
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
