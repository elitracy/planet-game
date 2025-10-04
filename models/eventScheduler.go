package models

type EventScheduler[E Event] struct {
	PriorityQueue []E
}

func (s *EventScheduler[E]) Push(e E) {
	for i, event := range s.PriorityQueue {
		if event.GetStart() >= e.GetStart() {
			s.PriorityQueue = append(s.PriorityQueue[:i+1], s.PriorityQueue[i:]...)
			s.PriorityQueue[i] = e
		}
	}
}

func (s *EventScheduler[E]) Pop() E {
	if len(s.PriorityQueue) == 0 {
		var zero E
		return zero
	}

	event := s.PriorityQueue[0]
	s.PriorityQueue = s.PriorityQueue[1:]
	return event
}

func (s *EventScheduler[E]) Peek() E {
	if len(s.PriorityQueue) == 0 {
		var zero E
		return zero
	}

	return s.PriorityQueue[0]
}
