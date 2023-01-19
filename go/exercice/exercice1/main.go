package main

import (
	"strconv"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type OperationService interface {
	Somme(int64, int64) (int64, error)
}

type operationService struct{}

func (operationService) Somme(nombre1 int64, nombre2 int64) (int64, error) {
	return nombre1 + nombre2, nil
}

type sommeRequest struct {
	Nombre1 string `json:"nombre1"`
	Nombre2 string `json:"nombre2"`
}

type sommeResponse struct {
	Resultat int64 `json:"resultat"`
	Err      string `json:"err,omitempty"`
}

func makeSommeEndpoint(svc OperationService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
        req := request.(sommeRequest)
        nombre1, err := strconv.ParseInt(req.Nombre1, 10, 64)
        if err != nil {
            return sommeResponse{-1, err.Error()}, nil
        }
        nombre2, err := strconv.ParseInt(req.Nombre2, 10, 64)
        if err != nil {
            return sommeResponse{-1, err.Error()}, nil
        }
 
        resultat, err := svc.Somme(nombre1, nombre2)
        if err != nil {
            return sommeResponse{-1, err.Error()}, nil
        }
        return sommeResponse{resultat, ""}, nil
    }
}

func main() {
	svc := operationService{}

	sommeHandler := httptransport.NewServer(
		makeSommeEndpoint(svc),
		decodeSommeRequest,
		encodeResponse,
	)

	http.Handle("/somme", sommeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeSommeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request sommeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}