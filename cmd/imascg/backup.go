package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo"
)

func init() {
	e.GET("/downloads", func(c echo.Context) error {
		dbPath := os.Getenv("DB_PATH")
		items, err := ioutil.ReadDir(dbPath)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}

		names := make([]string, len(items))
		for i, f := range items {
			names[i] = f.Name()
		}

		links := make([]string, len(names))
		for i, name := range names {
			links[i] = fmt.Sprintf("<li>\n<a href=\"/downloads/%s\">%s</a>\n</li>", name, name)
		}

		list := fmt.Sprintf("<ul>\n%s\n</ul>", strings.Join(links, "\n"))
		lines := []string{
			"<!DOCTYPE html>",
			"<html lang=\"ja\">",
			"<head>",
			"<meta charset=\"utf-8\">",
			"</head>",
			"<body>",
			fmt.Sprintf("    %s", list),
			"</body>",
			"</html>",
		}

		html := strings.Join(lines, "\n")

		return c.HTML(http.StatusOK, html)
	})

	e.GET("/downloads/:name", func(c echo.Context) error {
		name := c.Param("name")
		dbPath := os.Getenv("DB_PATH")
		dbFile := fmt.Sprintf("%s/%s", dbPath, name)
		return c.Attachment(dbFile, name)
	})
}

func createBackup(t time.Time) {
	dbPath := os.Getenv("DB_PATH")
	dbFile := fmt.Sprintf("%s/imascg.db", dbPath)

	name, err := t.MarshalText()
	if err != nil {
		log.Printf("backup: %s", err)
		return
	}

	backup := fmt.Sprintf("%s/%s", dbPath, string(name))

	src, err := os.Open(dbFile)
	if err != nil {
		log.Printf("backup: %s", err)
		return
	}
	defer src.Close()

	dst, err := os.Create(backup)
	if err != nil {
		log.Printf("backup: %s", err)
		return
	}
	defer src.Close()

	if _, err := io.Copy(dst, src); err != nil {
		log.Printf("backup: %s", err)
		return
	}

	log.Printf("created backup: %s", backup)
}
