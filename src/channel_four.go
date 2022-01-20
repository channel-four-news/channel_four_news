package main

import (
    "net/http"
    "golang.org/x/net/html"
    "encoding/json"
    "encoding/base64"
    "fmt"
    "io"
    "io/ioutil"
    "strings"
)

var (
    STATIC_CDN = "a.4cdn.org"
    MEDIA_CON = "i.4cdn.org"
    THREADS = "threads.json"
    BOARDS = "boards.json"
)

type FourBoard struct {
    Board           string          `json:"board"`
    Title           string          `json:"title"`
    Description     string          `json:"meta_description"`
}

type FourBoards struct {
    Boards          []FourBoard     `json:"boards"`
}

type FourThread struct {
    No              uint            `json:"no"`
    LastModified    uint            `json:"last_modified"`
    Replies         uint            `json:"replies"`
}

type FourThreadPage struct {
    Page            uint            `json:"page"`
    Threads         []FourThread    `json:"threads"`
}

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

type FourPosts struct {
    Posts           []FourPost      `json:"posts"`
}

func GetStatic(uri string) ([]byte, error) {
    response, err := http.Get("https://"+STATIC_CDN+"/"+uri)

    if err != nil {
        return nil, err
    }

    defer response.Body.Close()

    return ioutil.ReadAll(response.Body)
}

func GetThreads(board string) ([]FourThreadPage, error) {
    pages := make([]FourThreadPage, 10)

    data, err := GetStatic(board+"/"+THREADS)

    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(data, &pages)

    if err != nil {
        return nil, err
    }
    return pages, nil
}

func GetBoards() (*FourBoards, error) {
    boards := new(FourBoards)

    data, err := GetStatic(BOARDS)

    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(data, &boards)

    if err != nil {
        return nil, err
    }
    return boards, nil
}

func GetThread(board string, thread string) (*FourPosts, error) {
    posts := new(FourPosts)

    uri := fmt.Sprintf("%s/thread/%s.json", board, thread)
    data, err := GetStatic(uri)

    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(data, &posts)

    if err != nil {
        return nil, err
    }
    return posts, nil
}

func ProcessComment(comment string) string {
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

func PrintId(f io.Writer, id string) {
    raw_data, err := base64.StdEncoding.DecodeString(id)
    if err != nil {
        return
    }

    if raw_data[0] == 232 || raw_data[0] == 0 {
        raw_data[0] += 6
    }
    fmt.Fprintf(f, "(ID: \033[38;5;%d;1m%s\033[0m) ", raw_data[0], id)
}

func PrintThread(f io.Writer, thread string, board string) {
    posts, err := GetThread(board, thread)

    if err != nil {
        fmt.Fprintln(f, err)
        return
    }

    for _, p := range posts.Posts {
        if p.Sub != "" {
            fmt.Fprintf(f, "\033[34;1m%s\033[0m ", html.UnescapeString(p.Sub))
        }
        fmt.Fprintf(f, "\033[32m%s\033[0m ", p.Name)
        if p.Trip != "" {
            fmt.Fprintf(f, "\033[32m%s\033[0m ", p.Trip)
        }
        if p.Id != "" {
            PrintId(f, p.Id)
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
                        MEDIA_CON, ui_state.CurrentBoard,
                        p.ImageId, p.Ext)
        }
        fmt.Fprintf(f, "%s\n\n", ProcessComment(html.UnescapeString(p.Comment)))
    }
}

func PrintBoards(f io.Writer) {
    boards, err := GetBoards()
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, b := range boards.Boards {
        fmt.Fprintf(f, "%s\n", html.UnescapeString(b.Description))
    }
}

func PrintThreads(f io.Writer, board string) {
    pages, err := GetThreads(board)
    if err != nil {
        fmt.Fprintln(f, err)
        return
    }

    for _, p := range pages {
        for _, t := range p.Threads {
//            fmt.Fprintln(f, "Thread #", t.No, "\n\t", t.LastModified, "|", t.Replies)
            fmt.Fprintln(f, t.No)
        }
    }
}
