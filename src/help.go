package main

import (
    "fmt"
    "github.com/jroimartin/gocui"
)
 
var saved_view string

var (
    HELP_LINES = 9
    HELP_STRING = "                            Help\n" +
                  "==========================================================\n" +
                  "F1                            - Show/Hide Help\n" +
                  "Enter                         - Select the board\n" +
                  "Tab                           - Switch between panels\n" +
                  "Left/Right Arrow              - Prev/Next Thread\n" +
                  "Ctrl+P, Ctrl+N, Up/Down Arrow - Scrolling\n" +
                  "Ctrl+R                        - Refresh the current board\n" +
                  "Ctrl+C                        - Exit"
)

func show_help(g *gocui.Gui, v *gocui.View) error {

    saved_view = v.Name()
    maxX, maxY := g.Size()
    hv, err := g.SetView("help", maxX/2-29, maxY/2-HELP_LINES, maxX/2+30, maxY/2)
    if err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        fmt.Fprintf(hv, HELP_STRING)

        if _, err = g.SetCurrentView("help"); err != nil {
            return err
        }
    }
    return nil
}

func close_help(g *gocui.Gui, v *gocui.View) error {
    if err := g.DeleteView("help"); err != nil {
        return err
    }

    if _, err := g.SetCurrentView(saved_view); err != nil {
        return err
    }
    return nil
}

func set_help_binds(g *gocui.Gui) error {
    err := g.SetKeybinding("boards", gocui.KeyF1, gocui.ModNone, show_help)

    if err != nil {
        return err
    }

    err = g.SetKeybinding("body", gocui.KeyF1, gocui.ModNone, show_help)

    if err != nil {
        return err
    }

    err = g.SetKeybinding("threads", gocui.KeyF1, gocui.ModNone, show_help)

    if err != nil {
        return err
    }

    err = g.SetKeybinding("help", gocui.KeyF1, gocui.ModNone, close_help)

    if err != nil {
        return err
    }

    return nil
}
