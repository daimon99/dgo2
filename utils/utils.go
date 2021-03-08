package utils

import "os"

func IsExist(path string) (result bool, ret_err error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			result = false
		} else {
			// other error
			result = false
			ret_err = err
		}
	} else {
		//exist
		result = true
	}
	return result, ret_err
}
