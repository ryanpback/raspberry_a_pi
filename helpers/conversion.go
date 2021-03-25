package helpers

import "fmt"

// ConvertInterfaceToString converts an interface (likely from a map[string]interface{} type) to a string
func ConvertInterfaceToString(i interface{}) string {
	return fmt.Sprintf("%v", i)
}
