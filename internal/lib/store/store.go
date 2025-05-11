package store

import "sync"

type Store[T any] struct {
	mx	sync.Mutex
	d		T						// data
	ok	bool				// data existence
}

// Initialize Store
func New[T any](init T, ok bool) *Store[T] {
	return &Store[T]{
		mx: sync.Mutex{},
		d: init,
		ok: ok,
	}
}

// Return data
func (s *Store[T]) Get() (T, bool) {
	s.mx.Lock()
	defer s.mx.Unlock()
	return s.d, s.ok
}

// Return data and set ok to false
func (s *Store[T]) Pop() (T, bool) {
	s.mx.Lock()
	defer s.mx.Unlock()
	d, ok := s.d, s.ok
	s.ok = false
	return d, ok
}

// Set data
func (s *Store[T]) Set(d T) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.d = d
	s.ok = true
}

// Update data based on the previous Store
func (s *Store[T]) Update(f func(T) T) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.d = f(s.d)
	s.ok = true
}

// Non-blocking methods ///////////////////////////////////////////////////////

// The returning boolean value indicates one of the following:
//
//   - TryLock succeeded and data exists
//	 - TryLock succeeded and data does not exist
//   - TryLock failed
func (s *Store[T]) TryGet() (T, bool) {
	if s.mx.TryLock() {
		defer s.mx.Unlock()
		return s.d, s.ok
	}

	var empty T
	return empty, false
}

func (s *Store[T]) TryPop() (T, bool) {
	if s.mx.TryLock() {
		defer s.mx.Unlock()
		d, ok := s.d, s.ok
		s.ok = false
		return d, ok
	}

	var empty T
	return empty, false
}

// Return the result of TryLock()
func (s *Store[T]) TrySet(d T) bool {
	if s.mx.TryLock() {
		defer s.mx.Unlock()
		s.d = d
		s.ok = true
		return true
	}

	return false
}

// Return the result of TryLock()
func (s *Store[T]) TryUpdate(f func(T) T) bool {
	if s.mx.TryLock() {
		defer s.mx.Unlock()
		s.d = f(s.d)
		s.ok = true
		return true
	}

	return false
}
