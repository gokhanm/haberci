package main

import (
    "flag"
    "fmt"
    "os"

    "haberci/yts"
    "haberci/mail"
    "haberci/utils"
    "haberci/html"
)

var version = "0.1.2"

func PreMovie(d []api.Movie) {
    var html_table string

    html_head := yts_html.MovieHtmlHead()

    for i := 0; i < len(d); i++ {
        data := d[i]

        year := data.Year
        title := data.Title
        genres := data.Genres
        rating := data.Rating
        id := data.ImdbID
        date_uploaded := data.DateUploaded
        cover := data.MediumCover

        body := yts_html.MovieHtmlTable()
        html_table += fmt.Sprintf(body, cover, title, year, genres, rating, id, date_uploaded)
    }

    html_end := yts_html.HtmlEnd()
    yts_conf := toml.Yts()

    message := html_head + html_table + html_end
    mail.MailSend(yts_conf.Subject, yts_conf.ToRecipients, yts_conf.BccRecipients, message)
}

func Movies() {
    yts_conf := toml.Yts()

    if yts_conf.Enabled == "yes" {
        res, _ := api.GetNewMovies(yts_conf.PageLimit)
        PreMovie(res)
    }
}

func printVersion() {
    fmt.Printf("haberci version %v\n", version)
}

var helpMsg = `haberci = It will notify you of newly released movies or etc

usage: haberci [options]

options: 
`

func printHelp() {
	fmt.Println(helpMsg)
	flag.PrintDefaults()
}

func main() {
    var versionFlag = flag.Bool("v", false, "output version information and exit.")
    var helpFlag = flag.Bool("h", false, "display this help dialog")

    flag.Parse()

    if *versionFlag == true {
        printVersion()
        os.Exit(0)
    }

    if *helpFlag == true {
        printHelp()
        os.Exit(0)
    }

    Movies()
}

