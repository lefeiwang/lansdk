package tool

func IsExist(em interface{}, slice []interface{}) bool {
	for _, v := range slice {
		if v == em {
			return true
		}
	}
	return false
}
