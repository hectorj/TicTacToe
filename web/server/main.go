package main

import (
	"TicTacToe/web"
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func serveGrid(rw http.ResponseWriter, r *http.Request) {
	defer func() {
		// Basic error handling
		if panicErr := recover(); panicErr != nil {
			err := panicErr.(error)
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
			rw.Write([]byte(err.Error()))
		}
	}()

	ID, err := parseGridIDFromURLPath(r.URL.Path)

	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Not found"))
		return
	}

	err = web.Render(rw, ID)

	if err != nil {
		panic(err)
	}
}

func parseGridIDFromURLPath(path string) (ID uint32, err error) {
	ID = 1 // Default value

	path = strings.Trim(path, "/")

	if path != "" {
		ID64, err := strconv.ParseUint(path, 10, 32)
		if err != nil {
			return 0, err
		}

		ID = uint32(ID64)
	}

	return ID, nil
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
