package main

import (
	"3_1/wordz"
	"fmt"
	"github.com/Headrush95/utils"
	utilsV2 "github.com/Headrush95/utils/v2"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("Hello world")
	color.Blue("Hello world again")
	if isExistV2 := utilsV2.InSlice(wordz.Words, "Two"); isExistV2 {
		fmt.Println("Using utilsV2.InSlice and find value")
		return
	}
	if isExist := utils.Contains(wordz.Words, "Two"); isExist {
		fmt.Println("Slice Words contain finding value")
		return
	}
	if isExistInt := utils.ContainsInt([]int{1, 2, 3, 4, 5}, 5); isExistInt {
		fmt.Println("Slice Int contain finding value")
	}
	fmt.Println(wordz.Hello)
	for i := 0; i < 5; i++ {
		fmt.Println(wordz.Random())
	}
}
