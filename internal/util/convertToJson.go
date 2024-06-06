package util

import (
	"encoding/json"
	"log"
)

// StringToJson deserializa un string JSON a una estructura genérica de Go
func StringToJson[T any](jsonStr string) (*T, error) {
	var result T
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return nil, err
	}
	return &result, nil
}

func BodyToJson[T any](body []byte) (*T, error) {
	var result T
	err := json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("Error unmarshalling Body to JSON: %v", err)
		return nil, err
	}
	return &result, nil
}
