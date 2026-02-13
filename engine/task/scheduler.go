package task

type TaskID int

type TaskScheduler[T Task] struct {
	PriorityQueue []T
	currentID     TaskID
}

func (s *TaskScheduler[T]) Push(t T) {

	t.SetID(s.GetNextID())

	n := len(s.PriorityQueue)

	if n == 0 {
		s.PriorityQueue = append(s.PriorityQueue, t)
		return
	}

	idx := n
	for i, task := range s.PriorityQueue {
		if t.GetStartTick() < task.GetStartTick() {
			idx = i
			break
		}
	}

	s.PriorityQueue = append(s.PriorityQueue, t)
	copy(s.PriorityQueue[idx+1:], s.PriorityQueue[idx:])
	s.PriorityQueue[idx] = t
}

func (s *TaskScheduler[T]) Pop() T {
	if len(s.PriorityQueue) == 0 {
		var zero T
		return zero
	}

	task := s.PriorityQueue[0]
	s.PriorityQueue = s.PriorityQueue[1:]
	return task
}

func (s *TaskScheduler[T]) Peek() T {
	if len(s.PriorityQueue) == 0 {
		var zero T
		return zero
	}

	return s.PriorityQueue[0]
}

func (s *TaskScheduler[T]) GetNextID() TaskID {
	s.currentID++
	return s.currentID
}
