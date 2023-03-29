package main26

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	logTime := func() {
		fmt.Println(time.Since(start))
	}
	defer logTime()
	f, err := os.Open("./in.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	scn := bufio.NewScanner(f)
	i, line, countOfBytes := 0, "", 0
	fOut, err := os.OpenFile("out.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(fOut)

	for scn.Scan() {
		i++
		line = scn.Text()
		n, err := writer.WriteString(fmt.Sprintf("%d - %s\n", i, line))
		if err != nil {
			panic(err)
		}
		countOfBytes += n
	}
	err = writer.Flush()
	if err != nil {
		fOut.Close()
		panic(err)
	}
	err = fOut.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("its ok, total lines - %d, total bytes - %d\n", i, countOfBytes)
}
