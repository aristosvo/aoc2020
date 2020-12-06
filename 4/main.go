package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"log"
	"regexp"
	"strconv"
)

var inputFile = flag.String("inputFile", "input.txt", "Relative file path to use as input.")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		return
	}
	contents := string(bytes)
	split := strings.Split(contents, "\n\n")
	valid1 := 0
	valid2 := 0
	for _, passportInput := range split {
		fields := strings.Fields(passportInput)
		byr1 := false
		iyr1 := false
		eyr1 := false
		hgt1 := false
		hcl1 := false
		ecl1 := false
		pid1 := false

		byr2 := false
		iyr2 := false
		eyr2 := false
		hgt2 := false
		hcl2 := false
		ecl2 := false
		pid2 := false

		for _, field := range fields {
			fieldname := strings.Split(field, ":")
			switch {
			case fieldname[0] == "byr":
				byr1 = true
				// four digits; at least 1920 and at most 2002.
				byrRaw := fieldname[1]
				year, err := strconv.Atoi(byrRaw)
				if err != nil {
					log.Fatal(err)
				}
				if len(byrRaw) == 4 && year >= 1920 && year <= 2002 {
					byr2 = true
				}
			case fieldname[0] == "iyr":
				iyr1 = true
				// four digits; at least 2010 and at most 2020.
				iyrRaw := fieldname[1]
				year, err := strconv.Atoi(iyrRaw)
				if err != nil {
					log.Fatal(err)
				}
				if len(iyrRaw) == 4 && year >= 2010 && year <= 2020 {
					iyr2 = true
				}
			case fieldname[0] == "eyr":
				eyr1 = true
				// four digits; at least 2020 and at most 2030.
				eyrRaw := fieldname[1]
				year, err := strconv.Atoi(eyrRaw)
				if err != nil {
					log.Fatal(err)
				}
				if len(eyrRaw) == 4 && year >= 2020 && year <= 2030 {
					eyr2 = true
				}
			case fieldname[0] == "hgt":
				hgt1 = true
				// a number followed by either cm or in:
				// - If cm, the number must be at least 150 and at most 193.
				// - If in, the number must be at least 59 and at most 76.
				hgtRaw := fieldname[1]
				if len(hgtRaw) < 4 || len(hgtRaw) >5 {
					continue
				}
				length, err := strconv.Atoi(hgtRaw[:len(hgtRaw)-2])
				//fmt.Println(length)
				if err != nil {
					log.Fatal(err)
				}
				if len(hgtRaw) == 5 && strings.HasSuffix(hgtRaw, "cm") && length >= 150 && length <= 193 {
					hgt2 = true
				}
				if len(hgtRaw) == 4 && strings.HasSuffix(hgtRaw, "in") && length >= 59 && length <= 76 {
					hgt2 = true
				}
			case fieldname[0] == "hcl":
				hcl1 = true
				// a # followed by exactly six characters 0-9 or a-f.
				hclRaw := fieldname[1]
				re, err := regexp.Compile(`#[a-f0-9]{6}`)
				if err != nil {
					log.Fatal(err)
				}
				if re.MatchString(hclRaw) && len(hclRaw) == 7 {
					hcl2 = true
				}
			case fieldname[0] == "ecl":
				ecl1 = true
				// exactly one of: amb blu brn gry grn hzl oth.
				ecl := fieldname[1]
				if ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth" {
					ecl2 = true
				}
			case fieldname[0] == "pid":
				pid1 = true
				// a nine-digit number, including leading zeroes.
				pid := fieldname[1]
				re, err := regexp.Compile(`[0-9]{9}`)
				if err != nil {
					log.Fatal(err)
				}
				if re.MatchString(pid) && len(pid) == 9 {
					pid2 = true
				}
			}
		}
		if byr1 && iyr1 && eyr1 && hgt1 && hcl1 && ecl1 && pid1 {
			valid1++
		}

		//fmt.Printf("%t - %t - %t - %t - %t - %t - %t\n",byr2, iyr2, eyr2, hgt2, hcl2, ecl2, pid2)
		if byr2 && iyr2 && eyr2 && hgt2 && hcl2 && ecl2 && pid2 {
			valid2++
		}
	}
	fmt.Println(valid1)
	fmt.Println(valid2)
}
