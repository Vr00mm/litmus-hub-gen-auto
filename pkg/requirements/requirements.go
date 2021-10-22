package requirements

func checkIfRoot() (bool) {
  return false
}

func checkHaveLabels(composant interface{}) (bool) {
    labels := []string{"composant"}
    compliant := true
    for label := range labels {
        if _, found := composant[label]; found != true {
            compliant = false
        }
    }
    return compliant
}


func CheckRequirements(requirements []string) (bool) {
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
