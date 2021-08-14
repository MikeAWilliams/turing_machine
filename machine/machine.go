package machine

import (
	"errors"
	"fmt"
)

type OperationRow struct {
	symbol             string
	operations         Operations
	finalConfiguration string
}

func NewRow(symbol string, operations Operations, finalConfiguartion string) OperationRow {
	return OperationRow{symbol: symbol, operations: operations, finalConfiguration: finalConfiguartion}
}

type ConfigOP struct {
	configuration string
	row           OperationRow
}

func NewConfig(configuration string, row OperationRow) ConfigOP {
	return ConfigOP{configuration: configuration, row: row}
}

type Machine struct {
	rows                 map[string][]OperationRow
	currentConfiguration string
	midOperationRow      *OperationRow
	tape                 Tape
}

func (m Machine) Operate() (Machine, error) {
	if nil != m.midOperationRow {
		return Machine{}, errors.New("mid operation not implementd")
	}
	row, ok := m.rows[m.currentConfiguration]
	if !ok {
		return Machine{}, fmt.Errorf("current configuration does not exist in table %v", m.currentConfiguration)
	}
	_, newTape, err := row[0].operations.Operate(m.tape)
	if nil != err {
		return Machine{}, fmt.Errorf("performing operation %v", err)
	}

	return Machine{rows: m.rows, currentConfiguration: row[0].finalConfiguration, midOperationRow: nil, tape: newTape}, nil
}

func (m Machine) TapeAsString() string {
	return string(m.tape.squares)
}

func NewMachine(initialConfiguration string, rows []ConfigOP) Machine {
	result := Machine{}
	result.rows = map[string][]OperationRow{}
	result.currentConfiguration = initialConfiguration
	result.midOperationRow = nil
	result.tape = NewTape()

	for _, row := range rows {
		existingRow, ok := result.rows[row.configuration]
		if ok {
			result.rows[row.configuration] = append(existingRow, row.row)
		} else {
			result.rows[row.configuration] = []OperationRow{row.row}
		}
	}
	return result
}
