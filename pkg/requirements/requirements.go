package requirements

import "fmt"
import "strings"
import "hub-gen-auto/pkg/resources"

//import "hub-gen-auto/pkg/types"

func checkIfRoot(composant interface{}) bool {
	return false
}

func FindUniqueLabels(composants *resources.Resources) (*resources.Resources, bool) {
	unique := make(map[string][]string, len(composants.Objects))
	result := composants
	var dirty []resources.Object
	ready := true
	for _, composant := range composants.Objects {
		for labelName, labelValue := range composant.GetLabels() {
			skip := false
			label := fmt.Sprintf("%s=%s", labelName, labelValue)
			for _, composantO := range composants.Objects {
				if composant.GetName() != composantO.GetName() {
					for labelNameO, labelValueO := range composantO.GetLabels() {
						labelO := fmt.Sprintf("%s=%s", labelNameO, labelValueO)
						if label == labelO {
							skip = true
						}
					}
				}
			}
			if !skip {
				unique[composant.GetName()] = append(unique[composant.GetName()], label)
				composant.AddUniqueLabel(label)
				dirty = append(dirty, composant)
			}
		}
		if len(unique[composant.GetName()]) < 1 {
			ready = false
		}
	}
	result.Objects = dirty
	return result, ready
}

func CheckHaveLabels(namespace *resources.Resources, requiredLabels []string) (bool, string) {
	compliant := true
	var err []string
	//	fmt.Printf("TRACE: %v", namespace)
	for _, composant := range namespace.Objects {
		for _, requiredLabel := range requiredLabels {
			if _, found := composant.GetLabel(requiredLabel); found != true {
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
			result := checkIfRoot(composant)
			if result != true {
				compliant = false
			}
		}
	}
	return compliant
}
