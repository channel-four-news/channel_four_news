package main

import (
    "sync"
    "github.com/jroimartin/gocui"
)

type PostsMgr struct {

}

var pst_once sync.Once

func (posts_mgr PostsMgr) Layout(g *gocui.Gui) error {
    defer pst_once.Do(ui_state.Posts.WaitGroup.Done)

    maxX, maxY := g.Size()

    if view, err := g.SetView("body", X_LINE, Y_LINE, maxX, maxY); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }

        view.Wrap = true
        if _, err := g.SetCurrentView("body"); err != nil {
            return err
        }
    }
    return nil
}
