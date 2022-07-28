package main

import (
	"encoding/json"
	"fmt"
)

type mapStack struct {
	top  int
	vals []map[string]*string // A nil entry indicates that the key has been deleted
}

func (m *mapStack) String() string {
	var ret string
	for i := 0; i <= m.top; i++ {
		vJSON, _ := json.Marshal(m.vals[i])
		ret += fmt.Sprintf("%d - %s\n", i, string(vJSON))
	}

	return ret
}

func newMapStack() *mapStack {
	return &mapStack{
		top:  0,
		vals: []map[string]*string{map[string]*string{}},
	}
}

func (m *mapStack) set(key, val string) {
	m.vals[m.top][key] = &val
}

func (m *mapStack) get(key string) string {
	v := m.currentState(0)[key]
	if v == nil {
		return "Nil"
	}
	return *v
}

func (m *mapStack) unSet(key string) {
	if m.top > 0 {
		m.vals[m.top][key] = nil
	} else {
		delete(m.vals[m.top], key)
	}
}

func (m *mapStack) numEqualTo(targetValue string) int {
	count := 0
	for _, value := range m.currentState(0) {
		if value == nil {
			continue
		}
		if targetValue == *value {
			count++
		}
	}

	return count
}

func (m *mapStack) begin() {
	m.top++
	if m.top >= len(m.vals) {
		m.vals = append(m.vals, map[string]*string{})
	} else {
		m.vals[m.top] = map[string]*string{}
	}
}

func (m *mapStack) commit() bool {
	if m.top == 0 {
		// Nothing to commit
		return false
	}

	newTop := m.currentState(m.top - 1)

	m.top--

	m.vals[m.top] = newTop
	return true
}

func (m *mapStack) rollBack() bool {
	if m.top == 0 {
		// Don't allow removal of the base map
		return false
	}
	m.top--
	return true
}

func (m *mapStack) currentState(start int) map[string]*string {
	ret := make(map[string]*string, len(m.vals[0]))

	for i := start; i <= m.top; i++ {
		for key, value := range m.vals[i] {
			if start == 0 && value == nil {
				delete(ret, key)
			} else {
				ret[key] = value
			}
		}
	}

	return ret
}
