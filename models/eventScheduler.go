package models

type EventScheduler[E Event] struct {
	PriorityQueue []E
	currentID     int
}

func (s *EventScheduler[E]) Push(e E) {
	n := len(s.PriorityQueue)

	if n == 0 {
		s.PriorityQueue = append(s.PriorityQueue, e)
		return
	}

	idx := n
	for i, event := range s.PriorityQueue {
		if e.GetStart() < event.GetStart() {
			idx = i
			break
		}
	}

	s.PriorityQueue = append(s.PriorityQueue, e)
	copy(s.PriorityQueue[idx+1:], s.PriorityQueue[idx:])
	s.PriorityQueue[idx] = e
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

func (s *EventScheduler[E]) GetNextID() int {
	s.currentID++
	return s.currentID
}
