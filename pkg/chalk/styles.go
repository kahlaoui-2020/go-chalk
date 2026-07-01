package chalk

// Styles
const (
	Bold          Code = "\033[1m"
	Dim           Code = "\033[2m"
	Italic        Code = "\033[3m"
	Underline     Code = "\033[4m"
	Blink         Code = "\033[5m"
	Inverse       Code = "\033[7m"
	Hidden        Code = "\033[8m"
	Strikethrough Code = "\033[9m"
)

func (c *Chalk) Bold() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, Bold)
	return chalk
}
func (c *Chalk) Dim() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, Dim)
	return chalk
}
func (c *Chalk) Italic() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, Italic)
	return chalk
}
func (c *Chalk) Underline() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, Underline)

	return chalk
}
func (c *Chalk) Blink() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, Blink)
	return chalk
}
func (c *Chalk) Inverse() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, Inverse)
	return chalk
}
func (c *Chalk) Hidden() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, Hidden)
	return chalk
}
func (c *Chalk) Strikethrough() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, Strikethrough)
	return chalk
}
