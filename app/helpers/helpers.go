package helpers

import (
	"bytes"
	"encoding/json"
	"golang_gin/app/databases/model"
	"golang_gin/app/libraries"
	"log"
)

func GetToken(user *model.Users) string {
	token, err := libraries.NewPasetoToken().GenerateToken(user.ID)

	if err != nil {
		log.Fatal(err)
	}

	return token
}

func PrepareBody(body map[string]any) (*bytes.Buffer, error) {
	jsonBytes, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(jsonBytes), nil
}
