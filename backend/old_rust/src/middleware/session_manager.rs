use async_trait::async_trait;
use nickel::{Middleware, MiddlewareResult, Request, Response};
use crate::sessions::{SessionStore};

pub struct SessionManager<Store: SessionStore> {
    driver: Store
}

impl<Store: SessionStore> SessionManager<Store> {
    pub fn new(driver: Store) -> Self {
        SessionManager { driver }
    }
}

#[async_trait]
impl<D, Store> Middleware<D> for SessionManager<Store>
where
    D: Send + 'static + Sync,
    Store: SessionStore + Send + 'static + Sync
{
    fn invoke<'mw>(&'mw self,
                   req: &mut Request<'mw, '_, D>,
                   res: Response<'mw, D>
    ) -> MiddlewareResult<'mw, D> {
        res.next_middleware()
    }
}
