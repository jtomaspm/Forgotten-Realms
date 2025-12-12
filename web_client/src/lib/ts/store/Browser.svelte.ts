export function GetSessionToken(cookies?: { get: (name: string) => string | undefined }): string | null {
    if (cookies) {
        return cookies.get('popfrsid') || null;
    } else if (typeof document !== 'undefined') {
        const name = "popfrsid=";
        const decodedCookie = decodeURIComponent(document.cookie);
        const parts = decodedCookie.split(';');
        for (const part of parts) {
            const trimmedPart = part.trim();
            if (trimmedPart.startsWith(name)) {
                return trimmedPart.substring(name.length);
            }
        }
    }
    return null;
}

export function DeleteSessionToken(cookies?: { set: (name: string, value: string, options: any) => void }): void {
    if (cookies) {
        cookies.set('popfrsid', '', {
            path: '/',
            expires: new Date(0)
        });
    } else if (typeof document !== 'undefined') {
        document.cookie = "popfrsid=; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT";
    }
}

export function SetSessionToken(token: string, cookies?: { set: (name: string, value: string, options: any) => void }): void {
    const options = {
        path: '/',
        maxAge: 86400*7,
    };

    if (cookies) {
        cookies.set('popfrsid', token, options);
    } else if (typeof document !== 'undefined') {
        document.cookie = `popfrsid=${encodeURIComponent(token)}; path=${options.path}; max-age=${options.maxAge}`;
    }
}