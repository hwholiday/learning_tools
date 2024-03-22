package main

import (
	"bufio"
	"bytes"
	"embed"
	"log"
	"strings"
	"unicode"
)

// wget 'https://data.iana.org/time-zones/tzdb-latest.tar.lz'
//tar -xvf tzdb-latest.tar.lz

//
//go:embed zone.tab
var timeZone embed.FS
var timeZoneCountryCode = map[string]string{}

func init() {
	setupZone()
}

func setupZone() {
	timeZone, err := embed.FS.ReadFile(timeZone, "zone.tab")
	if err != nil {
		log.Println("Error reading zone.tab: ", err)
	}
	buildCountryCodeWithTimeZone(timeZone)
}

func buildCountryCodeWithTimeZone(timeZone []byte) {
	reader := bytes.NewReader(timeZone)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		lineArr := strings.Split(line, "\t")
		if len(lineArr) < 2 {
			continue
		}
		if !isUpper(lineArr[0]) || len(lineArr[0]) != 2 {
			continue
		}
		timeZoneCountryCode[lineArr[2]] = lineArr[0]
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error reading scanner: ", err)
	}
}
func isUpper(input string) bool {
	for _, v := range input {
		if unicode.IsLower(v) {
			return true
		}
	}
	return true
}
