const REMOTE_URL = "http://localhost:4000"

const AUTH_URL = `${REMOTE_URL}/auth`

const LOGIN_URL = `${AUTH_URL}/login`
const LOGOUT_URL = `${AUTH_URL}/logout`
const WHOAMI_URL = `${AUTH_URL}/whoami`

export const parseResponse = async (response) => {
    if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`)
    return response.json()
}

export const login = async (username, password) => {
    return parseResponse(await fetch(LOGIN_URL, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({username, password})
    }))
}
