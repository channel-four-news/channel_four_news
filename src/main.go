package main

/* TODO
 * Add an error view
 * Add an auto-update toggle
 * Implement caching
 * Implement post/thread/board links
 */

import (
    "fmt"
    "github.com/jroimartin/gocui"
)

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

    ui_state.LockViews()

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
