package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/dropbox/godropbox/errors"
	"github.com/hectorj/TicTacToe/web"
)

func serveGrid(rw http.ResponseWriter, r *http.Request) {
	defer func() {
		// Basic error handling
		if panicErr := recover(); panicErr != nil {
			err := panicErr.(error)
			log.Print(err)
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
			rw.Write([]byte(err.Error()))
		}
	}()

	ID, IAPlaysFirst, err := parseURLPath(r.URL.Path)

	if err != nil {
		log.Print(err)
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Not found"))
		return
	}

	templateData := web.PrepareData(ID)

	if !templateData.IsOver && (!templateData.FirstTurn || IAPlaysFirst) {
		templateData.PlayCPUTurn()
	}

	err = web.Render(rw, templateData)

	if err != nil {
		panic(err)
	}
}

func parseURLPath(path string) (ID uint32, IAPlaysFirst bool, err error) {
	path = strings.Trim(path, "/")

	if path == "" {
		return 1, false, nil // blank grid, X first
	} else {
		if !strings.HasSuffix(path, ".html") {
			return 0, false, errors.New("missing html suffix")
		}

		path = strings.TrimSuffix(path, ".html")

		if path == "cpu" {
			return 1, true, nil // blank grid, X first
		} else {
			ID64, err := strconv.ParseUint(path, 10, 32)
			if err != nil {
				return 0, false, err
			}

			return uint32(ID64), false, nil
		}
	}
}

func main() {
	_, basepath, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(basepath)
	assetsPath := filepath.Join(basepath, "../assets/dist/")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(assetsPath))))
	http.HandleFunc("/", serveGrid)

	var listenHost string
	flag.StringVar(&listenHost, "listen", "127.0.0.1:8080", "")
	flag.Parse()

	fmt.Println("Listening to ", listenHost)
	err := http.ListenAndServe(listenHost, nil)
	if err != nil {
		panic(err)
	}
}
