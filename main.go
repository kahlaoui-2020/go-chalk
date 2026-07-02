package main

import (
	"github.com/kahlaoui-2020/go-chalk/pkg/chalk"
)

func main() {
	c := chalk.New()
	c.
		Add(chalk.Red).
		Add(chalk.BgBlue).
		Add(chalk.Bold).
		Add(chalk.Underline)
	c.Debug().Println("Hello, World!")
	c.Disable().Debug().Print("Disabled styles\n")
	c.Enable().Clear().Debug().RGB(255, 0, 0).Printf("This is red text %s\n", "with RGB color")

}
