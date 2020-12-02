package points

import "net/http"

// USER POINT STATE
type transaction struct {
	payer  string
	points int
	date   string
}

var Transactions []transaction

func Add(res http.ResponseWriter, req *http.Request) {
	// decode incoming into transaction struct
	// validate all 3 fields
	// add transaction to Transactions slice
	// response
}

func Deduct(res http.ResponseWriter, req *http.Request) {
	type points struct {
		points int
	}
	// decode incoming into points struct
	// make sure transactions are sorted by date
	// calculate total points per payer (so that negative additions in the future do not cause payer to go negative)
	var deductions []transaction
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
