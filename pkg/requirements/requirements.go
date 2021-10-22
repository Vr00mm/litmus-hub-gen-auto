package requirements

func checkIfRoot() (bool) {
  return false
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
