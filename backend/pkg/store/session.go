package store

import "sync"

type Session struct {
	lock  *sync.RWMutex
	key   string
	value map[string]interface{}
}

func CreateSession(key string) Session {
	return Session{
		lock:  &sync.RWMutex{},
		key:   key,
		value: make(map[string]interface{}),
	}
}

func (s *Session) GetKey() string {
	s.lock.RLock()
	key := s.key
	s.lock.RUnlock()
	return key
}

func (s *Session) Set(key string, value interface{}) {
	s.lock.Lock()
	s.value[key] = value
	s.lock.Unlock()
}

func (s *Session) Get(key string) interface{} {
	s.lock.RLock()
	value := s.value[key]
	s.lock.RUnlock()
	return value
}

func (s *Session) Delete(key string) {
	s.lock.Lock()
	delete(s.value, key)
	s.lock.Unlock()
}

type SessionStore interface {
	Connect() error
	Disconnect() error
	Get(string) (Session, error)
	Set(*Session) error
	Create(string) (Session, error)
	Delete(string) error
}
