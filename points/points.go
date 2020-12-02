package points

import (
	"encoding/json"
	"net/http"
	"time"
)

// USER POINT STATE
type Transaction struct {
	Payer  string
	Points int
	Date   string
}

var Transactions []Transaction

// RESPONSE STRUCT
type Response struct {
	Message string
	Data    interface{}
}

func Add(res http.ResponseWriter, req *http.Request) {
	t := Transaction{}
	r := Response{}

	// decode incoming into transaction struct
	err := json.NewDecoder(req.Body).Decode(&t)
	if err != nil {
		r.Message = "Could not decode request body"
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(r)
		return
	}

	// validate payer field
	if t.Payer == "" {
		r.Message = "Request body must contain a `payer` field"
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(r)
		return
	}

	//  validate date field
	const layoutIso = "01/02/2006 03PM"
	_, err = time.Parse(layoutIso, t.Date)
	if err != nil {
		r.Message = "Improperly formatted `date` field"
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(r)
		return
	}

	// add transaction to Transactions slice
	Transactions = append(Transactions, t)

	// response
	r.Message = "Points added"
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(r)
}

func Deduct(res http.ResponseWriter, req *http.Request) {
	type points struct {
		points int
	}
	// decode incoming into points struct
	// make sure transactions are sorted by date
	// calculate total points per payer (so that negative additions in the future do not cause payer to go negative)
	// var deductions []Transaction
	// find first transaction where points != 0
	// deduct until points are exhausted, and create corresponding entry in deductions
	// response
}

func Get(res http.ResponseWriter, req *http.Request) {
	type totalPoints struct {
		payer  string
		points int
	}
	// loop through all transactions
	// if totalPoints member with payer name doesnt exist, create it
	// if it does exist, modify it to reflect new points value
	// response
}
