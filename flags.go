package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func declareFlags(prefix string) *flags {
	desc := func(message string) string {
		if prefix == "" {
			return fmt.Sprintf("%s.", message)
		}
		return fmt.Sprintf("%s for std%s.", message, prefix)
	}

	name := func(x string) string {
		if prefix == "" {
			return x
		}
		return fmt.Sprintf("%s-%s", prefix, x)
	}

	return &flags{
		pattern: flag.String(
			name("pattern"),
			"",
			desc("Restrict colorization to lines with this regex pattern"),
		),
		fg:     flag.String(name("fg"), "", desc("Foreground color")),
		bg:     flag.String(name("bg"), "", desc("Background color")),
		bold:   flag.Bool(name("bold"), false, desc("Use bold text")),
		faint:  flag.Bool(name("faint"), false, desc("Use faint text")),
		italic: flag.Bool(name("italic"), false, desc("Use italic text")),
		underline: flag.Bool(
			name("underline"),
			false,
			desc("Use underlined text"),
		),
		blinkSlow: flag.Bool(
			name("blink-slow"),
			false,
			desc("Use slowly blinking text"),
		),
		blinkRapid: flag.Bool(
			name("blink-rapid"),
			false,
			desc("Use rapidly blinking text"),
		),
		reverseVideo: flag.Bool(
			name("reverse-video"),
			false,
			desc("Show text in reverse video"),
		),
		concealed: flag.Bool(
			name("concealed"),
			false,
			desc("Use concealed text"),
		),
		crossedOut: flag.Bool(
			name("crossed-out"),
			false,
			desc("Use crossed out text"),
		),
	}
}

type flags struct {
	pattern      *string
	fg           *string
	bg           *string
	bold         *bool
	faint        *bool
	italic       *bool
	underline    *bool
	blinkSlow    *bool
	blinkRapid   *bool
	reverseVideo *bool
	concealed    *bool
	crossedOut   *bool
}

func (f *flags) ApplyTo(c *maybeColor) *maybeColor {
	c = c.Add(fgAttribute(*f.fg)).Add(bgAttribute(*f.bg))
	if *f.bold {
		c = c.Add(color.Bold)
	}
	if *f.faint {
		c = c.Add(color.Faint)
	}
	if *f.italic {
		c = c.Add(color.Italic)
	}
	if *f.underline {
		c = c.Add(color.Underline)
	}
	if *f.blinkSlow {
		c = c.Add(color.BlinkSlow)
	}
	if *f.blinkRapid {
		c = c.Add(color.BlinkRapid)
	}
	if *f.reverseVideo {
		c = c.Add(color.ReverseVideo)
	}
	if *f.concealed {
		c = c.Add(color.Concealed)
	}
	if *f.crossedOut {
		c = c.Add(color.CrossedOut)
	}
	return c
}

func (f *flags) ToColor() *maybeColor {
	return f.ApplyTo(newColor(color.Reset))
}

func fgAttribute(name string) color.Attribute {
	switch strings.ToLower(name) {
	case "black":
		return color.FgBlack
	case "red":
		return color.FgRed
	case "green":
		return color.FgGreen
	case "yellow":
		return color.FgYellow
	case "blue":
		return color.FgBlue
	case "magenta":
		return color.FgMagenta
	case "cyan":
		return color.FgCyan
	case "white":
		return color.FgWhite
	case "hiblack", "hi-black":
		return color.FgHiBlack
	case "hired", "hi-red":
		return color.FgHiRed
	case "higreen", "hi-green":
		return color.FgHiGreen
	case "hiyellow", "hi-yellow":
		return color.FgHiYellow
	case "hiblue", "hi-blue":
		return color.FgHiBlue
	case "himagenta", "hi-magenta":
		return color.FgHiMagenta
	case "hicyan", "hi-cyan":
		return color.FgHiCyan
	case "hiwhite", "hi-white":
		return color.FgHiWhite
	default:
		return color.Reset
	}
}

func bgAttribute(name string) color.Attribute {
	switch strings.ToLower(name) {
	case "black":
		return color.BgBlack
	case "red":
		return color.BgRed
	case "green":
		return color.BgGreen
	case "yellow":
		return color.BgYellow
	case "blue":
		return color.BgBlue
	case "magenta":
		return color.BgMagenta
	case "cyan":
		return color.BgCyan
	case "white":
		return color.BgWhite
	case "hiblack", "hi-black":
		return color.BgHiBlack
	case "hired", "hi-red":
		return color.BgHiRed
	case "higreen", "hi-green":
		return color.BgHiGreen
	case "hiyellow", "hi-yellow":
		return color.BgHiYellow
	case "hiblue", "hi-blue":
		return color.BgHiBlue
	case "himagenta", "hi-magenta":
		return color.BgHiMagenta
	case "hicyan", "hi-cyan":
		return color.BgHiCyan
	case "hiwhite", "hi-white":
		return color.BgHiWhite
	default:
		return color.Reset
	}
}
