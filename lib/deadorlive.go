package lib

import (
	"fmt"
	"net/http"
)

func DeadOrLive(uri string) bool {
	req, _ := http.NewRequest("GET", uri, nil)
	client := new(http.Client)

	resq, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return false
	}

	defer resq.Body.Close()

	if resq.StatusCode != 200 {
		fmt.Println("Error")
		return false
	} else {
		return true
	}

	// body, _ := io.ReadAll(resq.Body)
	// var result any

	// json.Unmarshal(body, &result)

	// fmt.Printf("%-v", result)
}
