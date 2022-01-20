package main

import (
    "regexp"
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
    PrintThread(body_view, thread, ui_state.CurrentBoard)

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

func get_board_name(line string) string {
    reg := regexp.MustCompile(`"/([[:alnum:]]+)/`)
    sub := reg.FindStringSubmatch(line)

    if len(sub) == 2 {
        return sub[1]
    }
    return ui_state.DefaultBoard
}

func get_board(g *gocui.Gui, v *gocui.View) error {
    _, y := v.Cursor()

    board_line, err := v.Line(y)
    if err != nil {
        return err
    }

    board := get_board_name(board_line)
    err = display_board(g, board)
    if err != nil {
        return err
    }
    return nil
}

func update_board(g *gocui.Gui, v *gocui.View) error {
    return display_board(g, ui_state.CurrentBoard)
}

func display_board(g *gocui.Gui, board_name string) error {
    threads_view, err := g.View("threads")

    if err != nil {
        return err
    }

    threads_view.Clear()
    PrintThreads(threads_view, board_name)
    threads_view.SetOrigin(0, 0)
    threads_view.SetCursor(0, 0)
    ui_state.CurrentBoard = board_name
    get_thread(g, threads_view)
    return nil
}
