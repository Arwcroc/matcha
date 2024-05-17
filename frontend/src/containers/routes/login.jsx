import React from 'react'
import {Box, Button, Checkbox, FormControlLabel, TextField, Typography} from "@mui/material";
import MatchaPage from "../layout/MatchaPage";
import {login} from "../../helper/requests";

const Login = () => {
    const handleSubmit = e => {
        e.preventDefault();
        const data = new FormData(e.currentTarget);
        login(data.get('username'), data.get('password')).catch((error) => {
            console.error(error);
        }).then((response) => {
            console.log(response)
        })
    }

    return (
        <Box>
            <Typography component="h1" variant="h5">
                Sign in
            </Typography>
            <Box component="form" onSubmit={handleSubmit} noValidate>
                <TextField
                    required
                    fullWidth
                    autoFocus
                    id="username"
                    name="username"
                    label="Username"
                    type="text"
                    autoComplete="username"
                />
                <TextField
                    required
                    fullWidth
                    id="password"
                    name="password"
                    label="Password"
                    type="password"
                    autoComplete="current-password"
                />
                <FormControlLabel
                    control={<Checkbox value="remember" color="primary" />}
                    label="Remember me"
                />
                <Button
                    fullWidth
                    type="submit"
                    variant="contained"
                >
                    Sign In
                </Button>
            </Box>
        </Box>
    )
}

export default MatchaPage(Login)