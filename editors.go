package main

import (
	"fmt"
	"github.com/asdine/storm"
	"github.com/labstack/echo"
	"net/http"
)

// Editor is the Storm model for characters.
type Editor struct {
	ID string `json:"id" storm:"id"`
}

// EditorHandlers defines the REST handlers for the Editor model.
type EditorHandlers struct {
	db *storm.DB
}

// Browse handler for the Editor model.
func (h EditorHandlers) Browse(c echo.Context) (err error) {
	list := make([]Editor, 0)

	if err = h.db.All(&list); err != nil {
		return c.JSON(http.StatusOK, list)
	}

	return c.JSON(http.StatusOK, list)
}

// Read handler for the Editor model.
func (h EditorHandlers) Read(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := Editor{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{Message: "Not found."})
	}

	return c.JSON(http.StatusOK, item)
}

// Edit handler for the Editor model.
func (h EditorHandlers) Edit(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := Editor{}

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

// Add handler for the Editor model.
func (h EditorHandlers) Add(c echo.Context) (err error) {
	item := Editor{}

	if err = c.Bind(&item); err != nil {
		fmt.Println(err)
		return err
	}

	if err = h.db.Save(&item); err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusCreated, item)
}

// Destroy handler for the Editor model.
func (h EditorHandlers) Destroy(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := Editor{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{Message: "Not found."})
	}

	if err = h.db.DeleteStruct(&item); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

// Wipe handler for the Editor model.
func (h EditorHandlers) Wipe(c echo.Context) (err error) {
	list := make([]Editor, 0)

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
