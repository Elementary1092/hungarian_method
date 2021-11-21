package main

import (
	"flag"
	"fmt"

	alg "hungarian_method"
	in "input_data"
)

const (
	SOURCE       = "src"

	WEB          = "web"
	CMD          = "cmd"
	FILE         = "file"

	PATH         = "path"
	HOST         = "host"
	PORT         = "port"

	DEFAULT_HOST = "localhost"
	DEFAULT_PATH = "source.txt"
	DEFAULT_PORT = 8080

	SOURCE_DESC = "Source of data.\nPossible values: cmd, web, file"
	HOST_DESC   = "Host where data should be taken"
	PATH_DESC   = "Path of the file where data is stored"
	PORT_DESC   = "Port which should be listened"
)

var source string
var host   string
var path   string
var port   int

func init() {
	flag.StringVar(&source, SOURCE, CMD, SOURCE_DESC)
	flag.StringVar(&host, HOST, DEFAULT_HOST, HOST_DESC)
	flag.StringVar(&path, PATH, DEFAULT_PATH, PATH_DESC)
	flag.IntVar(&port, PORT, DEFAULT_PORT, PORT_DESC)
}

func main() {
	flag.Parse()

	var inputMethod in.DataInput
	switch source {
	case CMD:
		inputMethod = &in.CMDInput{}
	case WEB:
		inputMethod = &in.WebInput{
			Host: host,
			Port: port,
		}
	case FILE:
		inputMethod = &in.FileInput{
			Path: path,
		}
	}

	data := inputMethod.GetData()

	t := alg.NewTable(data)
	fmt.Println(t.Solve())
}