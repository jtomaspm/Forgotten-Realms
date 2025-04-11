export function GetSessionToken() {
    return localStorage.getItem("popfrsid")
}

export function DeleteSessionToken() {
    return localStorage.removeItem("popfrsid")
}

export function SetSessionToken(token: string) {
    localStorage.setItem("popfrsid", token);
}