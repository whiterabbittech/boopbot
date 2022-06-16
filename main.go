package main

import (
	"fmt"

	"github.com/whiterabbittech/boopbot/packages/givebutter/contribution"
)

func main() {
	var input = contribution.ScanContributionsReq{}
	var output, err = contribution.Main(input)
	if err != nil {
		fmt.Println(err)
	}
	_ = output
}
