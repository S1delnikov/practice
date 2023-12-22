package main

import (
	"fmt"
	"os"

	"github.com/misha/lsb"
)

func main() {
	if len(os.Args) != 6 {
		fmt.Println("Usage: ./main <src.bmp> <src.txt> <dest.bmp> <dest.txt> <delimiter>")
		return
	}
	srcBmpPath := os.Args[1]
	srcTxtPath := os.Args[2]
	destBmpPath := os.Args[3]
	destTxtPath := os.Args[4]
	delimiter := os.Args[5]

	lsb.EncodeTxtToBmp(srcBmpPath, srcTxtPath)
	lsb.DecodeTxtFromBmp(destTxtPath, destBmpPath, delimiter)
}
