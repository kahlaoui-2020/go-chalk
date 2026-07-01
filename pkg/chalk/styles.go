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
