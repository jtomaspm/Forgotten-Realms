export type SdkConfiguration = {
    url: string
}

export type SdkError = {
    StatusCode: number
    Errors: string[]
} | undefined