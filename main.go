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

    html_head := html.MovieHtmlHead()

    for i := 0; i < len(d); i++ {
        data := d[i]

        year := data.Year
        title := data.Title
        genres := data.Genres
        rating := data.Rating
        id := data.ImdbID
        date_uploaded := data.DateUploaded
        cover := data.MediumCover

        body := html.MovieHtmlTable()
        html_table += fmt.Sprintf(body, cover, title, year, genres, rating, id, date_uploaded)
    }

        html_end := html.HtmlEnd()

        message := html_head + html_table + html_end
        mail.MailSend(message)
}

func Movies() {
    limit := toml.Yts()
    res, _ := api.GetNewMovies(limit.PageLimit)
    PreMovie(res)
}

func main() {
    Movies()
}

