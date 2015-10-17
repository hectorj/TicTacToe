package main

import (
	"TicTacToe"
	"TicTacToe/web"
	"flag"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	grids := TicTacToe.GetAllGrids()

	var destDir string
	flag.StringVar(&destDir, "dest", "./", "")
	flag.Parse()

	os.MkdirAll(destDir, os.ModeDir|0666)

	for _, grid := range grids {
		ID := grid.GetID()

		var filename string
		if ID == 1 {
			filename = "index"
		} else {
			filename = strconv.FormatUint(uint64(ID), 10)
		}

		filePath := filepath.Join(destDir, "/"+filename+".html")

		file, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}

		err = web.Render(file, ID)
		if err != nil {
			panic(err)
		}

		file.Close()
	}
}
