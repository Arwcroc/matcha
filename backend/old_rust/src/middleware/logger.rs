use std::io::Write;
use async_trait::async_trait;
use nickel::{Action, Middleware, MiddlewareResult, NickelError, Request, Response};
use log::{debug, info, warn};

pub struct Logger;

#[async_trait]
impl<D: Send + 'static + Sync> Middleware<D> for Logger {
	fn invoke<'mw>(&'mw self, req: &mut Request<'mw, '_, D>, res: Response<'mw, D>) -> MiddlewareResult<'mw, D> {
		info!("{}: {} {}",
			req.origin.remote_addr,
			req.origin.method,
			req.origin.uri,
		);
		debug!("{:?}", req.origin.headers);

		res.next_middleware()
	}
}

pub fn error_logger(err: &mut NickelError, req: &mut Request) -> Action {
	let message = err.message.clone();
	if let Some(ref mut res) = err.stream {
		let _ = res.write_all(message.as_bytes());
		warn!("{}: {} {} -- {}",
			req.origin.remote_addr,
			req.origin.method,
			req.origin.uri,
			res.status()
		);
		debug!("{:?}", req.origin.headers);
		debug!("{:?}", res.headers());
		return Action::Halt(())
	}
	Action::Continue(())
}
