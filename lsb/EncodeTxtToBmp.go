package lsb

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	_ "strings"
)

/*
encodeTxtBmp() зашифровывает текст из переданного утилите
txt-файла и переданную утилите bmp-изображение

srcBmpPath - путь к исходному bmp-файлу

srcTxt - путь к txt-файлу, который нужно зашифровать в bmp-файл
*/
func EncodeTxtToBmp(srcBmpPath string, srcTxtPath string) {
	srcBmp, err := os.Open(srcBmpPath)
	if err != nil {
		fmt.Printf("Failed open %s\n", srcBmpPath)
		os.Exit(1)
	}
	defer srcBmp.Close()

	srcTxt, err := os.OpenFile(srcTxtPath, os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Printf("Failed open %s\n", srcTxtPath)
		os.Exit(1)
	}
	defer srcTxt.Close()

	destBmp, err := os.OpenFile("dest.bmp", os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println("Failed open")
		os.Exit(1)
	}
	defer destBmp.Close()

	readerBmp := bufio.NewReader(srcBmp)
	readerTxt := bufio.NewReader(srcTxt)
	writerBmp := bufio.NewWriter(destBmp)

	pixels := readSrc(*readerBmp)
	data := readSrc(*readerTxt)
	var index int = BYTE_OFFSET
	for _, value := range data {
		tmp := fmt.Sprintf("%08b", value)
		for i := 0; i < len(tmp); i++ {
			str := fmt.Sprintf("%08b", pixels[index])[0:7] + string(tmp[i])
			backToByte, _ := strconv.ParseUint(str, 2, 8)
			pixels[index] = byte(backToByte)
			index++
		}
	}
	writerBmp.Write(pixels)

	fmt.Println("encodeTxtToBmp(): Done")
}
