package events

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventInsertEmpty(t *testing.T) {
	testingEvent := Event{nil, 0}

	em := InitializeEventManager()
	em.actions["test"] = insertEvent(Event{nil, 0}, []Event{})

	assert.Equal(t, 1, len(em.actions["test"]))
	assert.Equal(t, testingEvent, em.actions["test"][0])

	em.actions["test"] = insertEvent(Event{nil, 1}, em.actions["test"])
	assert.Equal(t, 2, len(em.actions["test"]))

	em.actions["test"] = insertEvent(Event{nil, 10}, em.actions["test"])
	em.actions["test"] = insertEvent(Event{nil, 4}, em.actions["test"])
	em.actions["test"] = insertEvent(Event{nil, 0}, em.actions["test"])
	em.actions["test"] = insertEvent(Event{nil, 0}, em.actions["test"])
	em.actions["test"] = insertEvent(Event{nil, 999}, em.actions["test"])
	em.actions["test"] = insertEvent(Event{nil, 45}, em.actions["test"])
	em.actions["test"] = insertEvent(Event{nil, 10}, em.actions["test"])

	assert.Equal(t, 9, len(em.actions["test"]))
}
