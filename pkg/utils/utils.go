package utils

import (
	"fmt"
	"os"
	"strings"
)

func WriteArrayToFile(sampledata []string, outfile string) {
	file, err := os.OpenFile(outfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		dbg := strings.Join(sampledata, "\n")
		fmt.Fprintf(os.Stderr, "%v", dbg)
		fmt.Fprintf(os.Stderr, "failed creating file: %v\n", err)
	}
	for _, data := range sampledata {
		fmt.Fprintln(file, data)
	}
	//      datawriter.Flush()
	file.Close()
}
