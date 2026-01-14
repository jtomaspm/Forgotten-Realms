import type { IComponent } from "./component";

export class Router {
    private _routes: Map<string, ()=>IComponent>

    constructor() {
        this._routes = new Map();
    }

    add_route(path: string, callback: ()=>IComponent) {
        this._routes.set(path, callback);
    }

    render(route: string) {
        const callback = this._routes.get(route);
        if (callback) callback().render();
    }
}