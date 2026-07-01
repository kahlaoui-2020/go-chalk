package chalk

import (
	"fmt"
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

func (c *Chalk) Black() *Chalk {
	newChalk := c.clone()
	newChalk.colors = append(newChalk.colors, Black)
	return newChalk
}
func (c *Chalk) Red() *Chalk {
	newChalk := c.clone()
	newChalk.colors = append(newChalk.colors, Red)
	return newChalk
}

func (c *Chalk) Green() *Chalk {
	newChalk := c.clone()
	newChalk.colors = append(newChalk.colors, Green)
	return newChalk
}
func (c *Chalk) Yellow() *Chalk {
	newChalk := c.clone()
	newChalk.colors = append(newChalk.colors, Yellow)
	return newChalk
}
func (c *Chalk) Blue() *Chalk {
	newChalk := c.clone()
	newChalk.colors = append(newChalk.colors, Blue)
	return newChalk
}
func (c *Chalk) Magenta() *Chalk {
	newChalk := c.clone()
	newChalk.colors = append(newChalk.colors, Magenta)
	return newChalk
}
func (c *Chalk) Cyan() *Chalk {
	newChalk := c.clone()
	newChalk.colors = append(newChalk.colors, Cyan)
	return newChalk
}
func (c *Chalk) White() *Chalk {
	newChalk := c.clone()
	newChalk.colors = append(newChalk.colors, White)
	return newChalk
}

// Background color methods
func (c *Chalk) BgBlack() *Chalk {
	newChalk := c.clone()
	newChalk.colors = append(newChalk.colors, BgBlack)
	return newChalk
}
func (c *Chalk) BgRed() *Chalk {
	newChalk := c.clone()
	newChalk.colors = append(newChalk.colors, BgRed)
	return newChalk
}
func (c *Chalk) BgGreen() *Chalk {
	newChalk := c.clone()
	newChalk.colors = append(newChalk.colors, BgGreen)
	return newChalk
}
func (c *Chalk) BgYellow() *Chalk {
	newChalk := c.clone()
	newChalk.colors = append(newChalk.colors, BgYellow)
	return newChalk
}
func (c *Chalk) BgBlue() *Chalk {
	newChalk := c.clone()
	newChalk.colors = append(newChalk.colors, BgBlue)
	return newChalk
}
func (c *Chalk) BgMagenta() *Chalk {
	newChalk := c.clone()
	newChalk.colors = append(newChalk.colors, BgMagenta)
	return newChalk
}
func (c *Chalk) BgCyan() *Chalk {
	newChalk := c.clone()
	newChalk.colors = append(newChalk.colors, BgCyan)
	return newChalk
}
func (c *Chalk) BgWhite() *Chalk {
	newChalk := c.clone()
	newChalk.colors = append(newChalk.colors, BgWhite)
	return newChalk
}

// Style methods
func (c *Chalk) Bold() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, Bold)
	return newChalk
}
func (c *Chalk) Dim() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, Dim)
	return newChalk
}
func (c *Chalk) Italic() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, Italic)
	return newChalk
}
func (c *Chalk) Underline() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, Underline)

	return newChalk
}
func (c *Chalk) Blink() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, Blink)
	return newChalk
}
func (c *Chalk) Inverse() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, Inverse)
	return newChalk
}
func (c *Chalk) Hidden() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, Hidden)
	return newChalk
}
func (c *Chalk) Strikethrough() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, Strikethrough)
	return newChalk
}

func (c *Chalk) Print(args ...any) {
	text := c.applyStyles(fmt.Sprint(args...))
	if _, err := fmt.Fprint(c.writer, text); err != nil {
		panic(err)
	}
}
func (c *Chalk) Println(args ...any) {
	text := c.applyStyles(fmt.Sprint(args...))
	if _, err := fmt.Fprintln(c.writer, text); err != nil {
		panic(err)
	}
}
func (c *Chalk) Printf(format string, args ...any) {
	text := c.applyStyles(fmt.Sprintf(format, args...))
	if _, err := fmt.Fprint(c.writer, text); err != nil {
		panic(err)
	}
}
func (c *Chalk) Sprintf(format string, args ...any) string {
	return c.applyStyles(fmt.Sprintf(format, args...))
}
func (c *Chalk) String() string {
	return c.applyStyles("")
}
