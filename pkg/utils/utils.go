package utils

import (
	"encoding/json"
	"log"
	"github.com/vaddepally-shiva-shankar/go-crudapi-sql/pkg/models"
)

// data -> JSON body from request
// obj -> pointer to the struct where the data needs to be unmarshalled
func UnmarshalFunc(data []byte) (*models.Book, error) {
	obj := models.Book{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		log.Println("Error unmarshalling data:", err)
		return nil, err
	}
	return &obj, nil
}

func MarshalFunc(obj interface{}) ([]byte, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		log.Println("Error marshalling object:", err)
		return nil, err
	}
	return data, nil
}