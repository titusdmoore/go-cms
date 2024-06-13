package events

import (
	"fmt"
	"slices"
)

type Event struct {
	Action   func(args ...interface{}) (interface{}, error)
	Priority int32
}

type EventManager struct {
	actions map[string][]Event
	filters map[string][]Event
}

func InitializeEventManager() EventManager {
	actions := make(map[string][]Event)
	filters := make(map[string][]Event)

	return EventManager{
		actions: actions,
		filters: filters,
	}
}

func (em *EventManager) DoAction(action string, opts ...interface{}) {
	for _, internalAction := range em.actions[action] {
		internalAction.Action(opts...)
	}
}

// Needs to return the value of the filter, may have to be string according to https://templ.guide/syntax-and-usage/expressions/#functions
func (em *EventManager) ApplyFilter(action string, opts ...interface{}) interface{} {
	var err error

	for _, internalAction := range em.actions[action] {
		// we need to figure out a way to return the values and pass to next in the chain
		internalAction.Action(opts...)
	}

	if err != nil {
		fmt.Println(err)
	}

	return opts[0]
}

func (em *EventManager) AddAction(action string, priority int32, actionFunc func(args ...interface{})) {
	em.actions[action] = insertEvent(Event{func(args ...interface{}) (interface{}, error) {
		actionFunc(args...)
		return nil, nil
	}, priority}, em.actions[action])
}

func (em *EventManager) AddFilter(action string, priority int32, actionFunc func(args ...interface{}) (interface{}, error)) {
	em.filters[action] = insertEvent(Event{actionFunc, priority}, em.filters[action])
}

// Private Helper function to quick insert the event into is corresponding map (actions or filters) will always will be sorted
// Currently this function sucks, so I will probably refactor it later
func insertEvent(event Event, events []Event) []Event {
	if len(events) == 0 {
		return append(events, event)
	}

	idx := 0
	for events[idx].Priority < event.Priority {
		if idx+1 == len(events) {
			idx = len(events)
			break
		}

		idx++
	}

	return slices.Insert(events, idx, event)
}
