package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}

	input := string(content)

	total := countValidPassprts(input)

	fmt.Printf("Valid passports: %d\n", total)
}

func countValidPassprts(input string) int {
	passports := strings.Split(input, "\n\n")

	validCount := 0
	for _, rawPassport := range passports {
		p := newPassport(rawPassport)

		if p.validate() {
			validCount = validCount + 1
		}
	}

	return validCount
}

type passport struct {
	fields map[string]string
}

func validateDate(v string, min, max int) bool {
	if len(v) != 4 {
		return false
	}
	date, err := strconv.Atoi(v)
	if err != nil {
		return false
	}
	if date < min || date > max {
		return false
	}

	return true
}

func validateHieght(v string) bool {
	re := regexp.MustCompile(`([0-9]*)(in|cm)`)
	matches := re.FindStringSubmatch(v)
	if len(matches) != 3 {
		return false
	}

	mt := matches[2]
	mm, _ := strconv.Atoi(matches[1])
	switch mt {
	case "in":
		if mm < 59 || mm > 76 {
			return false
		}
	case "cm":
		if mm < 150 || mm > 193 {
			return false
		}
	default:
		//unkown
		return false
	}

	return true
}

func (p *passport) validate() bool {
	if len(p.fields) < 7 {
		return false
	}

	found := false
	for key, value := range p.fields {
		switch key {
		case "byr":
			isValid := validateDate(value, 1920, 2002)
			if !isValid {
				return false
			}
		case "iyr":
			isValid := validateDate(value, 2010, 2020)
			if !isValid {
				return false
			}
		case "eyr":
			isValid := validateDate(value, 2020, 2030)
			if !isValid {
				return false
			}
		case "hgt":
			isValid := validateHieght(value)
			if !isValid {
				return false
			}
		case "hcl":
			matched, _ := regexp.MatchString(`^#[0-9|a-f]{6,6}$`, value)
			if !matched {
				return false
			}
		case "ecl":
			matched, _ := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, value)
			if !matched {
				return false
			}
		case "pid":
			matched, _ := regexp.MatchString(`^[0-9]{9,9}$`, value)
			if !matched {
				return false
			}
		case "cid":
			found = true
			continue
		default:
			// unkown key
			return false
		}
	}

	if len(p.fields) == 8 {
		fmt.Println(p)
		fmt.Println("Valid! All eight fields are valid.")
		return true
	}

	if len(p.fields) == 7 {
		if !found {
			fmt.Println(p)
			fmt.Println("Valid! Fields valid, only missing cid.")
			return true
		}
	}

	// unknown
	return false
}

func newPassport(rawPassport string) passport {
	fields := strings.Fields(rawPassport)

	m := make(map[string]string)
	for _, f := range fields {
		record := strings.Split(f, ":")
		m[record[0]] = record[1]
	}

	return passport{
		fields: m,
	}
}
