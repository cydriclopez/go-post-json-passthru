package treedata

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// ***********************************************************
// This JsonData should agree with Angular-side interface
type JsonData struct {
	Data string `json:"data"`
} // *********************************************************
// Angular-side interface in src/app/services/nodeservice.ts
// export interface JsonData {
//     data:   string;
// }
// ***********************************************************

// Small-case non-exported local identifier
type tData struct {
	Jdata  JsonData
	Config string
}

// Constructor pattern using factory method
func TData() *tData {
	t := new(tData)
	return t
}

func (t *tData) PostJsonData(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		var unmarshalTypeError *json.UnmarshalTypeError

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&t.Jdata); err != nil {

			if errors.As(err, &unmarshalTypeError) {
				jsonResponse(w, http.StatusBadRequest, "Error wrong data type: "+unmarshalTypeError.Field)

			} else {
				jsonResponse(w, http.StatusBadRequest, "Error: "+err.Error())
			}

			return
		}

		// Save json in db
		t.saveJsonData()

		jsonResponse(w, http.StatusOK, "Success")
		return
	}

	log.Print("http.NotFound")
	http.NotFound(w, r)
}

func (t *tData) saveJsonData() {
	// For now we will just print the data from the client
	log.Println("jsonData:", t.Jdata)
}

func jsonResponse(w http.ResponseWriter, statusCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// For production, use generic "Bad request or data error".
	// Detailed error message is not advised in production.
	w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, errorMsg) + "\n"))
}
