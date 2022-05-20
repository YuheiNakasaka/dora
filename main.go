package main

import (
	"embed"

	"github.com/YuheiNakasaka/dora/cmd"
)

//go:embed resources/*.mp3
var dora embed.FS

func main() {
	cmd.Execute(dora)
}
