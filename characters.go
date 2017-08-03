package main

import (
	"fmt"
	"net/http"
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/labstack/echo"
)

type Character struct {
	ID   string `json:"id"   storm:"id"`
	Name string `json:"name" storm:"unique"`
	Type string `json:"type"`
}

type CharacterHandlers struct {
	db *storm.DB
}

func (h *CharacterHandlers) Browse(c echo.Context) (err error) {
	search := c.QueryParam("search")

	list := make([]Character, 0)

	if len(search) > 0 {
		tmp := make([]CharacterReading, 0)

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
	} else{
		if err = h.db.All(&list); err != nil {
			return c.JSON(http.StatusOK, list)
		}
	}

	return c.JSON(http.StatusOK, list)
}

func (h *CharacterHandlers) Read(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := Character{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{ Message: "Not found." })
	}

	return c.JSON(http.StatusOK, item)
}

func (h *CharacterHandlers) Edit(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := Character{}

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

func (h *CharacterHandlers) Add(c echo.Context) (err error) {
	item := Character{}

	if err = c.Bind(&item); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("%+v\n", len(item.ID))

	if len(item.ID) == 0 {
		list := make([]Character, 0)

		if err = h.db.Find("Type", item.Type, &list); err != nil {
			return err
		}

		item.ID = fmt.Sprint(3000 + len(list))
	}

	if err = h.db.Save(&item); err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusCreated, item)
}

func (h *CharacterHandlers) Destroy(c echo.Context) (err error) {
	pk := c.Param("pk")

	item := Character{}

	if err = h.db.One("ID", pk, &item); err != nil {
		return c.JSON(http.StatusNotFound, Message{ Message: "Not found." })
	}

	if err = h.db.DeleteStruct(&item); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h * CharacterHandlers) Wipe(c echo.Context) (err error) {
	list := make([]Character, 0)

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
