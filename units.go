package main

import (
	"net/http"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/ktnyt/go.uuid"
	"github.com/labstack/echo"
)

func init() {
	g := e.Group("/units")
	Register(g, UnitHandlers{db: db})
}

// Unit is the Storm model for units.
type Unit struct {
	ID      string   `json:"id"   storm:"id"`
	Name    string   `json:"name" storm:"unique"`
	Members []string `json:"members"`
}

// UnitHandlers defines the REST handlers for the Unit model.
type UnitHandlers struct {
	db *storm.DB
}

// Browse handler for the Unit model.
func (h UnitHandlers) Browse(c echo.Context) (err error) {
	search := c.QueryParam("search")

	list := make([]Unit, 0)

	if len(search) > 0 {
		tmp := make([]UnitReading, 0)

		if err = h.db.Select(ReadingSubstr("ReadingTuple", search)).Find(&tmp); err != nil {
			return c.JSON(http.StatusOK, list)
		}

		pks := make([]string, len(tmp))

		for i := range tmp {
			pks[i] = tmp[i].ReadingTuple.ID
		}

		if err = h.db.Select(q.In("ID", pks)).Find(&list); err != nil {
			return c.JSON(http.StatusOK, list)
		}
	} else {
		if err = h.db.All(&list); err != nil {
			return c.JSON(http.StatusOK, list)
		}
	}

	return c.JSON(http.StatusOK, list)
}

// Read handler for the Unit model.
func (h UnitHandlers) Read(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := Unit{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{Message: "Not found."})
	}

	return c.JSON(http.StatusOK, item)
}

// Edit handler for the Unit model.
func (h UnitHandlers) Edit(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := Unit{}

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

// Add handler for the Unit model.
func (h UnitHandlers) Add(c echo.Context) (err error) {
	item := Unit{ID: uuid.NewV4().String()}

	if err = c.Bind(&item); err != nil {
		return err
	}

	if err = h.db.Save(&item); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, item)
}

// Destroy handler for the Unit model.
func (h UnitHandlers) Destroy(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := Unit{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{Message: "Not found."})
	}

	if err = h.db.DeleteStruct(&item); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

// Wipe handler for the Unit model.
func (h UnitHandlers) Wipe(c echo.Context) (err error) {
	list := make([]Unit, 0)

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
