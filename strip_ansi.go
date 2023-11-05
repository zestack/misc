package misc

import "regexp"

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var ansiRegex = regexp.MustCompile(ansi)

// StripAnsi 清除 ANSI 转义代码
// https://en.wikipedia.org/wiki/ANSI_escape_code
// https://github.com/acarl005/stripansi/blob/master/stripansi.go
func StripAnsi(str string) string {
	return ansiRegex.ReplaceAllString(str, "")
}
