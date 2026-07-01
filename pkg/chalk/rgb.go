package chalk

import "fmt"

func (c *Chalk) RGB(r, g, b int) *Chalk {
	r = clamp(r, 0, 255)
	g = clamp(g, 0, 255)
	b = clamp(b, 0, 255)
	return c.Add(Code(fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)))
}

func (c *Chalk) BgRGB(r, g, b int) *Chalk {
	r = clamp(r, 0, 255)
	g = clamp(g, 0, 255)
	b = clamp(b, 0, 255)
	return c.Add(Code(fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)))
}
