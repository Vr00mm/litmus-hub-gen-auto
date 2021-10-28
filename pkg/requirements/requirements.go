package requirements

import (
	"fmt"
	"hub-gen-auto/pkg/resources"
	"strings"
)

//import "hub-gen-auto/pkg/types"

func checkIfRoot(composant interface{}) bool {
	return false
}

func FindUniqueLabels(composants *resources.Resources) bool {
	ready := true
	for composant := range composants.Objects {
		for labelName, labelValue := range composants.Objects[composant].GetLabels() {
			skip := false
			label := fmt.Sprintf("%s=%s", labelName, labelValue)
			for _, composantO := range composants.Objects {
				if string(composants.Objects[composant].GetName()) != string(composantO.GetName()) {
					for labelNameO, labelValueO := range composantO.GetLabels() {
						labelO := fmt.Sprintf("%s=%s", labelNameO, labelValueO)
						if label == labelO {
							skip = true
							ready = false
						}
					}
				}
			}
			if !skip {
				composants.Objects[composant].AddUniqueLabel(label)
			}
		}
	}
	return ready
}

func CheckHaveLabels(namespace *resources.Resources, requiredLabels []string) (bool, string) {
	compliant := true
	var err []string
	//	fmt.Printf("TRACE: %v", namespace)
	for _, composant := range namespace.Objects {
		for _, requiredLabel := range requiredLabels {
			if _, found := composant.GetLabel(requiredLabel); !found {
				compliant = false
				msg := "Label " + requiredLabel + " not found on "
				err = append(err, msg)
			}
		}
	}
	return compliant, strings.Join(err, "\n")
}

func CheckRequirements(requirements []string, composant resources.Object) bool {
	compliant := true
	for _, requirement := range requirements {
		switch requirement {
		case "isRoot":
			isRoot := checkIfRoot(composant)
			if !isRoot {
				compliant = false
			}
		}
	}
	return compliant
}
