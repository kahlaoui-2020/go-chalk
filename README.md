# Chalk Package Documentation

[![Go Report Card](https://goreportcard.com/badge/github.com/chalk/chalk-go)](https://goreportcard.com/report/github.com/chalk/chalk-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## About

This package provides a simple way to apply styles (colors, styles) and apply them to text or formatted strings. It is designed to work with various types of writers such as `os.Stdout` for console output.

The purpose of this project is to demonstrate how easy it can be to create a simple library that you can use in your projects without having to worry about the intricacies of file descriptors, error handling, etc., allowing you to focus on writing great code instead of low-level stuff. We aim to make it as minimal and straightforward as possible while still being flexible enough for advanced users.

## Installation

```sh
go get github.com/kahlaoui-2020/go-chalk
```

## Getting Started

## Getting Started

### Basic Usage Example

```go
package main

import "github.com/kahlaoui-2020/go-chalk/chalk"

func main() {
    // Create a new instance of the chalk package.
    c := chalk.New()

    // Apply styles to text or formatted strings.
    c.BgRed().
        Bold().Println("Hello, World!")

    // You can also use named methods for better readability.
    c.Blue().
        Dim().Italic().
        Underline().
        Println("I am a fancy message.")
}
```

### Available Methods

Here's a list of all available methods:

```go
func (c *Chalk) Black()
func (c *Chalk) Red()
func (c *Chalk) Green()
func (c *Chalk) Yellow()
func (c *Chalk) Blue()
func (c *Chalk) Magenta()
func (c *Chalk) Cyan()
func (c *Chalk) White()

// Background color methods
func (c *Chalk) BgBlack()
func (c *Chalk) BgRed()
func (c *Chalk) BgGreen()
func (c *Chalk) BgYellow()
func (c *Chalk) BgBlue()
func (c *Chalk) BgMagenta()
func (c *Chalk) BgCyan()
func (c *Chalk) BgWhite()

// Style methods
func (c *Chalk) Bold()
func (c *Chalk) Dim()
func (c *Chalk) Italic()
func (c *Chalk) Underline()
func (c *Chalk) Blink()
func (c *Chalk) Inverse()
func (c *Chalk) Hidden()
func (c *Chalk) Strikethrough()

func (c *Chalk) Print(args ...any)
func (c *Chalk) Println(args ...any)
func (c *Chalk) Printf(format string, args ...any)
func (c *Chalk) Sprintf(format string, args ...any) string
func (c *Chalk) String() string

// Clone the current instance.
func (c *Chalk) clone() *Chalk {
    // ...
}

func main() {
    c := chalk.New()
    c.Add(chalk.Red).Add(chalk.BgBlue).Add(chalk.Bold)
}
```

### Examples

- **Basic Usage**: Apply styles to text or formatted strings directly.

```go
// Basic usage example.
c := New()
c.Black().Println("Hello, World!")

// Applying multiple styles at once.
c.Green().
    Underline().
    Println("This is a fancy message.")
```

### Contributing

Feel free to contribute by submitting issues or pull requests. Make sure to follow the existing coding standards and guidelines.

## License

MIT