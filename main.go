package main

import "github.com/kahlaoui-2020/go-chalk/pkg/chalk"

func main() {
	c := chalk.New()
	c.
		Add(chalk.Red).
		Add(chalk.BgBlue).
		Add(chalk.Bold).
		Add(chalk.Underline)
	c.Println("Hello, World!")

	c1 := chalk.New()
	c1.RGB(255, 0, 0).Println("This is red text")
}
