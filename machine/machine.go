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

func (m Machine) selectRowByCurrentSquare(rows []OperationRow) (OperationRow, error) {
	currentSymbol := m.tape.GetSymbol()
	for _, row := range rows {
		if row.symbolMatch(currentSymbol) {
			return row, nil
		}
	}
	return OperationRow{}, fmt.Errorf("row not found matching current square (%v)", string(currentSymbol))
}

func (m Machine) freshOperate() (Machine, error) {
	rows, ok := m.rows[m.currentConfiguration]
	if !ok {
		return Machine{}, fmt.Errorf("current configuration (%v) does not exist in table", m.currentConfiguration)
	}
	matchingRow, err := m.selectRowByCurrentSquare(rows)
	if nil != err {
		return Machine{}, err
	}
	return m.operate(matchingRow)
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
