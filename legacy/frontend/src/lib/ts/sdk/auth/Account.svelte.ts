import type { SdkConfiguration, SdkError } from "$lib/ts/types/Sdk.svelte";
import type { User } from "$lib/ts/types/User.svelte";

export async function VerifyEmail(configuration: SdkConfiguration, token: string): Promise<{error: SdkError}> {
    try {
        const response = await fetch(`${configuration.url}/api/account/verify?token=${token}`, {
            method: 'GET',
        });
        if (response.status !== 202) {
            return { error: {
                StatusCode: response.status,
                Errors: [await response.json()]
            }}
        }
        return {error: undefined};
    }
    catch (ex)
    {
        return { error: {
            StatusCode: 0,
            Errors: [(ex as Error).message]
        }};
    }
}

export async function CreateAccount(
    configuration: SdkConfiguration, 
    token: string, 
    request: { name: string, send_email_notifications: boolean }
) : Promise<{ token : string , error: SdkError }> {
    try {
        let response = await fetch(`${configuration.url}/api/account`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': "Bearer " +  token
            },
            body: JSON.stringify(request)
        });

        if (response.status !== 201) {
            return { token: "", error: {
                StatusCode: 0,
                Errors: [(await response.json()).error]
            }}
        }
        return {
            token: (await response.json()).token,
            error: undefined
        }
    }
    catch (ex)
    {
        return { token: "", error: {
            StatusCode: 0,
            Errors: [(ex as Error).message]
        }}
    }
}

export async function GetAccount(
    configuration: SdkConfiguration, 
    token: string, 
) : Promise<{ account : User | null , error: SdkError }> {
    try {
        const response = await fetch(`${configuration.url}/api/account?token=${token}`, {
            method: 'GET',
        });

        if (response.status !== 200) {
            return { account: null, error: {
                StatusCode: 0,
                Errors: [(await response.json()).error]
            }}
        }
        var data = await response.json();
        return {
            account: {
                Id: data.id,
                Email: data.email,
                Name: data.name,
                Role: data.role,
                Token: token
            },
            error: undefined
        }
    }
    catch (ex)
    {
        return { account: null, error: {
            StatusCode: 0,
            Errors: [(ex as Error).message]
        }}
    }
}