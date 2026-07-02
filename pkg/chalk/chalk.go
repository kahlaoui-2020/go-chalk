package chalk

import (
	"io"
	"os"
	"strings"
)

type Chalk struct {
	codes   []Code
	writer  io.Writer
	enabled bool
	debug   bool
}

func New() *Chalk {
	return &Chalk{
		codes:   []Code{},
		writer:  os.Stdout,
		enabled: isTerminal(),
		debug:   false,
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
func (c *Chalk) SetDebug(debug bool) *Chalk {
	c.debug = debug
	return c
}
func (c *Chalk) Debug() *Chalk {
	cl := c.clone()
	cl.debug = true
	return cl
}

func (c *Chalk) Add(code Code) *Chalk {
	c.codes = append(c.codes, code)
	return c
}
func (c *Chalk) String() string {
	return c.applyStyles("")
}
func (c *Chalk) Clear() *Chalk {
	c.codes = []Code{}
	return c
}
func (c *Chalk) applyStyles(text string) string {
	if !c.enabled || (len(c.codes) == 0) {
		return text
	}
	var sb strings.Builder
	for _, code := range c.codes {
		sb.WriteString(string(code))
	}

	sb.WriteString(text)
	sb.WriteString(string(Reset))
	return sb.String()
}
func (c *Chalk) clone() *Chalk {
	newChalk := &Chalk{
		codes:   append([]Code{}, c.codes...),
		writer:  c.writer,
		enabled: c.enabled,
	}
	return newChalk
}
