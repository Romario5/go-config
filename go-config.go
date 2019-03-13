package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var props map[string]string

// Loads configuration from file with given file name.
// Format of the configuration:
// - comments must begin from hash ($).
// - each property separated from value with equals symbol (=)
// - property name and value can contain spaces at boundaries (spaces will be trimmed)
// Returns number of read properties.
func LoadFile(fileName string) int64 {
	if props == nil {
		props = make(map[string]string)
	}

	f, err := os.OpenFile(fileName, os.O_RDONLY, 0)
	defer func() {
		err = f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		fmt.Println("Error reading file " + fileName + ":", err)
		return 0
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var propsLoaded int64 = 0
	for scanner.Scan() {

		line := scanner.Bytes()

		// Skip comments
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		key, value := parseProp(line)

		if key != "" {
			props[key] = value
			propsLoaded++
		}
		line = line[:0]
	}

	return propsLoaded
}

// Parses configuration line.
func parseProp(line []byte) (key, value string) {
	// Skip comments
	if len(line) == 0 || line[0] == '#' {
		return
	}

	lineArr := strings.Split(string(line), "=")
	key = strings.Trim(lineArr[0], " \n\r")

	if key != "" {
		if len(lineArr) > 1 {
			value = strings.Trim(lineArr[1], " \n\r")
		}
	}

	return
}

func GetString(param string, defaultValue string) string {
	if props == nil {
		return defaultValue
	}
	val, ok := props[param]
	if ok {
		return val
	}
	return defaultValue
}

func GetUint64(param string, defaultValue uint64) uint64 {
	if props == nil {
		return defaultValue
	}
	val, ok := props[param]
	if ok {
		i, _ := strconv.ParseUint(val, 10, 64)
		return i
	}
	return defaultValue
}

func GetInt64(param string, defaultValue int64) int64 {
	if props == nil {
		return defaultValue
	}
	val, ok := props[param]
	if ok {
		i, _ := strconv.ParseInt(val, 10, 64)
		return i
	}
	return defaultValue
}

func GetUint32(param string, defaultValue uint32) uint32 {
	if props == nil {
		return defaultValue
	}
	val, ok := props[param]
	if ok {
		i, _ := strconv.ParseUint(val, 10, 32)
		return uint32(i)
	}
	return defaultValue
}

func GetInt32(param string, defaultValue int32) int32 {
	if props == nil {
		return defaultValue
	}
	val, ok := props[param]
	if ok {
		i, _ := strconv.ParseInt(val, 10, 32)
		return int32(i)
	}
	return defaultValue
}

func GetUint16(param string, defaultValue uint16) uint16 {
	if props == nil {
		return defaultValue
	}
	val, ok := props[param]
	if ok {
		i, _ := strconv.ParseUint(val, 10, 16)
		return uint16(i)
	}
	return defaultValue
}

func GetInt16(param string, defaultValue int16) int16 {
	if props == nil {
		return defaultValue
	}
	val, ok := props[param]
	if ok {
		i, _ := strconv.ParseInt(val, 10, 16)
		return int16(i)
	}
	return defaultValue
}

func GetUint8(param string, defaultValue uint8) uint8 {
	if props == nil {
		return defaultValue
	}
	val, ok := props[param]
	if ok {
		i, _ := strconv.ParseUint(val, 10, 8)
		return uint8(i)
	}
	return defaultValue
}

func GetInt8(param string, defaultValue int8) int8 {
	if props == nil {
		return defaultValue
	}
	val, ok := props[param]
	if ok {
		i, _ := strconv.ParseInt(val, 10, 8)
		return int8(i)
	}
	return defaultValue
}

func GetFloat32(param string, defaultValue float32) float32 {
	if props == nil {
		return defaultValue
	}
	val, ok := props[param]
	if ok {
		i, _ := strconv.ParseFloat(val, 32)
		return float32(i)
	}
	return defaultValue
}

func GetFloat64(param string, defaultValue float64) float64 {
	if props == nil {
		return defaultValue
	}
	val, ok := props[param]
	if ok {
		i, _ := strconv.ParseFloat(val, 64)
		return i
	}
	return defaultValue
}

func GetBool(param string, defaultValue bool) bool {
	if props == nil {
		return defaultValue
	}
	val, ok := props[param]
	if ok {
		return val == "True" || val == "true" || val == "1" || val == "yes" || val == "Yes"
	}
	return false
}
