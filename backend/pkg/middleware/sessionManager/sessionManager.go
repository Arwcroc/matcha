package sessionManager

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"matcha/backend/pkg/slog"
	"matcha/backend/pkg/store"
)

type Config struct {
	Filter    func(c *fiber.Ctx) bool
	CookieKey string
}

type SessionManager struct {
	store  *store.SessionStore
	Config Config
}

func New(store store.SessionStore, config Config) *SessionManager {
	return &SessionManager{
		store:  &store,
		Config: config,
	}
}

func (s *SessionManager) SetStore(store *store.SessionStore) {
	s.store = store
}

func (s *SessionManager) NewHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if s.Config.Filter != nil && s.Config.Filter(c) {
			slog.Debug("SessionManager: Skipping middleware")
			return c.Next()
		}

		sessionCookie := c.Cookies(s.Config.CookieKey, "")

		session, err := (*s.store).Get(sessionCookie)
		if err != nil {
			if !errors.Is(err, fiber.ErrNotFound) {
				slog.Warn("SessionManager: Could not get session: " + err.Error())
				return err
			}
			slog.Debug("SessionManager: No session, creating...")
			session, err = (*s.store).Create(utils.UUIDv4())
			if err != nil {
				slog.Warn("SessionManager: Could not create session " + err.Error())
				return err
			}
			c.Cookie(&fiber.Cookie{
				Name:  s.Config.CookieKey,
				Value: session.GetKey(),
			})
			slog.Debug("SessionManager: session created")
		}

		if c.Locals("session", &session) == nil {
			return fiber.ErrInternalServerError
		}
		next := c.Next()

		if (*s.store).Set(&session) != nil {
			slog.Warn("SessionManager: Could not save session " + err.Error())
			return err
		}

		return next
	}
}
