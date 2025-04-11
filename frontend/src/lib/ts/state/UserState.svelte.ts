import type { User } from "../types/User.svelte";

export class UserState {
    private user: User | undefined = $state.raw()
    public LoggedIn = $derived(this.user!==undefined)

    Get(): User | undefined {
        return this.user;
    }

    Set(user: User) {
        this.user = user;
    }
}

