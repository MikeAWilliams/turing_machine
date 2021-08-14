package machine

import "errors"

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
	return Machine{}, errors.New("not implementd")
}

func NewMachine(initialConfiguration string, rows []ConfigOP) Machine {
	result := Machine{}
	result.currentConfiguration = initialConfiguration
	result.tape = NewTape()
	result.midOperationRow = nil
	result.rows = map[string][]OperationRow{}

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
