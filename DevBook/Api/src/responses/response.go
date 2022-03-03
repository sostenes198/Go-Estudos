package responses

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

type responseErro struct {
	Message string `json:"message"`
}

type responseErros []struct {
	Message string `json:"message"`
}

// JSON retorna um resposta em JSOn para a requisição
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data == nil {
		return
	}

	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		writeSliceResponse(w, data)
		return
	default:
		writeResponse(w, data)
		return
	}
}

func writeSliceResponse(w http.ResponseWriter, data interface{}) {
	rv := reflect.ValueOf(data)
	if rv.Len() > 0 {
		writeResponse(w, data)
	}
}

func writeResponse(w http.ResponseWriter, data interface{}) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatalln(err)
	}
}

// Erro retorna um responseErro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, err error) {
	Erros(w, statusCode, []error{err})
}

// Erros retorna um array de responseErros em formato JSON
func Erros(w http.ResponseWriter, statusCode int, errs []error) {
	var responseErros responseErros
	for _, err := range errs {
		responseErros = append(responseErros, responseErro{Message: err.Error()})
	}

	JSON(w, statusCode, responseErros)
}
