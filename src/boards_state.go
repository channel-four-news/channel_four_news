package main

import (
    "io"
    "fmt"
    "sync"
    "encoding/json"
    "golang.org/x/net/html"
)

var (
    BOARDS = "boards.json"
)

type FourBoard struct {
    Board           string          `json:"board"`
    Title           string          `json:"title"`
    Description     string          `json:"meta_description"`
}

type BoardsState struct {
    Boards          []FourBoard     `json:"boards"`
    CurrentBoard    int
    WaitGroup       sync.WaitGroup
    Mtx             sync.Mutex
}

func (boards_state *BoardsState) GetTitle(board int) string {
    assert(board < len(boards_state.Boards),
           "Board Index is out of range")
    return boards_state.Boards[board].Board
}

func (boards_state *BoardsState) GetCurrentTitle() string {
    assert(boards_state.CurrentBoard < len(boards_state.Boards),
           "Board Index is out of range")
    return boards_state.Boards[boards_state.CurrentBoard].Board
}

func (boards_state *BoardsState) SetCurrentBoard(board int) {
    assert(board < len(boards_state.Boards),
           "Board Index is out of range")
    boards_state.Mtx.Lock()
    boards_state.CurrentBoard = board
    boards_state.Mtx.Unlock()
}

func (boards_state *BoardsState) Fetch() error {
    data, err := get_static(BOARDS)

    if err != nil {
        return err
    }

    boards_state.Mtx.Lock()
    boards_state.Boards = nil
    err = json.Unmarshal(data, boards_state)
    boards_state.Mtx.Unlock()

    if err != nil {
        return err
    }
    return nil
}

func (boards_state *BoardsState) Print(f io.Writer) error {
    if len(boards_state.Boards) == 0 {
        err := boards_state.Fetch()
        if err != nil {
            return err
        }
    }

    for _, b := range boards_state.Boards {
        fmt.Fprintf(f, "%s\n", html.UnescapeString(b.Description))
    }
    return nil
}
