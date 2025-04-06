export type RealmListing = {
    id: string 
    name: string 
    status: 'open' | 'closed' | 'ended'
    registered: boolean
    created_at: string 
}
