package main

import (
    "github.com/jroimartin/gocui"
)

func cursor_up(g *gocui.Gui, v *gocui.View) error {

    if v != nil {
        v.MoveCursor(0, -1, false)
    }
    return nil
}

func cursor_down(g *gocui.Gui, v *gocui.View) error {
    if v != nil {
        v.MoveCursor(0, 1, false)
    }
    return nil
}

func next_thread(g *gocui.Gui, v *gocui.View) error {
    thread_view, err := g.View("threads")

    if err != nil {
        return err
    }

    if err = cursor_down(g, thread_view); err != nil {
        return err
    }

    if err = get_thread(g, thread_view); err != nil {
        return err
    }

    return nil
}

func prev_thread(g *gocui.Gui, v *gocui.View) error {
    thread_view, err := g.View("threads")

    if err != nil {
        return err
    }

    if err = cursor_up(g, thread_view); err != nil {
        return err
    }

    if err = get_thread(g, thread_view); err != nil {
        return err
    }

    return nil
}

func scroll(shift int, v *gocui.View) error {
    if v == nil {
        return nil
    }

    x, y := v.Origin()

    if y == 0 && shift < 0 {
        return nil
    }

    if err := v.SetOrigin(x, y+shift); err != nil {
        return err;
    }

    return nil
}

func scroll_down(g *gocui.Gui, v *gocui.View) error {
    return scroll(1, v)
}

func scroll_up(g *gocui.Gui, v *gocui.View) error {
    return scroll(-1, v)
}

func set_cursor_binds(g *gocui.Gui) error {
    err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone, next_thread)

    if err != nil {
        return err
    }

    err = g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone, prev_thread)

    if err != nil {
        return err
    }

    err = g.SetKeybinding("boards", gocui.KeyArrowDown, gocui.ModNone, cursor_down)

    if err != nil {
        return err
    }

    err = g.SetKeybinding("boards", gocui.KeyArrowUp, gocui.ModNone, cursor_up)

    if err != nil {
        return err
    }

    err = g.SetKeybinding("boards", gocui.KeyCtrlN, gocui.ModNone, cursor_down)

    if err != nil {
        return err
    }

    err = g.SetKeybinding("boards", gocui.KeyCtrlP, gocui.ModNone, cursor_up)

    if err != nil {
        return err
    }

    err = g.SetKeybinding("body", gocui.KeyCtrlN, gocui.ModNone, scroll_down)

    if err != nil {
        return err
    }

    err = g.SetKeybinding("body", gocui.KeyCtrlP, gocui.ModNone, scroll_up)

    if err != nil {
        return err
    }

    err = g.SetKeybinding("body", gocui.KeyArrowDown, gocui.ModNone, scroll_down)

    if err != nil {
        return err
    }

    err = g.SetKeybinding("body", gocui.KeyArrowUp, gocui.ModNone, scroll_up)

    if err != nil {
        return err
    }

    return nil
}
