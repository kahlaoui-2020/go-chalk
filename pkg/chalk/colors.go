package chalk

// Foreground codes
const (
	Black   Code = "\033[30m"
	Red     Code = "\033[31m"
	Green   Code = "\033[32m"
	Yellow  Code = "\033[33m"
	Blue    Code = "\033[34m"
	Magenta Code = "\033[35m"
	Cyan    Code = "\033[36m"
	White   Code = "\033[37m"

	BrightBlack   Code = "\033[90m"
	BrightRed     Code = "\033[91m"
	BrightGreen   Code = "\033[92m"
	BrightYellow  Code = "\033[93m"
	BrightBlue    Code = "\033[94m"
	BrightMagenta Code = "\033[95m"
	BrightCyan    Code = "\033[96m"
	BrightWhite   Code = "\033[97m"
)

// Background codes
const (
	BgBlack   Code = "\033[40m"
	BgRed     Code = "\033[41m"
	BgGreen   Code = "\033[42m"
	BgYellow  Code = "\033[43m"
	BgBlue    Code = "\033[44m"
	BgMagenta Code = "\033[45m"
	BgCyan    Code = "\033[46m"
	BgWhite   Code = "\033[47m"

	BgBrightBlack   Code = "\033[100m"
	BgBrightRed     Code = "\033[101m"
	BgBrightGreen   Code = "\033[102m"
	BgBrightYellow  Code = "\033[103m"
	BgBrightBlue    Code = "\033[104m"
	BgBrightMagenta Code = "\033[105m"
	BgBrightCyan    Code = "\033[106m"
	BgBrightWhite   Code = "\033[107m"
)

func (c *Chalk) Black() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, Black)
	return chalk
}
func (c *Chalk) Red() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, Red)
	return chalk
}

func (c *Chalk) Green() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, Green)
	return chalk
}
func (c *Chalk) Yellow() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, Yellow)
	return chalk
}
func (c *Chalk) Blue() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, Blue)
	return chalk
}
func (c *Chalk) Magenta() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, Magenta)
	return chalk
}
func (c *Chalk) Cyan() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, Cyan)
	return chalk
}
func (c *Chalk) White() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, White)
	return chalk
}

// Background color methods
func (c *Chalk) BgBlack() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, BgBlack)
	return chalk
}
func (c *Chalk) BgRed() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, BgRed)
	return chalk
}
func (c *Chalk) BgGreen() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, BgGreen)
	return chalk
}
func (c *Chalk) BgYellow() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, BgYellow)
	return chalk
}
func (c *Chalk) BgBlue() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, BgBlue)
	return chalk
}
func (c *Chalk) BgMagenta() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, BgMagenta)
	return chalk
}
func (c *Chalk) BgCyan() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, BgCyan)
	return chalk
}
func (c *Chalk) BgWhite() *Chalk {
	chalk := c.clone()
	chalk.codes = append(chalk.codes, BgWhite)
	return chalk
}
