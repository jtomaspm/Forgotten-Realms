export type AuthRegistrationCallbackResponse = {
    email: string
    expires_at: string 
    message: string 
    token: string 
}

export type AuthLoginCallbackResponse = {
    expires_at: string 
    token: string 
}

export type AccountCreated = {
    created: boolean
    token: string 
    name: string
}