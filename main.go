package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"simpleBackEndInGolang/repository"
	"strings"
)

var (
	mockRepository = repository.MockRepository{}
)

func readAllTransaction(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(responseWriter).Encode(mockRepository.GetAllTransaction())
}

func getBalanace(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(responseWriter).Encode(mockRepository.GetBalance())
}

func transactionWriterHandler(responseWriter http.ResponseWriter, request *http.Request) {
	var transaction repository.Transaction
	err := json.NewDecoder(request.Body).Decode(&transaction)
	if err != nil {
		fmt.Print("error with json decoding : ", err)
	}
	err = mockRepository.PostTransaction(&transaction)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		_, _ = responseWriter.Write([]byte("error inserting transaction: " + err.Error()))
		return
	}
	_, _ = responseWriter.Write([]byte("transaction suscessfull"))
}

func transactionReaderHandler(responseWriter http.ResponseWriter, request *http.Request) {
	id := strings.TrimPrefix(request.URL.Path, "/transaction/")
	fmt.Println("getting transaction with Id : " + id)
	transaction, err := mockRepository.GetTransactionById(id)
	if err != nil {

	} else {
		responseWriter.Header().Set("Content-Type", "application/json")
		if transaction == nil {
			_, _ = responseWriter.Write([]byte("No transaction found by given Id "))
		}

		_ = json.NewEncoder(responseWriter).Encode(transaction)
	}

}

func handleRequests() {
	http.HandleFunc("/", getBalanace)
	http.HandleFunc("/alltransactions", readAllTransaction)
	http.HandleFunc("/transaction", transactionWriterHandler)
	http.HandleFunc("/transaction/", transactionReaderHandler)
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func main() {
	handleRequests()
}
