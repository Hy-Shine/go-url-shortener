package main

import (
	"fmt"
	"os"

	"github.com/hy-shine/go-url-shortener/shortener"
	"github.com/hy-shine/go-url-shortener/store"
)

const helpInfo string = `Usage: ./client [OPTION] [URL]
	for example:
		./client -g https://www.google.com

	-h --help	show the client help info
	-g		generate a shorter url by the client
	-v --version	show the client version info`

func help(args []string) string {
	arg := args[0]
	if arg[0:1] != "-" {
		return fmt.Sprintf("unrecognized option '%s'\ntry './client -h' for more infomation.", arg)
	} else if arg[0:2] == "--" {
		if len(arg) == 2 {
			return "invalid option '--'\ntry './client -h' for more infomation."
		} else if arg[2:] != "help" && arg[2:] != "version" {
			return fmt.Sprintf("unrecognized option '%s'\ntry './client -h' for more infomation.", arg)
		}
	} else if arg[0:1] == "-" {
		if len(arg) == 1 {
			return "invalid option '-'\ntry './client -h' for more infomation."
		} else if arg[1:] != "h" && arg[1:] != "v" && arg[1:] != "g" {
			return fmt.Sprintf("invalid option -- '%s'\ntry './client -h' for more infomation.", arg[2:])
		}
	}
	switch arg {
	case "-h", "--help":
		return helpInfo
	case "-v", "--version":
		return "The client version: 1.0"
	case "-g":
		if len(args) == 1 {
			return "Missing a long url"
		}
		shortURL := shortener.GenerateShortLink(args[1])
		store.SaveURLMapping(shortURL, args[1])
		return "http://localhost/" + shortURL
	}
	return "Unkown error"
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println(helpInfo)
		return
	}
	info := help(args[1:])
	fmt.Println(info)
}
