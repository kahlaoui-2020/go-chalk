package chalk

import (
	"io"
	"os"
	"strings"
)

type Chalk struct {
	colors  []Color
	styles  []Style
	writer  io.Writer
	enabled bool
}

func New() *Chalk {
	return &Chalk{
		colors:  []Color{},
		styles:  []Style{},
		writer:  os.Stdout,
		enabled: isTerminal(),
	}
}

func isTerminal() bool {
	fileInfo, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}
func (c *Chalk) SetWriter(w io.Writer) *Chalk {
	c.writer = w
	return c
}
func (c *Chalk) Enable() *Chalk {
	c.enabled = true
	return c
}
func (c *Chalk) Disable() *Chalk {
	c.enabled = false
	return c
}
func (c *Chalk) AddColor(color Color) *Chalk {
	c.colors = append(c.colors, color)
	return c
}
func (c *Chalk) AddStyle(style Style) *Chalk {
	c.styles = append(c.styles, style)
	return c
}
func (c *Chalk) applyStyles(text string) string {
	if !c.enabled || (len(c.styles) == 0 && len(c.colors) == 0) {
		return text
	}
	var styledText strings.Builder
	for _, style := range c.styles {
		styledText.WriteString(string(style))
	}
	for _, color := range c.colors {
		styledText.WriteString(string(color))
	}
	styledText.WriteString(text)
	styledText.WriteString(string(Reset))
	return styledText.String()
}
func (c *Chalk) clone() *Chalk {
	newChalk := &Chalk{
		colors:  append([]Color{}, c.colors...),
		styles:  append([]Style{}, c.styles...),
		writer:  c.writer,
		enabled: c.enabled,
	}
	return newChalk
}
func (c *Chalk) String() string {
	return c.applyStyles("")
}
