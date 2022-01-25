package main

/* TODO
 * Add an error view
 * Add an auto-update toggle
 * Find a way of caching the data.
 * Implement post/thread/board links
 * Rather than flat scrolling, scoll from post to post? XXX Is it worth? XXX
 */

import (
    "fmt"
    "github.com/jroimartin/gocui"
)

var ui_state UIState

type UIState struct {
    Boards              BoardsState
    Posts               PostsState
    Threads             ThreadsState
}

func (s *UIState) WaitUntilInitialized() {
    s.Boards.WaitGroup.Wait()
    s.Posts.WaitGroup.Wait()
    s.Threads.WaitGroup.Wait()
}

func main() {
    var boards_mgr  BoardsMgr
    var posts_mgr   PostsMgr
    var threads_mgr ThreadsMgr

    g, err := gocui.NewGui(gocui.Output256)

    if err != nil {
        fmt.Println(err)
        return
    }

    defer g.Close()

    g.Cursor = true

    ui_state.Boards.WaitGroup.Add(1)
    ui_state.Posts.WaitGroup.Add(1)
    ui_state.Threads.WaitGroup.Add(1)

    g.SetManager(boards_mgr, posts_mgr, threads_mgr)

    if err = key_binds(g); err != nil {
        fmt.Println(err)
    }

    go refresh_thread(g)


    err = g.MainLoop()

    if err != nil && err != gocui.ErrQuit {
        fmt.Println(err)
    }
}
