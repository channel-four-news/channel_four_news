package main

/* TODO
 * Add an error view
 * Add an auto-update toggle
 * Instead of autistic regex, the board object should be used
 * Find an efficient way of caching the data.
 * Implement post/thread/board links
 * Rather than flat scrolling, scoll from post to post? XXX Is it worth? XXX
 */

import (
    "fmt"
    "time"
    "github.com/jroimartin/gocui"
)

var ui_state UIState

type UIState struct {
    CurrentBoard        string
    DefaultBoard        string
    Initialized         bool
}

func main() {
    ui_state.DefaultBoard = "pol" // Eventually this will be read from a config file
    ui_state.CurrentBoard = ui_state.DefaultBoard
    ui_state.Initialized = false
    g, err := gocui.NewGui(gocui.Output256)

    if err != nil {
        fmt.Println(err)
        return
    }

    defer g.Close()

    g.Cursor = true
    g.SetManagerFunc(manager)

    if err = key_binds(g); err != nil {
        fmt.Println(err)
    }

    go func (g *gocui.Gui) {
        for !ui_state.Initialized {
            time.Sleep(1 * time.Second)
        }
        for {
            g.Update(func(g *gocui.Gui) error {
                err = _get_thread(g, false)
                if err != nil {
                    return err
                }
                return nil
            })
            if err != nil {
                return
            }
            time.Sleep(1 * time.Second)
        }
    }(g)

    if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        fmt.Println(err)
        return
    }
}
