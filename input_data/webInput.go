package input_data

import (
	"fmt"
	//"io"
	"net"
)

type WebInput struct {
	Host string
	Port int
}

func (w *WebInput) GetData() [][]int64 {
	_, err := net.Listen("tcp", fmt.Sprintf("%s:%s", w.Host))

	if err != nil {
		fmt.Printf("Connection to the %s is aborted", w.Host)
		panic(err)
	}



	return nil
}