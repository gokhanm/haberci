package main

import (
    "fmt"
    "haberci/yts"
    "haberci/mail"
    "haberci/utils"
    "haberci/html"
)

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

func main() {
    Movies()
}

