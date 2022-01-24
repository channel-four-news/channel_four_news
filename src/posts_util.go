package main

import (
    "fmt"
    "io"
    "golang.org/x/net/html"
)

type FourPost struct {
    No              uint            `json:"no"`
    Now             string          `json:"now"`
    Name            string          `json:"name"`
    Comment         string          `json:"com"`
    Sub             string          `json:"sub"`
    Id              string          `json:"id"`
    CountryName     string          `json:"country_name"`
    FlagName        string          `json:"flag_name"`
    Trip            string          `json:"trip"`
    Filename        string          `json:"filename"`
    Ext             string          `json:"ext"`
    ImageId         uint            `json:"tim"`
}

func (p *FourPost) Print(f io.Writer) {
    if p.Sub != "" {
        fmt.Fprintf(f, "\033[34;1m%s\033[0m ", html.UnescapeString(p.Sub))
    }
    fmt.Fprintf(f, "\033[32m%s\033[0m ", p.Name)
    if p.Trip != "" {
        fmt.Fprintf(f, "\033[32m%s\033[0m ", p.Trip)
    }
    if p.Id != "" {
        print_id(f, p.Id)
    }
    if p.CountryName != "" {
        fmt.Fprintf(f, "(%s) ", p.CountryName)
    }
    if p.FlagName != "" {
        fmt.Fprintf(f, "(%s) ", p.FlagName)
    }
    fmt.Fprintf(f, "%s No.%d\n", p.Now, p.No)
    if p.ImageId != 0 {
        fmt.Fprintf(f,"\033[38;5;24mhttps://%s/%s/%d%s\n\033[0m",
                    MEDIA_CON, ui_state.Boards.GetCurrentTitle(),
                    p.ImageId, p.Ext)
    }
    fmt.Fprintf(f, "%s\n\n", process_comment(html.UnescapeString(p.Comment)))
}
