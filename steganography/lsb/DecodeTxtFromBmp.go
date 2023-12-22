package lsb

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
decodeTxtFromBmp() извлекает закодированное в .bmp файл
текстовое сообщение

destTxtPath - путь к txt-файлу, в который будут извлекаться данные

encodedBmpPath - путь к bmp-файлу, содержащему зашифрованную информацию

delim - последовательность из одинаковых символом, встретив которую,
функция будет считать, что вся нужная информация была изввлечена
*/
func DecodeTxtFromBmp(destTxtPath string, encodedBmpPath string, delim string) {
	encodedBmp, err := os.Open(encodedBmpPath)
	if err != nil {
		fmt.Println("decodeTxtFromBmp():", err)
		os.Exit(1)
	}
	defer encodedBmp.Close()

	destTxt, err := os.OpenFile(destTxtPath, os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println("decodeTxtFromBmp():", err)
		os.Exit(1)
	}
	defer destTxt.Close()

	reader := bufio.NewReader(encodedBmp)
	writer := bufio.NewWriter(destTxt)
	var str string = ""
	var del string = ""
	pixels := readSrc(*reader)[BYTE_OFFSET:]
	for _, value := range pixels {
		str += string(fmt.Sprintf("%08b", value)[7])
		if del == delim {
			writer.Flush()
			break
		}
		if len(str) == 8 {
			backToByte, _ := strconv.ParseUint(str, 2, 8)
			if backToByte == uint64(delim[0]) {
				del += string(rune(backToByte))
			} else {
				del = ""
			}
			writer.WriteString(fmt.Sprintf("%c", backToByte))
			str = ""
		}
	}
	fmt.Println("decodeTxtFromBmp: Done")
}
