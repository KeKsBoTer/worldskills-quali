package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"unicode"
)

type User struct {
	Name     string
	Password Password
}

type Password string

// regex that checks rule 1-3
var secureCheck = regexp.MustCompile("^[A-Z][a-zA-Z0-9]{5,}[A-Z]$")

func (p Password) IsSecure() bool {
	pwString := string(p)
	if !secureCheck.MatchString(pwString) {
		return false
	}
	letters, digits := countLettersAndDigits(pwString)
	if letters <= digits {
		return false
	}
	if hasConsecutiveDigits(pwString) {
		return false
	}
	return true
}

func forEachLine(filepath string, fn func(User)) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\r\n") {
		cred := strings.Split(line, " ")
		if len(cred) != 2 {
			// line has no credentials
			continue
		}
		fn(User{cred[0], Password(cred[1])})
	}
}

func main() {
	count := 0
	secure := map[string]bool{}
	pwCount := map[Password]int{
		"Kfgbg8tpXR":    0,
		"UewMik5KqM":    0,
		"NHkpZf4gQ":     0,
		"Y7ho2gu0anL":   0,
		"Fx8hf8egS":     0,
		"KyWedsxso7KW":  0,
		"JX3HMdfYWDfoD": 0,
	}
	var longestSecure, longestInsecure User
	forEachLine("Aufg_02_Data.txt", func(user User) {
		count++
		if _, ok := pwCount[user.Password]; ok {
			pwCount[user.Password]++
		}
		if user.Password.IsSecure() {

			secure[user.Name] = true
			if len(user.Password) > len(longestSecure.Password) {
				longestSecure = user
			}
		} else {
			if len(user.Password) > len(longestInsecure.Password) {
				longestInsecure = user
			}
		}
	})
	fmt.Println("Entries:", count)
	fmt.Println("Secure Passwords:", len(secure))
	for _, id := range []string{"decurion", "fiche", "piss", "ahearn", "areopagus", "victorvictoria", "anishaaniso", "merman"} {
		if secure[id] {
			fmt.Println(id, "has secure password")
		}
	}
	fmt.Println("Longest secure password:", longestSecure)
	fmt.Println("Longest insecure password:", longestInsecure)
	fmt.Println("Longest secure password length:", len(longestSecure.Password))
	fmt.Println("Longest insecure password length:", len(longestInsecure.Password))
	for pw, count := range pwCount {
		if count > 1 {
			fmt.Println(pw, "is used", count, "times")
		}
	}
}

func countLettersAndDigits(pw string) (letters int, nums int) {
	for _, l := range pw {
		if unicode.IsDigit(l) {
			nums++
		} else if unicode.IsLower(l) {
			letters++
		}
	}
	return
}

func hasConsecutiveDigits(pw string) bool {
	last := false
	for _, l := range pw {
		if unicode.IsDigit(l) {
			if last {
				return true
			}
			last = true
		} else {
			last = false
		}
	}
	return false
}
