package main27

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func showFile(f io.Reader) {
	fmt.Println("[showFile] start")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Printf("%v\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println("[showFile] end")
}

func main() {
	defer func() {
		file, err := os.OpenFile("2.7/out.txt", os.O_RDONLY, 0666)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		if recPan := recover(); recPan != nil {
			switch recPan {
			case nil:
				return
			default:
				showFile(file)
			}
		}
	}() //Если некорректные данные выводит те, что записались
	f, err := os.Open("2.7/in.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	//reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(f)

	fOut, err := os.OpenFile("2.7/out.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fOut.Close()

	//line := ""
	countRow := 0
	var tmp []string

	writer := bufio.NewWriter(fOut)
	defer func() {
		err := writer.Flush()
		if err != nil {
			panic(err)
		}
	}() // flush to out

	for scanner.Scan() {
		countRow++
		tmp = strings.Split(scanner.Text(), "|")
		if len(tmp) != 3 {
			panic(fmt.Sprintf("Parse error: empty field on string %d", countRow))
		}
		if countRow == 100 {
			fmt.Println(tmp[0])
		}
		_, err = writer.WriteString(fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", countRow, tmp[0], tmp[1], tmp[2]))

		if err != nil {
			panic(err)
		}
		//line, err = reader.ReadString('\n')
		//
		//if err == io.EOF {
		//	fmt.Println(countRow)
		//	break
		//}
		//if err != nil {
		//	panic(err)
		//}
		//tmp = strings.Split(line, "|")
		//if len(tmp) != 3 {
		//	panic(fmt.Sprintf("Parse error: empty field on string %d", countRow))
		//}
		//_, err = writer.WriteString(fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", countRow, tmp[0], tmp[1], tmp[2]))
		//
	}
	fmt.Println(countRow)
}
