package web

import (
	"html/template"
	"io"
	"path/filepath"
	"runtime"

	"github.com/hectorj/TicTacToe"
)

type templateData struct {
	TicTacToe.Grid
	Coordinates [3][3]TicTacToe.Coordinates
	IsOver      bool
	FirstTurn   bool
	Winner      TicTacToe.Player
}

var (
	templates *template.Template
)

func init() {
	_, basepath, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(basepath)
	templatesPath := filepath.Join(basepath, "/template/*")

	templates = template.Must(template.ParseGlob(templatesPath))
}

func prepareData(ID uint32) templateData {
	data := templateData{
		Grid: TicTacToe.GridFromID(ID),
		Coordinates: [3][3]TicTacToe.Coordinates{
			{
				TicTacToe.Coordinates{0, 0},
				TicTacToe.Coordinates{1, 0},
				TicTacToe.Coordinates{2, 0},
			},
			{
				TicTacToe.Coordinates{0, 1},
				TicTacToe.Coordinates{1, 1},
				TicTacToe.Coordinates{2, 1},
			},
			{
				TicTacToe.Coordinates{0, 2},
				TicTacToe.Coordinates{1, 2},
				TicTacToe.Coordinates{2, 2},
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
