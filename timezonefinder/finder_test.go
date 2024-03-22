package timezonefinder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckData(t *testing.T) {
	setupZone()
	for _, country := range timeZoneCountryCode {
		_, ok := countryToContinentMapping[country]
		assert.Equal(t, ok, true)
	}
}

func TestGetCountryAndContinentByTimeZone(t *testing.T) {
	var tests = []string{
		"GMT+08:00,Asia/Shanghai",
		"GMT+8,Asia/Shanghai",
		"GMT+08:00,Asia/Macau",
		"GMT+09:00,Asia/Seoul",
		"GMT-07:30,Asia/Ho_Chi_Minh",
		"GMT+09:00,Asia/Tokyo",
		"GMT+02:00,Europe/Paris",
		"GMT+12,Pacific/Apia",
		"GMT+02:00,Europe/Madrid",
		"GMT-07:00,America/Denver",
		"GMT+03:30,Asia/Tehran",
		"GMT+02:00,Africa/Cairo",
		"GMT+09:00,Asia/Yakutsk",
		"GMT-07:00,America/Los_Angeles",
		"GMT+09:30,Australia/Adelaide",
		"GMT+9,Asia/Shanghai",
		"GMT-02:30,Australia/Darwin",
		"GMT-03:30,America/St_Johns",
	}
	setupZone()
	for _, v := range tests {
		continent, country, err := GetCountryAndContinentByTimeZone(v)
		t.Log(continent, country, err)
		assert.Nil(t, err, nil)
		assert.NotEmpty(t, continent)
		assert.NotEmpty(t, country)
	}
}
