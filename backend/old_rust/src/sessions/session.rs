#[derive(Debug, Clone)]
pub struct Session {
	pub(crate) key: String,
	pub(crate) value: String
}

impl Session {
	pub fn new(key: String, value: String) -> Session {
		Session{
			key,
			value
		}
	}
}