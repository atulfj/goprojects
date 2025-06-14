package main

import (
	"encoding/json"
	"net/http"
)

type CalcRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumRequest struct {
	Array []int `json:"array"`
}

type CalcResponse struct {
	Res int `json:"result"`
}

func add(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	c := CalcRequest{}
	err := decoder.Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cr := CalcResponse{}
	cr.Res = c.A + c.B

	w.WriteHeader(http.StatusOK) // happens by default also
	encoder := json.NewEncoder(w)
	err = encoder.Encode(&cr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func subtract(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	c := CalcRequest{}
	if err := decoder.Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cr := CalcResponse{
		Res: c.A - c.B,
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(&cr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func multiply(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	c := CalcRequest{}
	if err := decoder.Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cr := CalcResponse{
		Res: c.A * c.B,
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(&cr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func divide(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	c := CalcRequest{}
	if err := decoder.Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if c.B == 0 {
		http.Error(w, "Cannot divide by 0", http.StatusBadRequest)
		return
	}

	cr := CalcResponse{
		Res: c.A / c.B,
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(&cr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func sum(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	s := SumRequest{}
	if err := decoder.Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sum := 0
	for _, n := range s.Array {
		sum += n
	}

	cr := CalcResponse{
		Res: sum,
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(&cr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/add", add)
	http.HandleFunc("/subtract", subtract)
	http.HandleFunc("/multiply", multiply)
	http.HandleFunc("/divide", divide)
	http.HandleFunc("/sum", sum)

	http.ListenAndServe(":3000", nil)
}
