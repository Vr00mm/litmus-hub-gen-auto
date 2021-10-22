package containerKill

import (
	"hub-gen-auto/pkg/requirements"
       )

var experimentRequirements = []string {"isRoot"}

func CheckRequirements() (bool) {
	return requirements.CheckRequirements(experimentRequirements)
}
