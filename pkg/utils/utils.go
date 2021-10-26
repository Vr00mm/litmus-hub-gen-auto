package main

import "fmt"

func getLabel(o map[string]string, label string) string {
	if _, found := o[label]; found {
		return o[label]
	}
	return ""
}

func writeArrayToFile(sampledata []string, outfile string) {
	file, err := os.OpenFile(outfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, strings.Join(sampledata, "\n"))
		fmt.Fprintf(os.Stderr, "failed creating file: %v\n", err)
	}
	for _, data := range sampledata {
		fmt.Fprintln(file, data)
	}
	//      datawriter.Flush()
	file.Close()
}
