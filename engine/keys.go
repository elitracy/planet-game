package engine

import (
	"fmt"
	"strings"
)

type KeyAction int

const (
	Back KeyAction = iota
	Select
	Up
	Down
	Scout
	Colonize
	Search
	InsertText
	Quit
)

type KeyBindings struct {
	bindings map[KeyAction]string
	order    []KeyAction
}

func NewKeyBindings() *KeyBindings {
	return &KeyBindings{
		bindings: make(map[KeyAction]string),
	}
}

func (kb *KeyBindings) Set(action KeyAction, key string) *KeyBindings {
	if _, ok := kb.bindings[action]; !ok {
		kb.order = append(kb.order, action)
	}
	kb.bindings[action] = key
	return kb
}

func (kb *KeyBindings) Unset(action KeyAction) *KeyBindings {
	delete(kb.bindings, action)
	for i, a := range kb.order {
		if a == action {
			kb.order = append(kb.order[:i], kb.order[i+1:]...)
			break
		}
	}
	return kb
}

func (kb *KeyBindings) Get(action KeyAction) string {
	if key, ok := kb.bindings[action]; ok {
		return key
	}
	return ""
}

func (kb *KeyBindings) Clear() *KeyBindings {
	kb.bindings = make(map[KeyAction]string)
	kb.order = []KeyAction{}
	return kb
}

func (kb *KeyBindings) String() string {
	var parts []string
	for _, action := range kb.order {
		if key, ok := kb.bindings[action]; ok {
			parts = append(parts, fmt.Sprintf("%v: %v", action.String(), key))
		}

	}
	return strings.Join(parts, " | ")
}
