package lsb

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
readSrc() считывает все байты из исходного файла
и возвращает срез с ними
*/
func readSrc(reader bufio.Reader) []byte {
	var data []byte = make([]byte, 0)
	for {
		tmp, err := reader.ReadByte()
		if err == io.EOF {
			return data
		} else if err != nil {
			fmt.Println("readSrc():", err)
			os.Exit(1)
		}

		data = append(data, tmp)
	}
}
