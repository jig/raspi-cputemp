package cputemp

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const termalzone = "/sys/class/thermal/thermal_zone0/temp"

// Get returns the temperature of the Raspberry processor
func Get() (float64, error) {
	raw, err := ioutil.ReadFile(termalzone)
	if err != nil {
		return 0, fmt.Errorf("error reading %q: %v", termalzone, err)
	}

	cpuTemp, err := strconv.Atoi(strings.TrimSpace(string(raw)))
	if err != nil {
		return 0, errors.New("invalid format")
	}
	return float64(cpuTemp) / 1000, nil
}
