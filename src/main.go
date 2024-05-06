package main

import (
	"embed"

	"github.com/GleisonEm/bot-claudio-zap-golang/cmd"
)

//go:embed views/index.html
var embedIndex embed.FS

//go:embed views
var embedViews embed.FS

func main() {
	cmd.Execute(embedIndex, embedViews)
}
