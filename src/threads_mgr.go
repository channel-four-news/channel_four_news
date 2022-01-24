package main

import (
    "sync"
    "github.com/jroimartin/gocui"
)

type ThreadsMgr struct {

}

var thr_once sync.Once

func (threads_mgr ThreadsMgr) Layout(g *gocui.Gui) error {
    defer thr_once.Do(ui_state.Threads.WaitGroup.Done)

    _, maxY := g.Size()
    if view, err := g.SetView("threads", -1, Y_LINE, X_LINE, maxY); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }

        ui_state.Boards.WaitGroup.Wait()
        ui_state.Posts.WaitGroup.Wait()

        ui_state.Threads.Print(view, ui_state.Boards.GetCurrentTitle())

        view.Highlight = true
        view.Wrap = true
        view.SelBgColor = gocui.ColorGreen
        view.SelFgColor = gocui.ColorBlack
        get_thread(g, view)
    }
    return nil
}

