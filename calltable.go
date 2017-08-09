package main

import (
	"fmt"
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"strings"
)

func init() {
	g := e.Group("/calltable")
	Register(g, CallTableHandlers{db: db})
}

// CallTableTuple holds the calling information body.
type CallTableTuple struct {
	Caller string `json:"caller"`
	Callee string `json:"callee"`
	Called string `json:"called"`
	Remark string `json:"remark"`
}

// CallTable is the Storm model for the call table.
type CallTable struct {
	ID             string `json:"id" storm:"id"`
	CallTableTuple `storm:"inline,unique"`
}

// CallTableHandlers defines the REST handlers for the CallTable model.
type CallTableHandlers struct {
	db *storm.DB
}

// Browse handler for the CallTable model.
func (h CallTableHandlers) Browse(c echo.Context) (err error) {
	list := make([]CallTable, 0)

	if len(c.QueryParams()) > 0 {
		caller := c.QueryParam("caller")
		callee := c.QueryParam("callee")
		called := c.QueryParam("called")
		remark := c.QueryParam("remark")
		limit := c.QueryParam("limit")
		skip := c.QueryParam("skip")

		qs := make([]q.Matcher, 0)

		if len(caller) > 0 {
			qs = append(qs, q.In("Caller", strings.Split(caller, ",")))
		}

		if len(callee) > 0 {
			qs = append(qs, q.In("Callee", strings.Split(callee, ",")))
		}

		if len(called) > 0 {
			qs = append(qs, Ss("Called", called))
		}

		if len(remark) > 0 {
			qs = append(qs, Ss("Remark", remark))
		}

		query := h.db.Select(qs...)

		if len(limit) > 0 {
			value, convErr := strconv.Atoi(limit)
			if convErr != nil {
				return convErr
			}
			query = query.Limit(value)
		}

		if len(skip) > 0 {
			value, convErr := strconv.Atoi(skip)
			if convErr != nil {
				return convErr
			}
			query = query.Skip(value)
		}

		if err = query.Find(&list); err != nil {
			return c.JSON(http.StatusOK, list)
		}
	} else {
		if err = h.db.All(&list); err != nil {
			return c.JSON(http.StatusOK, list)
		}
	}

	return c.JSON(http.StatusOK, list)
}

// Read handler for the CallTable model.
func (h CallTableHandlers) Read(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := CallTable{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{Message: "Not found."})
	}

	return c.JSON(http.StatusOK, item)
}

// Edit handler for the CallTable model.
func (h CallTableHandlers) Edit(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := CallTable{}

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

func formatID(caller string, callee string, n int) string {
	return fmt.Sprintf("%s%s%d", caller, callee, n)
}

// Add handler for the CallTable model.
func (h CallTableHandlers) Add(c echo.Context) (err error) {
	item := CallTable{}

	if err = c.Bind(&item); err != nil {
		return err
	}

	tmp := CallTable{}

	n := 0
	id := formatID(item.Caller, item.Callee, n)

	for {
		if err = h.db.One("ID", id, &tmp); err != nil {
			break
		} else {
			id = formatID(item.Caller, item.Callee, n)
			n++
		}
	}

	item.ID = id

	if err = h.db.Save(&item); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, item)
}

// Destroy handler for the CallTable model.
func (h CallTableHandlers) Destroy(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := CallTable{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{Message: "Not found."})
	}

	if err = h.db.DeleteStruct(&item); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

// Wipe handler for the CallTable model.
func (h CallTableHandlers) Wipe(c echo.Context) (err error) {
	list := make([]CallTable, 0)

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
