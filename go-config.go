package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"utils/log"
)

var props map[string]string

// Loads configuration from file with given file name.
// Format of the configuration:
// - comments must begin from hash ($).
// - each property separated from value with equals symbol (=)
// - property name and value can contain spaces at boundaries (spaces will be trimmed)
func LoadFile(fileName string) {
	if props == nil {
		props = make(map[string]string)
	}

	f, err := os.OpenFile(fileName, os.O_RDONLY, 0)
	defer f.Close()

	if err != nil {
		log.Error("Error reading configuration file" + fileName)
		return
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

	fmt.Println("Configuration loaded from", fileName, "("+strconv.FormatInt(propsLoaded, 10)+" properties)")
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
		return ""
	}
	val, ok := props[param]
	if ok {
		return val
	}
	return defaultValue
}

func GetUint64(param string, defaultValue uint64) uint64 {
	if props == nil {
		return 0
	}
	val, ok := props[param]
	if ok {
		i, _ := strconv.ParseUint(val, 10, 64)
		return i
	}
	return defaultValue
}

func GetBool(param string, defaultValue bool) bool {
	if props == nil {
		return false
	}
	val, ok := props[param]
	if ok {
		return val == "True" || val == "true"
	}
	return false
}
