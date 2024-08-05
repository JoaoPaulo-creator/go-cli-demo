package utils

import (
	"encoding/json"
)

func BodyParser(body interface{}) []byte {

	parsedBody, err := json.Marshal(&body)
	if err != nil {
		panic("Erro ao parsear body")
	}

	return parsedBody
}

func ResponseBodyParser(response []byte, body interface{}) []byte {

	err := json.Unmarshal(response, &body)
	if err != nil {
		panic("deu ruimno Unmarshal")
	}

	return response
}
