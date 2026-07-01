package main

import "github.com/kahlaoui-2020/go-chalk/pkg/chalk"

func main() {
	c := chalk.New()
	c.
		AddColor(chalk.Red).
		AddColor(chalk.BgBlue).
		AddStyle(chalk.Bold).
		AddStyle(chalk.Underline)
	c.Println("Hello, World!")

	c1 := chalk.New()
	c1.Red().Println("This is red text")
}
