package main

import (
    "fmt"
    "io"
    "strings"
    "encoding/base64"
    "golang.org/x/net/html"
)

func process_comment(comment string) string {
    ret := ""
    tkn := html.NewTokenizer(strings.NewReader(comment))

    for {
        it := tkn.Next()
        if it == html.ErrorToken {
            break
        }

        if it == html.StartTagToken {
            tag := tkn.Token()
            if tag.Data == "span" {
                ret += "\033[32;1m"
            }

            if tag.Data == "a" {
                ret += "\033[31;4m"
            }

            if tag.Data == "br" {
                ret += "\n"
            }
        }

        if it == html.EndTagToken {
            ret += "\033[0m"
        }

        if it == html.TextToken {
            text := tkn.Token()
            ret += text.Data
        }
    }

    return ret
}

func print_id(f io.Writer, id string) {
    raw_data, err := base64.StdEncoding.DecodeString(id)
    if err != nil {
        return
    }

    if raw_data[0] == 232 || raw_data[0] == 0 {
        raw_data[0] += 6
    }
    fmt.Fprintf(f, "(ID: \033[38;5;%d;1m%s\033[0m) ", raw_data[0], id)
}
