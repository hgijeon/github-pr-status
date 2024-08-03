package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/go-gh/v2/pkg/tableprinter"
	"github.com/cli/go-gh/v2/pkg/text"
	"golang.org/x/term"
)

type String string

func createLink(s, url string) string {
	trimmedText := strings.TrimSpace(s)
	ss := String(trimmedText).addLink(url).addUnderline()

	return strings.Replace(s, trimmedText, string(ss), 1)
}

func (s String) addLink(url string) String {
	return String(fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", url, string(s)))
}

func (s String) addUnderline() String {
	underlineStart := "\033[4m"
	underlineEnd := "\033[24m"

	return String(fmt.Sprintf("%s%s%s", underlineStart, string(s), underlineEnd))
}

func newTablePrinterWithHeaders(w *os.File, headers []string, cs *iostreams.ColorScheme) tableprinter.TablePrinter {
	isTTY := term.IsTerminal(int(w.Fd()))
	width, _, err := term.GetSize(int(w.Fd()))
	if err != nil {
		width = 80
	}

	tp := tableprinter.New(w, isTTY, width)
	if isTTY && len(headers) > 0 {
		upperCasedHeaders := make([]string, len(headers))
		for i, header := range headers {
			upperCasedHeaders[i] = strings.ToUpper(header)
		}

		var paddingFunc func(int, string) string
		if cs.Enabled() {
			paddingFunc = text.PadRight
		}

		tp.AddHeader(upperCasedHeaders, tableprinter.WithPadding(paddingFunc), tableprinter.WithColor(cs.LightGrayUnderline))
	}

	return tp
}
