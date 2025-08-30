package utils

import "runtime"

func GetGOVersion() string {
	return runtime.Version()[2:]
}
