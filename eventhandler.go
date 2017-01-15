package minatsubot

import (
	"reflect"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type EventHandler struct {
	mu       *sync.Mutex
	rmu      *sync.RWMutex
	handlers map[interface{}][]reflect.Value
}

func (e *EventHandler) validateHandler(handler interface{}) reflect.Type {
	handlerType := reflect.TypeOf(handler)
	if handlerType.NumIn() != 1 {
		log.Error("Unable to add event handler, handler must be of the type func(*minatsubot.EventType)")
		return nil
	}

	eventType := handlerType.In(0)

	if eventType.Kind() == reflect.Interface {
		eventType = nil
	}
	return eventType
}

func (e *EventHandler) initialize() {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.handlers != nil {
		return
	}
	e.handlers = map[interface{}][]reflect.Value{}
}

func (e *EventHandler) AddHandler(handler interface{}) func() {
	e.initialize()

	eventType := e.validateHandler(handler)

	e.mu.Lock()
	defer e.mu.Unlock()

	h := reflect.ValueOf(handler)
	e.handlers[eventType] = append(e.handlers[eventType], h)

	// This must be done as we need a consistent reference to the
	// reflected value, otherwise a RemoveHandler method would have
	// been nice.
	return func() {
		e.mu.Lock()
		defer e.mu.Unlock()

		handlers := e.handlers[eventType]
		for i, v := range handlers {
			if h == v {
				e.handlers[eventType] = append(handlers[:i], handlers[i+1:]...)
				return
			}
		}
	}
}

// Handle is the "emitter", call this function whenever you want to call an event
func (e *EventHandler) Handle(event interface{}) {
	e.rmu.RLock()
	defer e.rmu.RUnlock()

	if e.handlers == nil {
		return
	}

	handlerParameters := []reflect.Value{reflect.ValueOf(event)}

	// Call all handlers that listens to all events
	if handlers, ok := e.handlers[nil]; ok {
		for _, handler := range handlers {
			go handler.Call(handlerParameters)
		}
	}

	if handlers, ok := e.handlers[reflect.TypeOf(event)]; ok {
		for _, handler := range handlers {
			go handler.Call(handlerParameters)
		}
	}
}

func (e *EventHandler) handler(s *discordgo.Session, event interface{}) {
	e.Handle(event)
}
