package main

import (
    "sync"
    "github.com/jroimartin/gocui"
)

type BoardsMgr struct {

}

var brd_once sync.Once

func (boards_mgr BoardsMgr) Layout(g *gocui.Gui) error {
    defer brd_once.Do(ui_state.Boards.WaitGroup.Done)

    maxX, _ := g.Size()

    if view, err := g.SetView("boards", -1, -1, maxX, 5); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        ui_state.Boards.Print(view)
        view.Highlight = true
        view.Wrap = true
        view.SelBgColor = gocui.ColorGreen
        view.SelFgColor = gocui.ColorBlack
   }

   return nil
}
