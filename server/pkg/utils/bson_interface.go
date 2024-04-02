package utils

import (
	"captioner.com.ng/pkg/constants"
)

func BsonInterfacer(v []constants.Identifier) map[string]interface{} {
	result := make(map[string]interface{})

	for _, val := range v {
		result[val.Key] = val.Value
	}

	return result
}
