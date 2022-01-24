package main

import (
    "io"
    "fmt"
    "sync"
    "encoding/json"
)

var (
    MEDIA_CON = "i.4cdn.org"
)

type PostsState struct {
    Posts       []FourPost      `json:"posts"`
    WaitGroup   sync.WaitGroup
    Mtx         sync.Mutex
}

func (posts_state *PostsState) Fetch(board string, thread string) error {
    uri := fmt.Sprintf("%s/thread/%s.json", board, thread)
    data, err := get_static(uri)

    if err != nil {
        return err
    }

    posts_state.Mtx.Lock()
    posts_state.Posts = nil
    err = json.Unmarshal(data, posts_state)
    posts_state.Mtx.Unlock()

    if err != nil {
        return err
    }

    return nil
}

func (posts_state *PostsState) Print(f io.Writer, thread string, board string) {
    err := posts_state.Fetch(board, thread)

    if err != nil {
        fmt.Fprintln(f, err)
        return
    }

    for _, p := range posts_state.Posts {
        p.Print(f)
    }
}
