package main

import (
    "time"
    "net/http"
    "io/ioutil"
    "github.com/jroimartin/gocui"
)

var (
    STATIC_CDN = "a.4cdn.org"
    THREAD_REFRESH_RATE = 30 * time.Second
)

func get_static(uri string) ([]byte, error) {
    response, err := http.Get("https://"+STATIC_CDN+"/"+uri)

    if err != nil {
        return nil, err
    }

    defer response.Body.Close()

    return ioutil.ReadAll(response.Body)
}

func _refresh_thread(g *gocui.Gui) error {
    err := _get_thread(g, false)
    if err != nil {
        return err
    }
    return nil
}

func refresh_thread(g *gocui.Gui) {
    ui_state.WaitUntilInitialized()

    for {
        g.Update(_refresh_thread)
        time.Sleep(THREAD_REFRESH_RATE)
    }
}

func assert(cond bool, str string) {
    if (!(cond)) {
        panic(str)
    }
}
