package utils

import (
	"fmt"
	"hub-gen-auto/pkg/types"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func MakeHttpRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func RemoveDuplicateMap(strMap []struct {
	Name  string "yaml:\"name\""
	Email string "yaml:\"email\""
}) []struct {
	Name  string "yaml:\"name\""
	Email string "yaml:\"email\""
} {
	allKeys := make(map[string]bool)
	var list []struct {
		Name  string "yaml:\"name\""
		Email string "yaml:\"email\""
	}
	for _, item := range strMap {
		if _, value := allKeys[item.Name]; !value {
			allKeys[item.Name] = true
			list = append(list, struct {
				Name  string "yaml:\"name\""
				Email string "yaml:\"email\""
			}{item.Name, item.Email})
		}
	}
	return list
}

func GetExperimentsManifests(experimentsList []string) map[string]types.ChaosChart {
	experiments := make(map[string]types.ChaosChart, len(experimentsList))
	for _, experimentName := range experimentsList {
		var chaosChart types.ChaosChart
		var chaosChartVersion types.ChaosChartVersion
		var chaosEngine types.ChaosEngine
		var chaosExperiment types.ChaosExperiment

		chaosChartVersionManifest := MakeHttpRequest("https://raw.githubusercontent.com/litmuschaos/chaos-charts/master/charts/generic/" + experimentName + "/" + experimentName + ".chartserviceversion.yaml")
		chaosEngineManifest := MakeHttpRequest("https://raw.githubusercontent.com/litmuschaos/chaos-charts/master/charts/generic/" + experimentName + "/engine.yaml")
		chaosExperimentManifest := MakeHttpRequest("https://raw.githubusercontent.com/litmuschaos/chaos-charts/master/charts/generic/" + experimentName + "/experiment.yaml")
		chaosExperimentPsp := MakeHttpRequest("https://raw.githubusercontent.com/litmuschaos/chaos-charts/master/charts/generic/" + experimentName + "/rbac-psp.yaml")
		chaosExperimentRbac := MakeHttpRequest("https://raw.githubusercontent.com/litmuschaos/chaos-charts/master/charts/generic/" + experimentName + "/rbac.yaml")

		chaosIcon := []byte(MakeHttpRequest("https://raw.githubusercontent.com/litmuschaos/chaos-charts/master/charts/generic/icons/" + experimentName + ".png"))

		if err := yaml.Unmarshal(chaosChartVersionManifest, &chaosChartVersion); err != nil {
			panic(err)
		}
		if err := yaml.Unmarshal(chaosEngineManifest, &chaosEngine); err != nil {
			panic(err)
		}
		if err := yaml.Unmarshal(chaosExperimentManifest, &chaosExperiment); err != nil {
			panic(err)
		}

		chaosChart.ChaosEngine = chaosEngine
		chaosChart.ChaosExperiment = chaosExperiment
		chaosChart.ChartVersion = chaosChartVersion
		chaosChart.Icon = chaosIcon
		chaosChart.PSP = chaosExperimentPsp
		chaosChart.RBAC = chaosExperimentRbac
		experiments[experimentName] = chaosChart
	}

	return experiments
}

func WriteToFile(filename string, data string) {
	dir := filepath.Dir(filename)
	os.MkdirAll(dir, os.ModePerm)

	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
	}
}

func WriteToFileAsYaml(filename string, data interface{}) {
	dir := filepath.Dir(filename)
	os.MkdirAll(dir, os.ModePerm)

	file, err := yaml.Marshal(data)
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
	}
	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
	}
}

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
