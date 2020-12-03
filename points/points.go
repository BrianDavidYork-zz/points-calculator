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
	DateTime_ time.Time `json:"-"`
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

	// append dateTime in new format, set deducted to False
	t.DateTime_ = dateTime

	// add transaction to Transactions slice
	Transactions = append(Transactions, t)

	// response
	r.Message = "Points added"
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(r)
}

func Deduct(res http.ResponseWriter, req *http.Request) {
	type Deduct struct {
		Points int
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

	var deductions []Transaction
	// loop through all point additions
	for i, v := range Transactions {
		// if payer has points to deduct and v.Points is positive (negative numbers are implicitly deducted)
		if balance[v.Payer] > 0 && v.Points > 0 && d.Points > 0 {
			// get deduct amount
			var deductAmount int
			if d.Points >= balance[v.Payer] {
				deductAmount = balance[v.Payer]
			} else if d.Points < balance[v.Payer] {
				deductAmount = d.Points
			}
			// subtract deductAmount from balance[v.Payer]    (payer point total)
			balance[v.Payer] -= deductAmount
			// subtract deductAmount from v.Points            (transaction record)
			Transactions[i].Points -= deductAmount
			// subtract deductAmount from d.Points            (amount to be deducted from call)
			d.Points -= deductAmount
			// create new transaction and append to deductions
			t := Transaction{}
			t.Payer = Transactions[i].Payer
			t.Points = -(deductAmount)
			t.Date = "now"
			deductions = append(deductions, t)
		}
	}
	// response
	if len(deductions) == 0 {
		r.Message = "No available points to deduct"
	} else {
		r.Message = "Points deducted"
	}
	r.Data = deductions
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(r)
}

func Balance(res http.ResponseWriter, req *http.Request) {
	r := Response{}

	// get points balance
	balance := getTotalPointsPerPayer()

	// response
	r.Message = "Total points balance"
	r.Data = balance
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(r)
}

func getTotalPointsPerPayer() (TotalPointsList map[string]int) {
	TotalPointsList = make(map[string]int)
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
