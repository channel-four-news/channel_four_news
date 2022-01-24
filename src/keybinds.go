package main

import (
    "github.com/jroimartin/gocui"
)

var (
    X_LINE = 15
    Y_LINE = 5
)

func switch_view(g *gocui.Gui, v *gocui.View) error {
    if v == nil {
        _, err := g.SetCurrentView("boards")
        return err
    }
    if v.Name() == "body" {
        _, err := g.SetCurrentView("boards")
        return err
    }

    if v.Name() == "boards" {
        _, err := g.SetCurrentView("body")
        return err
    }
    return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
    return gocui.ErrQuit
}

func set_update_binds(g *gocui.Gui) error {
    err := g.SetKeybinding("threads", gocui.KeyEnter, gocui.ModNone, get_thread)

    if err != nil {
        return err
    }

    err = g.SetKeybinding("boards", gocui.KeyEnter, gocui.ModNone, get_board)

    if err != nil {
        return err
    }

    err = g.SetKeybinding("", gocui.KeyCtrlR, gocui.ModNone, update_board)

    if err != nil {
        return err
    }

    return nil
}

func key_binds(g *gocui.Gui) error {
    if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
        return err
    }

    if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, switch_view); err != nil {
        return err
    }

    if err := set_cursor_binds(g); err != nil {
        return err
    }

    if err := set_update_binds(g); err != nil {
        return err
    }

    if err := set_help_binds(g); err != nil {
        return err
    }

    return nil
}
