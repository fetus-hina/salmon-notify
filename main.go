package main

import (
	"errors"
	"fmt"
	"github.com/fetus-hina/salmon-notify/conf"
	"github.com/fetus-hina/salmon-notify/discord"
	"github.com/fetus-hina/salmon-notify/s2ink"
	"net/url"
	"os"
	"strings"
)

func main() {
	err := doMain()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	os.Exit(0)
}

func doMain() error {
	urlOrOption, err := getCliOptions()
	switch {
	case err != nil:
		return err

	case urlOrOption == "--help" || urlOrOption == "-f":
		fmt.Fprintln(os.Stderr, usage())
		return nil

	case urlOrOption == "--version" || urlOrOption == "-v":
		fmt.Fprintln(os.Stderr, conf.CliVersion())
		return nil
	}

	schedules, err := s2ink.FetchCoopSchedules()
	if err != nil {
		return err
	}

	for _, detail := range schedules.Details {
		if urlOrOption == "--stdout" {
			fmt.Println(discord.FormatMessage(detail))
			return nil
		}

		if detail.IsOpened() && detail.IsOpenRecently() {
			webhook := discord.Webhook{
				UserName: "サーモンランスケジュール",
				Content:  discord.FormatMessage(detail),
			}
			err = webhook.Post(urlOrOption)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getCliOptions() (string, error) {
	if len(os.Args) != 2 {
		return "", errors.New(usage())
	}

	file := strings.Trim(os.Args[1], " ")
	if file == "-v" || file == "--version" || file == "-h" || file == "--help" || file == "--stdout" || isValidURL(file) {
		return file, nil
	}

	return "", errors.New(usage())
}

func usage() string {
	return fmt.Sprintf(
		"%s <webhook-url>\n"+
			"\n"+
			"Usage:\n"+
			"  -h --help     Show help\n"+
			"  -v --version  Show version\n"+
			"     --stdout   Write message to STDOUT",
		os.Args[0],
	)
}

func isValidURL(urlStr string) bool {
	u, err := url.ParseRequestURI(urlStr)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
