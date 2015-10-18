package web

import (
	"html/template"
	"io"
	"path/filepath"
	"runtime"

	"strings"

	"github.com/hectorj/TicTacToe"
)

type TemplateData struct {
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

func PrepareData(ID uint32) TemplateData {
	data := TemplateData{
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

	data.IsOver, data.Winner = data.Grid.IsGameOver()

	return data
}

func (data *TemplateData) PlayCPUTurn() {
	coordinates := TicTacToe.BestNextMove(data.Grid)
	data.Grid.Play(coordinates)
	data.IsOver, data.Winner = data.Grid.IsGameOver()
}

func Render(wr io.Writer, data TemplateData) error {
	return templates.ExecuteTemplate(wr, "main", data)
}
