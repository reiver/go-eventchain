package eventchaindriver


import (
	"reflect"
	"time"

	"testing"
)


func TestMock(t *testing.T) {

	tests := []struct{
		MockedEvents []Event
	}{
		{
			MockedEvents: []Event{},
		},
		{
			MockedEvents: []Event{
				&MockEvent{
					MockedID: "apple",
				},
			},
		},
		{
			MockedEvents: []Event{
				&MockEvent{
					MockedID: "apple",
				},
				&MockEvent{
					MockedID: "banana",
				},
			},
		},
		{
			MockedEvents: []Event{
				&MockEvent{
					MockedID: "apple",
				},
				&MockEvent{
					MockedID: "banana",
				},
				&MockEvent{
					MockedID: "cherry",
				},
			},
		},
		{
			MockedEvents: []Event{
				&MockEvent{
					MockedID:                       "apple",
					MockedWhenCreated:              time.Now().Add(1 * time.Hour),
					MockedPrevID:                   "",
					MockedPrevIDErr:                internalNotFoundComplainer{},
					MockedInteractionIdentifier:    "some.unique_VALUE",
					MockedInteractionIdentifierErr: nil,
					MockedWhenHappened:             time.Now().Add(5 * time.Hour),
					MockedEventName:                "HUNGRY_HUMAN_REGISTERED",
					MockedVersion:                  "1.0.0",
					MockedData:                     "one two three",
				},
				&MockEvent{
					MockedID:          "banana",
					MockedWhenCreated: time.Now().Add(2 * time.Hour),
					MockedPrevID:      "apple",
					MockedPrevIDErr:    nil,
					MockedInteractionIdentifier:    "",
					MockedInteractionIdentifierErr: internalNotFoundComplainer{},
					MockedWhenHappened:             time.Now().Add(8 * time.Hour),
					MockedEventName:                "HUNGRY_HUMAN_ATE_BREAKFAST",
					MockedVersion:                  "2.3.0",
					MockedData:                     map[string]interface{}{
						"one":   1,
						"two":   2,
						"three": 3,
					},
				},
				&MockEvent{
					MockedID:          "cherry",
					MockedWhenCreated: time.Now().Add(3 * time.Hour),
					MockedPrevID:      "banana",
					MockedPrevIDErr:    nil,
					MockedInteractionIdentifier:    "some.OTHER.unique_VALUE",
					MockedInteractionIdentifierErr: nil,
					MockedWhenHappened:             time.Now().Add(13 * time.Hour),
					MockedEventName:                "HUNGRY_HUMAN_ATE_SMAKE",
					MockedVersion:                  "5.8.13",
					MockedData:                     map[string]interface{}{
						"once":   1,
						"twice":  2,
						"thrice": 3,
						"fource": []interface{}{"four", 4, 4.0, "FOUR"},
					},
				},
			},
		},
	}


	for testNumber, test := range tests {

		mockedEvents := append([]Event(nil), test.MockedEvents...)

		mockEvents := MockEvents{
			Events: mockedEvents,
		}

		mockEventStore :=  MockEventStore{
			LoadEventsFunc: func(eventChainID string)(Events, error) {
				return &mockEvents, nil
			},
		}

		mockDataStore := MockDataStore{
			OpenEventStoreFunc: func(eventStoreName string, eventNames ...string)(EventStore, error) {
				return &mockEventStore, nil
			},
		}

		mockDriver := MockDriver{
			OpenDataStoreFunc: func(args ...interface{}) (DataStore, error) {
				return &mockDataStore, nil
			},
		}


		var driver = &mockDriver

		dataStore, err := driver.OpenDataStore()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		if nil == dataStore {
			t.Errorf("For test #%d, did not expect nil, but actually got: %v", testNumber, dataStore)
			continue
		}

		eventStore, err := dataStore.OpenEventStore("does_not_matter",
			"ATE_AN_APPLE",
			"ATE_A_BANANA",
			"ATE_A_CHERRY",
		)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		if nil == eventStore {
			t.Errorf("For test #%d, did not expect nil, but actually got: %v", testNumber, eventStore)
			continue
		}

		events, err := eventStore.LoadEvents("it.matters_not")
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		if nil == events {
			t.Errorf("For test #%d, did not expect nil, but actually got: %v", testNumber, events)
			continue
		}

		var eventsInASlice []Event
		{
			actual := 0
			for events.Next() {
				event, err := events.Event()
				if nil != err {
					t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
					continue
				}

				eventsInASlice = append(eventsInASlice, event)
				actual++
			}
			if err := events.Err(); nil != err {
				t.Errorf("For test #%d, did not expect an error. but actually got one: (%T) %v", testNumber, err, err)
				continue
			}

			if expected := len(mockedEvents); expected != actual {
				t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
				continue
			}
		}

		for eventNumber, actualEvent := range test.MockedEvents {

			expectedEvent := eventsInASlice[eventNumber]

			if expected, actual := expectedEvent.ID(), actualEvent.ID(); expected != actual {
				t.Errorf("For test #%d and event #%d, expected %q, but actually got %q.", testNumber, eventNumber, expected, actual)
				continue
			}
			if expected, actual := expectedEvent.WhenCreated(), actualEvent.WhenCreated(); ! expected.Equal(actual) {
				t.Errorf("For test #%d and event #%d, expected %q, but actually got %q.", testNumber, eventNumber, expected, actual)
				continue
			}
			{
				expected, err := expectedEvent.PrevID()
				if nil != err {
					if _, notFound := err.(NotFoundComplainer); !notFound {
						t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
						continue
					}
				}

				actual, err := actualEvent.PrevID()
				if nil != err {
					if _, notFound := err.(NotFoundComplainer); !notFound {
						t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
						continue
					}
				}

				if expected != actual {
					t.Errorf("For test #%d and event #%d, expected %q, but actually got %q.", testNumber, eventNumber, expected, actual)
					continue
				}
			}
			{
				expected, err := expectedEvent.InteractionIdentifier()
				if nil != err {
					if _, notFound := err.(NotFoundComplainer); !notFound {
						t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
						continue
					}
				}

				actual, err := actualEvent.InteractionIdentifier()
				if nil != err {
					if _, notFound := err.(NotFoundComplainer); !notFound {
						t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
						continue
					}
				}

				if expected != actual {
					t.Errorf("For test #%d and event #%d, expected %q, but actually got %q.", testNumber, eventNumber, expected, actual)
					continue
				}
			}
			if expected, actual := expectedEvent.WhenHappened(), actualEvent.WhenHappened(); ! expected.Equal(actual) {
				t.Errorf("For test #%d and event #%d, expected %q, but actually got %q.", testNumber, eventNumber, expected, actual)
				continue
			}
			if expected, actual := expectedEvent.EventName(), actualEvent.EventName(); expected != actual {
				t.Errorf("For test #%d and event #%d, expected %q, but actually got %q.", testNumber, eventNumber, expected, actual)
				continue
			}
			if expected, actual := expectedEvent.Version(), actualEvent.Version(); expected != actual {
				t.Errorf("For test #%d and event #%d, expected %q, but actually got %q.", testNumber, eventNumber, expected, actual)
				continue
			}
			if expected, actual := expectedEvent.Data(), actualEvent.Data(); !reflect.DeepEqual(expected, actual) {
				t.Errorf("For test #%d and event #%d, expected %#v, but actually got %#v.", testNumber, eventNumber, expected, actual)
				continue
			}
		}
	}
}
