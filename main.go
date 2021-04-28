package main

import (
	"khoros-community-cli/app"
	"khoros-community-cli/plumbing"
)

func main() {
	cmn := plumbing.Config()
	app.Run(cmn)
}
