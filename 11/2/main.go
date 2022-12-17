package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

type item struct {
	worryLevel int
	next       *item
	previous   *item
}

type list struct {
	head *item
	last *item
}

func (l *list) first() *item {
	return l.head
}

func (l *list) add(i *item) {
	if l.last == nil {
		l.head = i
		l.last = i
	} else {
		l.last.next = i
		i.previous = l.last
		l.last = i
	}
}

func (l *list) remove(i *item) {
	if l.last == i {
		l.last = i.previous
	}
	if i.next != nil {
		i.next.previous = i.previous
	}
	if i != l.head {
		i.previous.next = i.next
	} else {
		l.head = i.next
	}
	i.previous = nil
	i.next = nil
}

type monkeyTest struct {
	divisor                       int
	positiveTestMonkeyDestination int
	negativeTestMonkeyDestination int
}

func (mt *monkeyTest) test(i *item) int {
	if i.worryLevel%mt.divisor == 0 {
		return mt.positiveTestMonkeyDestination
	} else {
		return mt.negativeTestMonkeyDestination
	}
}

type monkeyOperation struct {
	operandOne string
	operandTwo string
	operation  string
}

func getValue(operand string, currentValue int) int {
	if operand == "old" {
		return currentValue
	} else {
		res, _ := strconv.Atoi(operand)
		return res
	}
}

func (mo *monkeyOperation) operate(i *item) {
	op1 := getValue(mo.operandOne, i.worryLevel)
	op2 := getValue(mo.operandTwo, i.worryLevel)
	switch mo.operation {
	case "+":
		i.worryLevel = op1 + op2
	case "-":
		i.worryLevel = op1 - op2
	case "*":
		i.worryLevel = op1 * op2
	case "/":
		i.worryLevel = op1 / op2
	}
}

type monkey struct {
	items       list
	test        monkeyTest
	operation   monkeyOperation
	inspections int
}

func (m *monkey) takeTurn(monkeys []*monkey, factor int) {
	for current := m.items.head; current != nil; {
		m.inspections++
		next := current.next
		m.operation.operate(current)
		current.worryLevel %= factor
		monkeyDestination := m.test.test(current)
		m.items.remove(current)
		monkeys[monkeyDestination].items.add(current)
		current = next
	}
}

func main() {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "11/input.txt")
	dat, e := os.ReadFile(path)

	if e != nil {
		panic(e)
	}

	lines := bytes.Split(dat, []byte("\n"))

	var monkeys []*monkey
	factor := 1
	for i := 0; i+5 < len(lines); i += 7 {
		m := monkey{}

		itemsSubs := lines[i+1][18:]
		itemsVals := bytes.Split(itemsSubs, []byte(", "))
		for j := 0; j < len(itemsVals); j++ {
			val, _ := strconv.Atoi(string(itemsVals[j]))
			m.items.add(&item{worryLevel: val})
		}

		operationSubs := lines[i+2][19:]
		opParts := bytes.Split(operationSubs, []byte(" "))
		m.operation.operandOne = string(opParts[0])
		m.operation.operation = string(opParts[1])
		m.operation.operandTwo = string(opParts[2])

		testSubs := lines[i+3][21:]
		positiveDestinationSubs := lines[i+4][29:]
		negativeDestinationSubs := lines[i+5][30:]

		m.test.divisor, _ = strconv.Atoi(string(testSubs))
		m.test.positiveTestMonkeyDestination, _ = strconv.Atoi(string(positiveDestinationSubs))
		m.test.negativeTestMonkeyDestination, _ = strconv.Atoi(string(negativeDestinationSubs))

		factor *= m.test.divisor

		monkeys = append(monkeys, &m)
	}

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkeys[j].takeTurn(monkeys, factor)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	fmt.Println(monkeys[0].inspections * monkeys[1].inspections)
}
