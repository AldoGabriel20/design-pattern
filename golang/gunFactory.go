package golang

import "fmt"

func getGun(gunType string) (iGun, error) {
	if gunType == "bazooka" {
		return newBazooka(), nil
	}

	if gunType == "sniper" {
		return newSniper(), nil
	}

	return nil, fmt.Errorf("Gun type %s not recognized", gunType)
}
