package parse

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MikeAWilliams/turing_machine/machine"
)

func parseLine(line string, delimiter string, runningConfigName string) (machine.ConfigOP, string, error) {
	parts := strings.Split(line, delimiter)
	if 4 != len(parts) {
		return machine.ConfigOP{}, "", fmt.Errorf("found %v parts for configuration line 4 required", len(parts))
	}

	configName := parts[0]
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
