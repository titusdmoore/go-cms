package events

type Event struct {
	Action   func(args ...interface{})
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

func (em *EventManager) ApplyFilter(action string, opts ...interface{}) {
	for _, internalAction := range em.actions[action] {
		internalAction.Action(opts...)
	}
}

// Private Helper function to quick insert the event into is corresponding map (actions or filters) will always will be sorted
func insertEvent(event Event, events []Event) []Event {
	if len(events) == 0 {
		return []Event{event}
	}

	if len(events) == 1 {
		if event.Priority < events[0].Priority {
			return []Event{event, events[0]}
		}

		return []Event{events[0], event}
	}

	testIndex := (len(events) - 1) / 2
	left := events[:testIndex]
	right := events[testIndex:]

	if event.Priority < events[testIndex].Priority {
		return append(insertEvent(event, left), right...)
	} else if event.Priority > events[testIndex].Priority {
		return append(left, insertEvent(event, right)...)
	}

	return append(append(left, event), right...)
}
