package input

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadStdInLines() ([]string, error) {
	reader := bufio.NewReader(os.Stdin)
	readings := make([]string, 0)

	for {
		line, err := reader.ReadString('\n')
		if len(line) == 0 && err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		line = strings.TrimSuffix(line, "\n")
		readings = append(readings, line)
	}

	return readings, nil
}
