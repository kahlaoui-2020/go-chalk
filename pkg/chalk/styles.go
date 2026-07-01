package chalk

type Style string

// Styles
const (
	Bold          Style = "\033[1m"
	Dim           Style = "\033[2m"
	Italic        Style = "\033[3m"
	Underline     Style = "\033[4m"
	Blink         Style = "\033[5m"
	Inverse       Style = "\033[7m"
	Hidden        Style = "\033[8m"
	Strikethrough Style = "\033[9m"
)

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
