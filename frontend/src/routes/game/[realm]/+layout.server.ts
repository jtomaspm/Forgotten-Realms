import { GetPlayableRealms } from '$lib/ts/sdk/hub/Realms.svelte';
import { GetSessionToken } from '$lib/ts/store/Browser.svelte';
import { redirect } from '@sveltejs/kit';

export async function load({ params, cookies }) {
    const realmName = params.realm;
    let token = GetSessionToken(cookies);
    if (!token) { exit(); }
    let r = await GetPlayableRealms({ url: "http://hub:8081" }, token!);
    if (r.error || !r.realms) { exit(); }
    const current_realm = r.realms.find(x => x.name === realmName);
    if (!current_realm) { exit(); }
    return {
        realm: current_realm!,
        realms: r.realms
    };
}

function exit() {
    throw redirect(302, "/");
}