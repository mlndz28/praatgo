// Package praatgo provides tools to interpret and process
// Praat output and input objects as defined in 
// https://www.fon.hum.uva.nl/praat/manual/Types_of_objects 
package praatgo

import (
	"bytes"
	"errors"
	"regexp"
)

// DeserializeTextGrid recurses the content of a TextGrid file 
// and returns its values into tg  
func DeserializeTextGrid(input []byte) (tg TextGrid, err error) {
	content := bytes.ReplaceAll(input, []byte("\x00"), []byte{}) // praat usually writes null bytes between characters
	pattern := regexp.MustCompile(`"(.*?)"`)
	headers := pattern.FindAllSubmatch(content, 2)
	if !bytes.Equal(headers[1][1], []byte("TextGrid")) {
		return tg, errors.New("Not a TextGrid file")
	}
	tg.FileType = string(headers[0][1])
	cursor := pattern.FindAllSubmatchIndex(content, 2)[1][1]
	content = content[cursor:]
	pattern = regexp.MustCompile(`\s[0-9\.]+|"(.*?)"|<exists>|<absent>`)
	data := pattern.FindAll(content, -1)
	tg.Xmin, err = parseNumber(data[0])
	if err != nil { return }
	tg.Xmax, err = parseNumber(data[1])
	if err != nil { return }
	tg.Tiers, err = parseBool(data[2])
	if err != nil { return }
	tg.Size, err = parseIndex(data[3])
	if err != nil { return }
	j := 4
	for i1 := 0; i1 < tg.Size; i1++ {
		switch string(data[j]) {
		case `"IntervalTier"`:
			tier,err := parseIntervalTier(data[j+1:])
			if err != nil { return tg, err}
			j += 5 + 3*len(tier.Intervals)
			tg.Item = append(tg.Item, tier)
		default:
			return tg, errors.New("Bad format")
		}
	}
	return
}
