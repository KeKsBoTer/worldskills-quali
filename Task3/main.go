package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Transfer struct {
	From, To string
	Amount   float64
}

func forEachLine(filepath string, fn func(Transfer)) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\r\n") {
		trans := strings.Split(line, " ")
		if len(trans) != 3 {
			// line has no credentials
			continue
		}
		amount, err := strconv.ParseFloat(trans[2], 64)
		if err != nil {
			continue
		}
		fn(Transfer{
			From:   trans[0],
			To:     trans[1],
			Amount: amount,
		})
	}
}

func main() {
	transferCount := 0
	users := map[string]float64{}
	var aReceived, eReceived, transfered float64
	guarneriMoney := 0
	forEachLine("Test_03_Data.txt", func(t Transfer) {
		transferCount++
		transfered += t.Amount
		users[t.From] -= t.Amount
		users[t.To] += t.Amount
		if strings.HasPrefix(t.To, "a") {
			aReceived += t.Amount
		} else if strings.HasPrefix(t.To, "e") {
			eReceived += t.Amount
		}
		if t.To == "guarneri" {
			guarneriMoney++
		}
	})
	fmt.Println("Users:", len(users))
	checkNegative := []string{"bisexual", "rombert", "aeroscope", "vannesavanness", "yearbook", "espousal", "apodaca", "friseur", "fagin"}
	for _, user := range checkNegative {
		if users[user] < 0 {
			fmt.Println(user, "has negative balance of:", users[user])
		}
	}
	fmt.Println("aeroscope:", users["aeroscope"], ", bisexual(x2):", users["bisexual"]*2)
	fmt.Println("Transfers:", transferCount)
	fmt.Println("intellectual:", users["intellectual"], ", bonzer:", users["bonzer"])
	fmt.Println("'a' received:", aReceived, ", 'e' received:", eReceived)
	fmt.Println("transfered money:", transfered)
	fmt.Println(guarneriMoney, "users received guarneri money")
}
