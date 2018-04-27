package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Conf is global toml conf struct
var Conf *TomlConfig

var version = "1.0.2"

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

func preMovie(d []Movie) error {
	var table string

	if len(d) != 0 {
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
			BccRecipients: Conf.Yts.BccRecipients,
			Message:       message,
		}

		err := mb.Send()
		if err != nil {
			return err
		}

		return nil
	}
	return fmt.Errorf("content not found")
}

func movie() error {

	if Conf.Yts.Enabled {
		res, err := GetNewMovies(Conf.Yts.PageLimit)
		if err != nil {
			return err
		}
		err = preMovie(res)
		if err != nil {
			return err
		}
	}

	return nil
}

func reportError(e string) {
	mb := MailBody{
		Subject:      Conf.Yts.Subject,
		ToRecipients: Conf.Yts.ToRecipients,
		Message:      e,
	}

	err := mb.Send()
	if err != nil {
		log.Printf("mail sending error %v", err)
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

	err := movie()
	if err != nil {
		reportError(err.Error())
	}
}
