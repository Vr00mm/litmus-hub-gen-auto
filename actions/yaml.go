package actions

import (
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"log"
	"strings"
)

func ReadYamlFile(filename string) []runtime.Object {
	dataByte, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	dataString := string(dataByte)
	yamlfiles := strings.Split(dataString, "---")
	var resources []runtime.Object

	for _, file := range yamlfiles {
		file = strings.Trim(file, " ")
		file = strings.Trim(file, "\n")
		if file == "\n" || file == "" {
			continue
		}

		decode := scheme.Codecs.UniversalDeserializer().Decode
		obj, _, err := decode([]byte(file), nil, nil)

		if err != nil {
			log.Fatal(err)
			continue
		}
		resources = append(resources, obj)
	}
	return resources
}
