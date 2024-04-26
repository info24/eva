export const Authorization = "Authorization"
export const setToken = (token: string) => {
    sessionStorage.setItem(Authorization, token)
}

export const getToken = () => {
    return sessionStorage.getItem(Authorization)
}
