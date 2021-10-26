package requirements

import "fmt"

func checkIfRoot() bool {
	return false
}

func FindUniqueLabelsBb(composants map[string]map[string]string) map[string][]string {
	unique := make(map[string][]string, len(composants))
	for composant, o := range composants {
		for labelName, labelValue := range o {
			skip := false
			label := fmt.Sprintf("%s=%s", labelName, labelValue)
			for composantO, uo := range composants {
				if composant != composantO {
					for labelNameO, labelValueO := range uo {
						labelO := fmt.Sprintf("%s=%s", labelNameO, labelValueO)
						if label == labelO {
							skip = true
						}
					}
				}
			}
			if !skip {
				unique[composant] = append(unique[composant], label)
				//fmt.Printf("Valeur de unique: %v\n", unique)
			}
		}
	}
	return unique
}

func FindUniqueLabels(composants map[string]map[string]string) map[string][]string {
	unique := make(map[string][]string, len(composants))
	for composant, o := range composants {
		skip := false
		for ko, vo := range o {
			labelValue := fmt.Sprintf("%s=%s", ko, vo)
			for _, u := range unique {
				for _, vu := range u {
					labelu := string(vu)
					if labelValue == labelu {
						skip = true
						break
					}
				}
			}
			if !skip {
				unique[composant] = append(unique[composant], labelValue)
				//fmt.Printf("Valeur de unique: %v\n", unique)
			}
		}
	}
	return unique
}

func CheckHaveLabels(labels map[string]string) bool {
	requiredLabels := []string{"composant"}
	compliant := true
	for _, requiredLabel := range requiredLabels {
		if _, found := labels[requiredLabel]; found != true {
			compliant = false
		}
	}
	return compliant
}

func CheckRequirements(requirements []string) bool {
	compliant := true
	for _, requirement := range requirements {
		switch requirement {
		case "isRoot":
			result := checkIfRoot()
			if result != true {
				compliant = false
			}
		}
	}
	return compliant
}
