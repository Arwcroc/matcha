package memory

import (
	"github.com/gofiber/fiber/v2"
	"matcha/backend/pkg/store"
	"sync"
)

type Store struct {
	store.SessionStore
	lock     *sync.RWMutex
	sessions map[string]store.Session
}

func New() *Store {
	return &Store{
		lock:     &sync.RWMutex{},
		sessions: make(map[string]store.Session),
	}
}

func (s *Store) Connect() error { return nil }

func (s *Store) Disconnect() error { return nil }

func (s *Store) Get(key string) (store.Session, error) {
	var err error = nil

	s.lock.RLock()
	session, ok := s.sessions[key]
	s.lock.RUnlock()

	if !ok {
		err = fiber.ErrNotFound
	}
	return session, err
}

func (s *Store) Set(session *store.Session) error {
	s.lock.Lock()
	s.sessions[session.GetKey()] = *session
	s.lock.Unlock()
	return nil
}

func (s *Store) Delete(key string) error {
	s.lock.Lock()
	delete(s.sessions, key)
	s.lock.Unlock()
	return nil
}

func (s *Store) Create(key string) (store.Session, error) {
	s.sessions[key] = store.CreateSession(key)
	return s.sessions[key], nil
}
