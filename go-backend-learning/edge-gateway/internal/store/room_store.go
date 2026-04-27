package store

type RoomStore[T any] struct {
	rooms map[string]*T
}

func NewRoomStore[T any]() *RoomStore[T] {
	return &RoomStore[T]{
		rooms: map[string]*T{},
	}
}

func (s *RoomStore[T]) Save(id string, value *T) {
	s.rooms[id] = value
}

func (s *RoomStore[T]) Get(id string) (*T, bool) {
	value, ok := s.rooms[id]
	return value, ok
}

func (s *RoomStore[T]) Delete(id string) {
	delete(s.rooms, id)
}

func (s *RoomStore[T]) All() map[string]*T {
	return s.rooms
}
