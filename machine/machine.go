package machine

import (
	"fmt"
)

type OperationRow struct {
	symbolMatch        SymbolMatch
	operations         Operations
	finalConfiguration string
}

func NewRow(symbol SymbolMatch, operations Operations, finalConfiguartion string) OperationRow {
	return OperationRow{symbolMatch: symbol, operations: operations, finalConfiguration: finalConfiguartion}
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
		return m.continueOperate()
	}

	return m.freshOperate()
}

func (m Machine) operate(row OperationRow) (Machine, error) {
	newOp, newTape, err := row.operations.Operate(m.tape)
	if nil != err {
		return Machine{}, fmt.Errorf("performing operation %v", err)
	}

	var newMidOperationPointer *OperationRow
	newConfiguration := row.finalConfiguration
	if !newOp.IsDone() {
		newMidOperationPointer = &OperationRow{symbolMatch: row.symbolMatch, operations: newOp, finalConfiguration: row.finalConfiguration}
	}

	return Machine{rows: m.rows, currentConfiguration: newConfiguration, midOperationRow: newMidOperationPointer, tape: newTape}, nil
}

func (m Machine) continueOperate() (Machine, error) {
	return m.operate(*m.midOperationRow)
}

func (m Machine) freshOperate() (Machine, error) {
	row, ok := m.rows[m.currentConfiguration]
	if !ok {
		return Machine{}, fmt.Errorf("current configuration does not exist in table %v", m.currentConfiguration)
	}

	// todo handle multi row configurations
	return m.operate(row[0])
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
