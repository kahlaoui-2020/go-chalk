package chalk

import "fmt"

func (c *Chalk) Print(args ...any) {
	text := c.applyStyles(fmt.Sprint(args...))
	if _, err := fmt.Fprint(c.writer, text); err != nil {
		panic(err)
	}
}
func (c *Chalk) Println(args ...any) {
	text := c.applyStyles(fmt.Sprintln(args...))
	if _, err := fmt.Fprint(c.writer, text); err != nil {
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
