package main

import (
	"fmt"
	"strings"
)

func GetCountryAndContinentByTimeZone(timeZone string) (continent, country string, err error) {
	country, err = GetCountryCodeByTimeZone(timeZone)
	if err != nil {
		return "", "", err
	}
	continent, err = GetContinentByCountry(country)
	if err != nil {
		return "", "", err
	}
	return continent, country, nil

}

func getRegion(timeZone string) string {
	arr := strings.Split(timeZone, ",")
	if len(arr) == 2 {
		return arr[1]
	}
	return ""
}

func GetCountryCodeByTimeZone(timeZone string) (string, error) {
	region := getRegion(timeZone)
	if region == "" {
		return "", fmt.Errorf("region not found for timezone: %s", timeZone)
	}
	val, ok := timeZoneCountryCode[region]
	if !ok {
		return "", fmt.Errorf("country code not found for timezone: %s", timeZone)
	}
	return val, nil
}

func GetContinentByCountry(country string) (string, error) {
	val, ok := countryToContinentMapping[country]
	if !ok {
		return "", fmt.Errorf("continent not found for country: %s", country)
	}
	return val, nil
}
