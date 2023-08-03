package ascii

func CheckColor(userValue string) string {
	Colorss := map[string]string{"red": "\033[31m", "green": "\033[32m", "yellow": "\033[33m", "blue": "\033[34m", "purple": "\033[35m", "cyan": "\033[36m", "white": "\033[37m"}
	//traverse through the map
	for a, value := range Colorss {
		//check if present value is equals to userValue
		if a == userValue {
			//if same return true
			return value
		}
	}
	//if value not found return No
	return "NO"
}
