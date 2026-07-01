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

// Styles
const (
	Bold          Color = "\033[1m"
	Dim           Color = "\033[2m"
	Italic        Color = "\033[3m"
	Underline     Color = "\033[4m"
	Blink         Color = "\033[5m"
	Inverse       Color = "\033[7m"
	Hidden        Color = "\033[8m"
	Strikethrough Color = "\033[9m"
)

const (
	Reset Color = "\033[0m"
)
