import type { Faction, Player } from "$lib/ts/types/Player.svelte";
import type { SdkConfiguration, SdkError } from "$lib/ts/types/Sdk.svelte";

export async function RegisterPlayer(configuration: SdkConfiguration, token: string, faction: Faction) : Promise<{error: SdkError}> {
    try {
        let headers : any = {
            'Content-Type': 'application/json',
        };
        headers.Authorization = "Bearer " + token
        let response = await fetch(`${configuration.url}/api/player`, {
            method: 'POST',
            headers: headers,
            body: JSON.stringify({
                faction: faction
            })
        })
        if (!(response.status === 201)) {
            return {
                error: {
                    StatusCode: response.status,
                    Errors: [await response.json()]
                }
            };
        }
        return { error: undefined };
    }
    catch (ex)
    {
        return { error: {
            StatusCode: 0,
            Errors: [(ex as Error).message]
        }};
    }
}

export async function GetPlayer(configuration: SdkConfiguration, playerId: string) : Promise<{ player: Player | null, error: SdkError }> {
    try {
        let headers : any = {
            'Content-Type': 'application/json',
        };
        let response = await fetch(`${configuration.url}/api/player/${playerId}`, {
            method: 'GET',
            headers: headers
        })
        if (!(response.status === 200)) {
            return {
                player: null,
                error: {
                    StatusCode: response.status,
                    Errors: [await response.json()]
                }
            };
        }
        return {error: undefined, player: (await response.json()) };
    }
    catch (ex)
    {
        return { player: null, error: {
            StatusCode: 0,
            Errors: [(ex as Error).message]
        }};
    }
}