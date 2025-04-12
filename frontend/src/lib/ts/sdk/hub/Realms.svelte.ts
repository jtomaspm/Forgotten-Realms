import type { PlayableRealm, Realm, RealmListing } from "$lib/ts/types/Realm.svelte";
import type { SdkConfiguration, SdkError } from "$lib/ts/types/Sdk.svelte";

export async function GetPlayableRealms(configuration: SdkConfiguration, token: string) : Promise<{ realms: PlayableRealm[], error: SdkError }> {
    try {
        let headers : any = {
            'Content-Type':  'application/json',
            'Authorization': 'Bearer ' + token
        };
        let response = await fetch(`${configuration.url}/api/realm/account`, {
            method: 'GET',
            headers: headers
        })
        if (!(response.status === 200)) {
            return {
                realms: [],
                error: {
                    StatusCode: response.status,
                    Errors: [await response.json()]
                }
            };
        }
        return {error: undefined, realms: (await response.json()).realms };
    }
    catch (ex)
    {
        console.log("Exception:", ex);
        return { realms: [], error: {
            StatusCode: 0,
            Errors: [(ex as Error).message]
        }};
    }
}

export async function GetRealms(configuration: SdkConfiguration, token: string | undefined) : Promise<{ realms: RealmListing[], error: SdkError }> {
    try {
        let headers : any = {
            'Content-Type': 'application/json',
        };
        if (token) {
            headers.Authorization = "Bearer " + token
        }
        let response = await fetch(`${configuration.url}/api/realm`, {
            method: 'GET',
            headers: headers
        })
        if (!(response.status === 200)) {
            return {
                realms: [],
                error: {
                    StatusCode: response.status,
                    Errors: [await response.json()]
                }
            };
        }
        return {error: undefined, realms: (await response.json()).realms };
    }
    catch (ex)
    {
        return { realms: [], error: {
            StatusCode: 0,
            Errors: [(ex as Error).message]
        }};
    }
}

export async function GetRealm(configuration: SdkConfiguration, realmId: string) : Promise<{ realm: Realm | null, error: SdkError }> {
    try {
        let headers : any = {
            'Content-Type': 'application/json',
        };
        let response = await fetch(`${configuration.url}/api/realm/${realmId}`, {
            method: 'GET',
            headers: headers
        })
        if (!(response.status === 200)) {
            return {
                realm: null,
                error: {
                    StatusCode: response.status,
                    Errors: [await response.json()]
                }
            };
        }
        return {error: undefined, realm: (await response.json()) };
    }
    catch (ex)
    {
        return { realm: null, error: {
            StatusCode: 0,
            Errors: [(ex as Error).message]
        }};
    }
}