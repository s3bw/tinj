package tinj

import (
	"bufio"
	"io"
	"os"
)

// NewLine rune
const NewLine = '\n'

func ReadStdin(format string) {
	var line []rune

	lineFormatter := DeconstructFormat(format)
	stdin := bufio.NewReader(os.Stdin)

	for {
		nextChar, _, err := stdin.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		line = append(line, nextChar)

		if nextChar == NewLine {
			lineFormatter.Format(line)
			line = nil
		}
	}
}
