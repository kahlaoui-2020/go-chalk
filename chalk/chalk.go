package chalk

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Chalk struct {
	styles  []Color
	writer  io.Writer
	enabled bool
}

func New() *Chalk {
	return &Chalk{
		styles:  []Color{},
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
func (c *Chalk) Add(style Color) *Chalk {
	c.styles = append(c.styles, style)
	return c
}
func (c *Chalk) applyStyles(text string) string {
	if !c.enabled || len(c.styles) == 0 {
		return text
	}
	var styledText strings.Builder
	for _, style := range c.styles {
		styledText.WriteString(string(style))
	}
	styledText.WriteString(text)
	styledText.WriteString(string(Reset))
	return styledText.String()
}
func (c *Chalk) clone() *Chalk {
	newChalk := &Chalk{
		styles:  append([]Color{}, c.styles...),
		writer:  c.writer,
		enabled: c.enabled,
	}
	return newChalk
}

func (c *Chalk) Black() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, Black)
	return newChalk
}
func (c *Chalk) Red() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, Red)
	return newChalk
}

func (c *Chalk) Green() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, Green)
	return newChalk
}
func (c *Chalk) Yellow() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, Yellow)
	return newChalk
}
func (c *Chalk) Blue() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, Blue)
	return newChalk
}
func (c *Chalk) Magenta() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, Magenta)
	return newChalk
}
func (c *Chalk) Cyan() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, Cyan)
	return newChalk
}
func (c *Chalk) White() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, White)
	return newChalk
}

// Background color methods
func (c *Chalk) BgBlack() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, BgBlack)
	return newChalk
}
func (c *Chalk) BgRed() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, BgRed)
	return newChalk
}
func (c *Chalk) BgGreen() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, BgGreen)
	return newChalk
}
func (c *Chalk) BgYellow() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, BgYellow)
	return newChalk
}
func (c *Chalk) BgBlue() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, BgBlue)
	return newChalk
}
func (c *Chalk) BgMagenta() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, BgMagenta)
	return newChalk
}
func (c *Chalk) BgCyan() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, BgCyan)
	return newChalk
}
func (c *Chalk) BgWhite() *Chalk {
	newChalk := c.clone()
	newChalk.styles = append(newChalk.styles, BgWhite)
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
	fmt.Fprint(c.writer, text)
}
func (c *Chalk) Println(args ...any) {
	text := c.applyStyles(fmt.Sprint(args...))
	fmt.Fprintln(c.writer, text)
}
func (c *Chalk) Printf(format string, args ...any) {
	text := c.applyStyles(fmt.Sprintf(format, args...))
	fmt.Fprint(c.writer, text)
}
func (c *Chalk) Sprintf(format string, args ...any) string {
	return c.applyStyles(fmt.Sprintf(format, args...))
}
func (c *Chalk) String() string {
	return c.applyStyles("")
}

var DefaultChalk = New()

func Print(args ...any) {
	DefaultChalk.Print(args...)
}
func Println(args ...any) {
	DefaultChalk.Println(args...)
}
func Printf(format string, args ...any) {
	DefaultChalk.Printf(format, args...)
}
func Sprintf(format string, args ...any) string {
	return DefaultChalk.Sprintf(format, args...)
}
