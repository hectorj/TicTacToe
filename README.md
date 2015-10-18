# TicTacToe
A simple human vs CPU TicTacToe game, powered by Go. You can play it here: http://hectorj.net/TicTacToe/

[![Build Status](https://travis-ci.org/hectorj/TicTacToe.svg?branch=master)](https://travis-ci.org/hectorj/TicTacToe) [![GoDoc](https://godoc.org/github.com/hectorj/TicTacToe?status.svg)](https://godoc.org/github.com/hectorj/TicTacToe/) [![Coverage Status](https://coveralls.io/repos/hectorj/TicTacToe/badge.svg?branch=master)](https://coveralls.io/r/hectorj/TicTacToe?branch=master)

## Usage
There is currently 3 ways of using this package:
- Build your own UI and just use the [library](https://godoc.org/github.com/hectorj/TicTacToe/)
- `go run TicTacToe/web/server/main.go -listen=":80"` will start an HTTP server on port 80
- `cd $GOPATH/src/github.com/hectorj/TicTacToe/web/static/ && ./build.sh -dest=./public/` will generate all the static files needed in `$GOPATH/src/github.com/hectorj/TicTacToe/web/static/public`, so that you can just open the index.html file in your browser, or use any HTTP server.

Coming soonish: a GopherJS version
