package chalk

type Color string

// Foreground colors
const (
	Black   Color = "\033[30m"
	Red     Color = "\033[31m"
	Green   Color = "\033[32m"
	Yellow  Color = "\033[33m"
	Blue    Color = "\033[34m"
	Magenta Color = "\033[35m"
	Cyan    Color = "\033[36m"
	White   Color = "\033[37m"

	BrightBlack   Color = "\033[90m"
	BrightRed     Color = "\033[91m"
	BrightGreen   Color = "\033[92m"
	BrightYellow  Color = "\033[93m"
	BrightBlue    Color = "\033[94m"
	BrightMagenta Color = "\033[95m"
	BrightCyan    Color = "\033[96m"
	BrightWhite   Color = "\033[97m"
)

// Background colors
const (
	BgBlack   Color = "\033[40m"
	BgRed     Color = "\033[41m"
	BgGreen   Color = "\033[42m"
	BgYellow  Color = "\033[43m"
	BgBlue    Color = "\033[44m"
	BgMagenta Color = "\033[45m"
	BgCyan    Color = "\033[46m"
	BgWhite   Color = "\033[47m"

	BgBrightBlack   Color = "\033[100m"
	BgBrightRed     Color = "\033[101m"
	BgBrightGreen   Color = "\033[102m"
	BgBrightYellow  Color = "\033[103m"
	BgBrightBlue    Color = "\033[104m"
	BgBrightMagenta Color = "\033[105m"
	BgBrightCyan    Color = "\033[106m"
	BgBrightWhite   Color = "\033[107m"
)

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
