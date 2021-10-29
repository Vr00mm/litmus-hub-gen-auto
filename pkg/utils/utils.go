package utils

import (
	"fmt"
	"hub-gen-auto/pkg/types"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func makeHttpRequest(url string) []byte {
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

func GetExperimentsManifests(experimentsList []string) map[string]types.ChaosChart {
	experiments := make(map[string]types.ChaosChart, len(experimentsList))
	for _, experimentName := range experimentsList {
		var chaosChart types.ChaosChart

		var chaosChartVersion types.ChaosChartVersion
		var chaosEngine types.ChaosEngine
		var chaosExperiment types.ChaosExperiment

		chaosChartVersionManifest := makeHttpRequest("https://raw.githubusercontent.com/litmuschaos/chaos-charts/master/charts/generic/" + experimentName + "/" + experimentName + ".chartserviceversion.yaml")
		chaosEngineManifest := makeHttpRequest("https://raw.githubusercontent.com/litmuschaos/chaos-charts/master/charts/generic/" + experimentName + "/engine.yaml")
		chaosExperimentManifest := makeHttpRequest("https://raw.githubusercontent.com/litmuschaos/chaos-charts/master/charts/generic/" + experimentName + "/experiment.yaml")
		chaosIcon := []byte(makeHttpRequest("https://raw.githubusercontent.com/litmuschaos/chaos-charts/master/charts/generic/icons/" + experimentName + ".png"))

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
		experiments[experimentName] = chaosChart
	}

	return experiments
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
