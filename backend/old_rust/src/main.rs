mod sessions;
mod middleware;
mod stores;

use nickel::{Action, HttpRouter, MiddlewareResult, Nickel, NickelError, Request, Response};
use nickel::status::StatusCode;
use crate::middleware::error_logger;

fn todo<'a>(_req: &mut Request, res: Response<'a>) -> MiddlewareResult<'a> {
	Err(NickelError::new(res, "TODO", StatusCode::NotImplemented))
}

fn auth_router() -> nickel::Router {
	let mut router = Nickel::router();

	router.post("/auth/login", todo);
	router.get("/auth/logout", todo);
	router.get("/auth/whoami", todo);

	router
}

fn user_router() -> nickel::Router {
	let mut router = Nickel::router();

	router.get("/users", todo);
	router.post("/users", todo);
	router.get("/users/:id", todo);
	router.put("/users/:id", todo);
	router.delete("/users/:id", todo);

	router
}

fn main() {
	env_logger::init();
	let mut server = Nickel::new();
	let session_manager = middleware::SessionManager::new(
		stores::Memory::new()
	);
	server.utilize(session_manager);
	server.utilize(middleware::Logger);
	server.utilize(auth_router());
	server.utilize(user_router());

	let error_logger: fn(&mut NickelError<()>, &mut Request<()>) -> Action = error_logger;
	server.handle_error(error_logger);


	server.listen("127.0.0.1:3000").unwrap();
}
