package apis

import (
	"encoding/json"
	"net/http"

	"github.com/Shubhacker/interview-2.0.git~/service"
	"github.com/gorilla/mux"
)

// On API level, we can add all business logic. Such as validations, checks.

type errorStruct struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    string `json:"errorCode"`
}

var i service.CurrencyInterface

func RegisterCache() {
	var d service.Cache
	d.CountryCache = map[string]service.Data{}
	i = &d
}

func GetCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cName := vars["name"]
	ctx := r.Context()

	res, err := i.GetCountry(ctx, cName)
	if err != nil {
		var e errorStruct
		e.ErrorMessage = err.Error()
		// later can add different http code
		e.ErrorCode = "500"
		json.NewEncoder(w).Encode(e)
	}

	json.NewEncoder(w).Encode(res)

}
