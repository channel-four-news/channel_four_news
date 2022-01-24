package main

import (
    "github.com/jroimartin/gocui"
)

func _get_thread(g *gocui.Gui, rewind bool) error {
    thread_view, err := g.View("threads")

    if err != nil {
        return err
    }

    _, y := thread_view.Cursor()

    thread, err := thread_view.Line(y)
    if err != nil {
        return err
    }

    body_view, err := g.View("body")

    if err != nil {
        return err
    }

    body_view.Clear()
    ui_state.Posts.Print(body_view, thread, ui_state.Boards.GetCurrentTitle())

    if rewind {
        body_view.SetOrigin(0, 0)
    }

    return nil
}

func update_thread(g *gocui.Gui, v *gocui.View) error {
    return _get_thread(g, false)
}

func get_thread(g *gocui.Gui, v *gocui.View) error {
    return _get_thread(g, true)
}

func get_board(g *gocui.Gui, v *gocui.View) error {
    _, cy := v.Cursor()
    _, oy := v.Origin()
    y := oy + cy

    board := ui_state.Boards.GetTitle(y)

    ui_state.Boards.SetCurrentBoard(y)

    err := display_board(g, board)

    if err != nil {
        return err
    }

    return nil
}

func update_board(g *gocui.Gui, v *gocui.View) error {
    return display_board(g, ui_state.Boards.GetCurrentTitle())
}

func display_board(g *gocui.Gui, board_name string) error {
    threads_view, err := g.View("threads")

    if err != nil {
        return err
    }

    threads_view.Clear()
    ui_state.Threads.Print(threads_view, board_name)
    threads_view.SetOrigin(0, 0)
    threads_view.SetCursor(0, 0)
    get_thread(g, threads_view)
    return nil
}
