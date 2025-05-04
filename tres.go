package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/potato", resultPrint)
	fmt.Print("http://localhost:8000/potato")
	http.ListenAndServe(":8000", nil)
}

func resultPrint(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprintln(w, "potato!")

	x := req.URL.Query().Get("x")
	y := req.URL.Query().Get("y")

	i, err1 := strconv.Atoi(x)
	j, err2 := strconv.Atoi(y)

	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	switch req.Method {
	case "PUT":
		sum := i + j
		fmt.Fprintf(w, "x: %d, y:%d | Sum: %d\n", i, j, sum)
	case "GET":
		diff := i - j
		fmt.Fprintf(w, "x: %d, y:%d | Difference: %d\n", i, j, diff)
	case "POST":
		prod := i * j
		fmt.Fprintf(w, "x: %d, y:%d | Product: %d\n", i, j, prod)
	case "DELETE":
		if j == 0 {
			http.Error(w, "Division by zero", http.StatusBadRequest)
			return
		} else {
			quot := i / j
			fmt.Fprintf(w, "x: %d, y:%d | Quotient: %d\n", i, j, quot)
		}
	default:
		{
			fmt.Fprintf(w, "nihil!", http.StatusBadRequest)
		}
	}
}
