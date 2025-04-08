import type { RealmListing } from "$lib/ts/types/Realm.svelte";
import type { SdkConfiguration, SdkError } from "$lib/ts/types/Sdk.svelte";

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