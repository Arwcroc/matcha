use std::collections::HashMap;
use crate::sessions::errors::SessionError;
use crate::sessions::{Session, SessionStore};

pub struct Memory {
	sessions: HashMap<String, String>
}

impl Memory {
	pub fn new() -> Memory {
		Memory{
			sessions: HashMap::new()
		}
	}
}

impl SessionStore for Memory {
	fn connect(&self) -> Result<(), SessionError> {
		Ok(())
	}

	fn disconnect(&self) -> Result<(), SessionError> {
		Ok(())
	}

	fn get(&self, key: &str) -> Result<Option<Session>, SessionError> {
		Ok(
			self.sessions.get(key).map(|value| Session::new(
					key.to_string(),
					value.to_string()
				))
		)
	}

	fn set(&mut self, session: &Session) -> Result<(), SessionError> {
		self.sessions.insert(session.key.clone(), session.value.clone());
		Ok(())
	}

	fn delete(&mut self, key: &str) -> Result<(), SessionError> {
		self.sessions.remove(key);
		Ok(())
	}

	fn clone(self) -> dyn SessionStore {
		Memory{
			sessions: self.sessions
		}
	}
}