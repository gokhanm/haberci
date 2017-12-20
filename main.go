package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Conf is global toml conf struct
var Conf *TomlConfig

var version = "1.0.0"

var helpMsg = `haberci = It will notify you of newly released movies or etc

usage: haberci [options]

options: 
`

func printVersion() {
	fmt.Printf("haberci version %v\n", version)
}

func printHelp() {
	fmt.Println(helpMsg)
	flag.PrintDefaults()
}

func preMovie(d []Movie) {
	var table string

	head := MovieHTMLHead()

	for i := 0; i < len(d); i++ {
		data := d[i]

		year := data.Year
		title := data.Title
		genres := data.Genres
		rating := data.Rating
		id := data.ImdbID
		dateUploaded := data.DateUploaded
		cover := data.MediumCover

		body := MovieHTMLTable()
		table += fmt.Sprintf(body, cover, title, year, genres, rating, id, dateUploaded)
	}

	htmlEnd := HTMLEnd()

	message := head + table + htmlEnd

	mb := MailBody{
		Subject:       Conf.Yts.Subject,
		ToRecipients:  Conf.Yts.ToRecipients,
		BccRecipients: Conf.Yts.BccRecipients,
		Message:       message,
	}

	err := mb.Send()
	if err != nil {
		log.Printf("Mail send error: %v", err)
	}
}

func movie() {

	if Conf.Yts.Enabled {
		res, _ := GetNewMovies(Conf.Yts.PageLimit)
		preMovie(res)
	}
}

func main() {
	var versionFlag = flag.Bool("v", false, "output version information and exit.")
	var helpFlag = flag.Bool("h", false, "display this help dialog")
	var confPath = flag.String("c", "/etc/haberci.toml", "config file path.")

	flag.Parse()

	if *versionFlag == true {
		printVersion()
		os.Exit(0)
	}

	if *helpFlag == true {
		printHelp()
		os.Exit(0)
	}

	Load(*confPath)
	Conf, _ = Parse()

	movie()
}
