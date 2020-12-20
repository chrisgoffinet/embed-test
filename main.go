package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed static/*
var static embed.FS
var port = "8080"

func main() {
	fmt.Println("printing out embeded static assets:")
	fs.WalkDir(static, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("\t==> %s\n", path)
		return nil
	})
	fmt.Printf("listening on port: %s\n", port)
	http.Handle("/static/", http.FileServer(http.FS(static)))
	http.ListenAndServe(":"+port, nil)
}
