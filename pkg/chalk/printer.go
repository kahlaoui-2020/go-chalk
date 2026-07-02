package chalk

import (
	"fmt"
	"time"
)

func (c *Chalk) Print(args ...any) {
	text := c.applyStyles(fmt.Sprint(args...))
	if c.debug {
		text = fmt.Sprintf("[DEBUG][%s]: %s", time.Now().Format("15:04:05"), text)
	}
	if _, err := fmt.Fprint(c.writer, text); err != nil {
		panic(err)
	}
}
func (c *Chalk) Println(args ...any) {
	text := c.applyStyles(fmt.Sprintln(args...))
	if c.debug {
		text = fmt.Sprintf("[DEBUG][%s]: %s", time.Now().Format("15:04:05"), text)
	}
	if _, err := fmt.Fprint(c.writer, text); err != nil {
		panic(err)
	}
}
func (c *Chalk) Printf(format string, args ...any) {
	text := c.applyStyles(fmt.Sprintf(format, args...))
	if c.debug {
		text = fmt.Sprintf("[DEBUG][%s]: %s", time.Now().Format("15:04:05"), text)
	}
	if _, err := fmt.Fprint(c.writer, text); err != nil {
		panic(err)
	}
}
func (c *Chalk) Sprintf(format string, args ...any) string {
	text := c.applyStyles(fmt.Sprintf(format, args...))
	if c.debug {
		text = fmt.Sprintf("[DEBUG][%s]: %s", time.Now().Format("15:04:05"), text)
	}
	return text
}
