package main

import (
	"fmt"
	"github.com/PichuChen/tw_bussiness_tax-golang"
	"os"
)

func main() {
	fmt.Println("Ho" + btax.EXT_BUSINESS_INPUT_OUTPUT_FILE)

	f, err := os.Open("54834795.TET")
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	records := btax.ParseTETFile(f)
	for _, r := range records {
		fmt.Println(r)
	}
}

func runTXTFile() {
	f, err := os.Open("54834795.TXT")
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	records := btax.ParseTXTFile(f)
	for _, r := range records {
		fmt.Println(r)
	}

}
