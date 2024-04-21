package pango

import (
	"fmt"
)

type Args struct {
	Bg   string
	Fg   string
	Size string
}

func Powerline(separator, text, fg, bg string, isFirst bool) string {
	edgeBg := fg

	if isFirst {
		edgeBg = "#000000"
	}

	edge := Span(separator, Args{
		Fg:   bg,
		Bg:   edgeBg,
		Size: "15pt",
	})

	content := Span(text, Args{
		Bg:   bg,
		Fg:   fg,
		Size: "14pt",
	})

	return edge + content
}

func Span(s string, a Args) string {
	argString := ""

	if a.Bg != "" {
		argString += fmt.Sprintf("background=\"%s\"", a.Bg)
	}
	if a.Fg != "" {
		argString += fmt.Sprintf("foreground=\"%s\"", a.Fg)
	}
	if a.Size != "" {
		argString += fmt.Sprintf("size=\"%s\"", a.Size)
	}

	return fmt.Sprintf("<span %s>%s</span>", argString, s)
}
