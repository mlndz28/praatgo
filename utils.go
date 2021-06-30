package praatgo

import (
	"bytes"
	"errors"
	"strconv"
)

func parseString(data []byte) string {
	val := bytes.ReplaceAll(data, []byte(`"`), []byte{})
	val = bytes.TrimSpace(val)
	return string(val)
}
func parseNumber(data []byte) (float64, error) {
	return strconv.ParseFloat(string(bytes.TrimSpace(data)), 64)
}

func parseIndex(data []byte) (int, error) {
	return strconv.Atoi(string(bytes.TrimSpace(data)))
}

func parseBool(data []byte) (bool, error) {
	var err error
	switch string(data) {
	case "<exists>":
		return true, err
	case "<absent>":
		return false, err
	default:
		err = errors.New("<exists> or <absent> boolean type expected")
	}
	return false, err
}

func parseInterval(data [][]byte) (Interval, error){
	xmin, err := parseNumber(data[0])
	if err != nil { return Interval{}, err }
	xmax, err := parseNumber(data[1])
	if err != nil { return Interval{}, err }
	return Interval{Xmin: xmin, Xmax: xmax, Text: parseString(data[2])}, err
}

func parseIntervalTier(data [][]byte) (tier IntervalTier, err error){
	xmin, err := parseNumber(data[1])
	if err != nil { return }
	xmax, err := parseNumber(data[2])
	if err != nil { return }
	size, err := parseIndex(data[3])
	if err != nil { return }
	tier = IntervalTier{Class:"IntervalTier", Name: parseString(data[0]), Xmin: xmin, Xmax: xmax, Size: size}
	for i2 := 0; i2 < size; i2++ {
		interval, err := parseInterval(data[i2*3+4:i2*3+7])
		if err != nil { return tier, err}
		tier.Intervals = append(tier.Intervals, interval)
	}
	return tier, err
}