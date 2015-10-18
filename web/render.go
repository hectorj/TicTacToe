package web

import (
	"html/template"
	"io"
	"path/filepath"
	"runtime"

	"strings"

	"github.com/hectorj/TicTacToe"
)

type templateData struct {
	TicTacToe.Grid
	Coordinates [3][3]TicTacToe.Coordinates
	IsOver      bool
	FirstTurn   bool
	Winner      TicTacToe.NullPlayer
}

var (
	templates *template.Template
)

func init() {
	_, basepath, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(basepath)
	if strings.HasSuffix(basepath, "_test/_obj_test") {
		// This little hack is necessary for tests with coverage. @TODO: find a better solution
		basepath = "."
	}

	templatesPath := filepath.Join(basepath, "/template/*")

	templates = template.Must(template.ParseGlob(templatesPath))
}

func prepareData(ID uint32) templateData {
	data := templateData{
		Grid: TicTacToe.GridFromID(ID),
		Coordinates: [3][3]TicTacToe.Coordinates{
			{
				TicTacToe.Coordinates{X: 0, Y: 0},
				TicTacToe.Coordinates{X: 1, Y: 0},
				TicTacToe.Coordinates{X: 2, Y: 0},
			},
			{
				TicTacToe.Coordinates{X: 0, Y: 1},
				TicTacToe.Coordinates{X: 1, Y: 1},
				TicTacToe.Coordinates{X: 2, Y: 1},
			},
			{
				TicTacToe.Coordinates{X: 0, Y: 2},
				TicTacToe.Coordinates{X: 1, Y: 2},
				TicTacToe.Coordinates{X: 2, Y: 2},
			},
		},
		FirstTurn: ID <= 1,
	}

	if !data.FirstTurn {
		data.IsOver, data.Winner = data.Grid.IsGameOver()

		if !data.IsOver {
			coordinates := TicTacToe.BestNextMove(data.Grid)
			data.Grid.Play(coordinates)
			data.IsOver, data.Winner = data.Grid.IsGameOver()
		}
	}

	return data
}

func Render(wr io.Writer, ID uint32) error {
	return templates.ExecuteTemplate(wr, "main", prepareData(ID))
}
