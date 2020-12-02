package points

import (
	"encoding/json"
	"net/http"
	"time"
)

// USER POINT STATE
type Transaction struct {
	Payer     string
	Points    int
	Date      string
	_DateTime time.Time
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
		r.Message = "Request body must contain a `Payer` field"
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(r)
		return
	}

	// validate points field
	if t.Points == 0 {
		r.Message = "Request body must contain a `Points` field that is not 0"
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(r)
		return
	}

	//  validate date field
	// validate payer field
	if t.Date == "" {
		r.Message = "Request body must contain a `Date` field"
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(r)
		return
	}
	const layoutIso = "01/02/2006 03PM"
	dateTime, err := time.Parse(layoutIso, t.Date)
	if err != nil {
		r.Message = "Improperly formatted `date` field"
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(r)
		return
	}

	// append dateTime in new format
	t._DateTime = dateTime

	// add transaction to Transactions slice
	Transactions = append(Transactions, t)

	// sort transactions by _DateTime

	// response
	r.Message = "Points added"
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(r)
}

func Deduct(res http.ResponseWriter, req *http.Request) {
	type Deduct struct {
		points int
	}

	d := Deduct{}
	r := Response{}

	// decode incoming into points struct
	err := json.NewDecoder(req.Body).Decode(&d)
	if err != nil {
		r.Message = "Could not decode request body"
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(r)
		return
	}

	// calculate total points per payer (so that negative additions in the future do not cause player to go negative)
	balance := getTotalPointsPerPayer()

	// var deductions []Transaction
	// find first transaction where points != 0
	// deduct until points are exhausted, and create corresponding entry in deductions
	// response
}

func Balance(res http.ResponseWriter, req *http.Request) {
	r := Response{}

	// response
	r.Message = "Total points balance"
	r.Data = getTotalPointsPerPayer()
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(r)
}

func getTotalPointsPerPayer() (TotalPointsList map[string]int) {
	for _, v := range Transactions {
		_, ok := TotalPointsList[v.Payer]
		if ok {
			TotalPointsList[v.Payer] += v.Points
		} else {
			TotalPointsList[v.Payer] = v.Points
		}
	}
	return
}
