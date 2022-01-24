package main

import (
    "io"
    "fmt"
    "sync"
    "encoding/json"
)

var (
    THREADS = "threads.json"
)

type FourThread struct {
    No              uint            `json:"no"`
    LastModified    uint            `json:"last_modified"`
    Replies         uint            `json:"replies"`
}

type FourThreadPage struct {
    Page            uint            `json:"page"`
    Threads         []FourThread    `json:"threads"`
}

type ThreadsState struct {
    Pages           []FourThreadPage
    WaitGroup       sync.WaitGroup
    Mtx             sync.Mutex
}

func (threads_state *ThreadsState) Fetch(board string) error {
    data, err := get_static(board+"/"+THREADS)

    if err != nil {
        return err
    }

    threads_state.Mtx.Lock()
    threads_state.Pages = nil
    err = json.Unmarshal(data, &threads_state.Pages)
    threads_state.Mtx.Unlock()

    if err != nil {
        return err
    }
    return nil
}

func (threads_state *ThreadsState) Print(f io.Writer, board string) {
    err := threads_state.Fetch(board)
    if err != nil {
        fmt.Fprintln(f, err)
        return
    }

    for _, p := range threads_state.Pages {
        for _, t := range p.Threads {
            fmt.Fprintln(f, t.No)
        }
    }
}
