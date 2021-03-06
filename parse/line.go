package parse

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MikeAWilliams/turing_machine/machine"
)

func splitLine(line string, delimiter string) []string {
	parts := strings.Split(line, delimiter)
	var result []string
	for index, part := range parts {
		if 0 == index {
			result = append(result, part)
		} else if len(part) > 0 {
			result = append(result, part)
		}
	}
	return result
}

func parseLine(line string, delimiter string, runningConfigName string) (machine.ConfigOP, string, error) {
	parts := splitLine(line, delimiter)
	if 4 != len(parts) {
		return machine.ConfigOP{}, "", fmt.Errorf("found %v parts for configuration line 4 required", len(parts))
	}

	configName := strings.TrimSpace(parts[0])
	if 0 == len(configName) && 0 == len(runningConfigName) {
		return machine.ConfigOP{}, "", errors.New("both line config name and running config name are empty")
	}
	if 0 == len(configName) {
		configName = runningConfigName
	}

	symbolMatcher, err := parseSymbolMatch(parts[1])
	if nil != err {
		return machine.ConfigOP{}, "", err
	}

	operations, err := parseOperations(parts[2])
	if nil != err {
		return machine.ConfigOP{}, "", err
	}

	if 0 == len(parts[3]) {
		return machine.ConfigOP{}, "", errors.New("final configuration cannot be empty")
	}

	return machine.NewConfig(configName, machine.NewRow(symbolMatcher, machine.NewOperations(operations), parts[3])), configName, nil
}
