import { Router } from "../framework/router";
import { MapPage } from "./map/page";

export function GetGameRouter(screen: HTMLElement): Router {
    const router = new Router();
    router.add_route("map", () => new MapPage(screen));
    return router;
}