package main

import (
	"os"

	"github.com/kahlaoui-2020/go-chalk/pkg/chalk"
)

func main() {
	c := chalk.New()
	c.
		Add(chalk.Red).
		Add(chalk.BgBlue).
		Add(chalk.Bold).
		Add(chalk.Underline)
	c.Println("Hello, World!")
	c.Disable().Println("Disabled styles")
	c.Enable().Clear().RGB(255, 0, 0).Println("This is red text")
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	c.SetWriter(f).Println("This will be written to the file")

}
