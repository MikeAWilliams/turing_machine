package parse

import (
	"strings"

	"github.com/MikeAWilliams/turing_machine/machine"
)

func Machine(machineText string, lineDelimitor string) (machine.Machine, error) {
	lines := strings.Split(machineText, "\n")

	var runningConfig string
	var initialConfig string
	var configurations []machine.ConfigOP
	for _, line := range lines {
		row, runningConfig, err := parseLine(line, lineDelimitor, runningConfig)
		if nil != err {
			return machine.Machine{}, err
		}
		if 0 == len(configurations) {
			initialConfig = runningConfig
		}
		configurations = append(configurations, row)
	}
	return machine.NewMachine(initialConfig, configurations), nil
}
