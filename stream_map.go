package mtserver

import (
	"sync"

	"google.golang.org/grpc"
)

type StreamMap[K comparable, T grpc.ServerStream, P any] struct {
	m  map[K][]*T
	mu sync.RWMutex
}

func NewStreamMap[K comparable, T grpc.ServerStream, P any]() *StreamMap[K, T, P] {
	return &StreamMap[K, T, P]{m: make(map[K][]*T)}
}

func (s *StreamMap[K, T, P]) Add(key K, stream T) error {
	s.mu.Lock()
	if s.m[key] == nil {
		s.m[key] = []*T{}
	}
	s.m[key] = append(s.m[key], &stream)
	s.mu.Unlock()

	for {
		<-stream.Context().Done()
		s.Remove(key, stream)
		return nil
	}
}

func (s *StreamMap[K, T, P]) Remove(key K, stream T) {
	if s.m[key] == nil {
		return
	}

	s.mu.RLock()
	streams := s.m[key]
	s.mu.RUnlock()

	ptr := &stream

	for i, st := range streams {
		if st == ptr {
			s.mu.Lock()
			s.m[key] = append(streams[:i], streams[i+1:]...)
			s.mu.Unlock()
			return
		}
	}
}

func (s *StreamMap[K, T, P]) Send(key K, msg P) {
	if s.m[key] == nil {
		return
	}

	s.mu.RLock()
	streams := s.m[key]
	s.mu.RUnlock()

	for _, stream := range streams {
		var t T = *stream
		if err := t.Context().Err(); err != nil {
			s.Remove(key, t)
			continue
		}
		t.SendMsg(msg)
	}
}
