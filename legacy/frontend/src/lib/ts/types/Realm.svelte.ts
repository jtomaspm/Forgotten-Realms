export type RealmListing = {
    id: string 
    name: string 
    status: 'open' | 'closed' | 'ended'
    registered: boolean
    created_at: string 
}

export type Realm = {
    id: string 
    name: string 
    api: string 
    status: 'open' | 'closed' | 'ended'
    created_at: string 
    updated_at: string 
}

export type PlayableRealm = {
    id: string 
    name: string 
    api: string 
}