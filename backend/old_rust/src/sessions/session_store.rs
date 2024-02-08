use crate::sessions::errors::SessionError;
use crate::sessions::Session;

pub trait SessionStore {
	fn connect(&self) -> Result<(), SessionError>;
	fn disconnect(&self) -> Result<(), SessionError>;
	fn get(&self, key: &str) -> Result<Option<Session>, SessionError>;
	fn set(&mut self, session: &Session) -> Result<(), SessionError> ;
	fn delete(&mut self, key: &str) -> Result<(), SessionError>;
}