import React from 'react';
import {Container, CssBaseline} from "@mui/material";

const MatchaPage = (WrappedComponent) => {
    return (props) => {
        return (
            <Container maxWidth="sm">
                <CssBaseline />
                <WrappedComponent {...props} />
            </Container>
        )
    }
}

export default MatchaPage
