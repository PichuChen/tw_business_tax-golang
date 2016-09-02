package main

import (
	"fmt"
	"github.com/PichuChen/tw_bussiness_tax-golang"
)

func main() {
	fmt.Println("Ho" + btax.EXT_BUSINESS_INPUT_OUTPUT_FILE)
	btax.ParseTXTFile()
}
