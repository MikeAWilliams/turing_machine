package machine

import (
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
		return m.continueOperate()
	}

	return m.freshOperate()
}

func (m Machine) continueOperate() (Machine, error) {
	newOp, newTape, err := m.midOperationRow.operations.Operate(m.tape)
	if nil != err {
		return Machine{}, fmt.Errorf("performing operation %v", err)
	}

	var newMidOperationPointer *OperationRow
	newConfiguration := m.midOperationRow.finalConfiguration
	if !newOp.IsDone() {
		newMidOperationPointer = &OperationRow{symbol: m.midOperationRow.symbol, operations: newOp, finalConfiguration: m.midOperationRow.finalConfiguration}
	}

	return Machine{rows: m.rows, currentConfiguration: newConfiguration, midOperationRow: newMidOperationPointer, tape: newTape}, nil
}

func (m Machine) freshOperate() (Machine, error) {
	row, ok := m.rows[m.currentConfiguration]
	if !ok {
		return Machine{}, fmt.Errorf("current configuration does not exist in table %v", m.currentConfiguration)
	}

	newOp, newTape, err := row[0].operations.Operate(m.tape)
	if nil != err {
		return Machine{}, fmt.Errorf("performing operation %v", err)
	}
	var newMidOperationPointer *OperationRow
	newConfiguration := row[0].finalConfiguration
	if !newOp.IsDone() {
		newMidOperationPointer = &OperationRow{symbol: row[0].symbol, operations: newOp, finalConfiguration: row[0].finalConfiguration}
	}

	return Machine{rows: m.rows, currentConfiguration: newConfiguration, midOperationRow: newMidOperationPointer, tape: newTape}, nil
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
