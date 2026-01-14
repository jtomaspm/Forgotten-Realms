import type { Coords, Village } from "../../gateway/game_api/types";
import { VirtualMap } from "./virtual_map";

export class GameMapRender {
    private _canvas: HTMLCanvasElement;
    private _ctx: CanvasRenderingContext2D;
    private _center: Coords;
    private readonly _SLOT_SIZE: number = 65;
    private readonly _COORD_INDICATOR_SIZE: number = 25;
    private _virtual_map: VirtualMap;
    constructor(
        canvas: HTMLCanvasElement, 
        ctx: CanvasRenderingContext2D,
        center: Coords,
    ) {
        this._canvas = canvas;
        this._ctx = ctx;
        this._center = center;
        this._virtual_map = new VirtualMap();
    }

    update_center(coords: Coords) {
        this._center = coords;
    }

    camera_world_position(): Coords {
        return {
            x: this._center.x * this._SLOT_SIZE,
            y: this._center.y * this._SLOT_SIZE,
        };
    }
    camera_screen_position(): Coords {
        return {
            x: this._canvas.width / 2,
            y: this._canvas.height / 2,
        };
    }

    coords_to_pixel(coords: Coords, camera_world_position: Coords, camera_screen_position: Coords): Coords {
        const world_pixels: Coords = {
            x: coords.x * this._SLOT_SIZE,
            y: coords.y * this._SLOT_SIZE,
        };
        const offset_to_camera: Coords = {
            x: world_pixels.x - camera_world_position.x,
            y: world_pixels.y - camera_world_position.y,
        }
        return {
            x: camera_screen_position.x + offset_to_camera.x,
            y: camera_screen_position.y + offset_to_camera.y,
        }
    }

    render_background() {
        this._ctx.fillStyle = "rgb(186, 134, 44)";
        this._ctx.fillRect(0, 0, this._canvas.width, this._canvas.height);
    }

    render_village(village: Village, camera_world_position: Coords, camera_screen_position: Coords) {
        const screen_position = this.coords_to_pixel(village.coords, camera_world_position, camera_screen_position);
        switch (village.player.faction) {
            case "caldari":
                this._ctx.fillStyle = "rgb(0, 55, 194)";
                break;
            case "dawnhold":
                this._ctx.fillStyle = "rgb(255, 242, 0)";
                break;
            case "varnak":
                this._ctx.fillStyle = "rgb(54, 0, 0)";
                break;
        }
        this._ctx.fillRect(screen_position.x-this._SLOT_SIZE/2, screen_position.y-this._SLOT_SIZE/2, this._SLOT_SIZE, this._SLOT_SIZE);
    }

    render_villages() {
        const sp = this.camera_screen_position();
        const wp = this.camera_world_position();
        const camara_position = this.coords_to_pixel(this._center, wp, sp);
        const offset: Coords = {
            x: Math.floor(camara_position.x / this._SLOT_SIZE),
            y: Math.floor(camara_position.y / this._SLOT_SIZE),
        };

        const min: Coords = {
            x: Math.floor(this._center.x-offset.x-2),
            y: Math.floor(this._center.y-offset.y-2),
        };
        const max: Coords = {
            x: Math.floor(this._center.x+offset.x+2),
            y: Math.floor(this._center.y+offset.y+2),
        };

        for(let x = min.x; x <= max.x; x++){
            for(let y = min.y; y <= max.y; y++){
                let village = this._virtual_map.get_village({x, y});
                if (village) this.render_village(village, wp, sp);
            }
        }
    }

    render_coord_indicators() {
        this._ctx.fillStyle = "rgb(255 255 255)";
        this._ctx.fillRect(0, 0, this._COORD_INDICATOR_SIZE, this._canvas.height);
        this._ctx.fillRect(0, this._canvas.height-this._COORD_INDICATOR_SIZE, this._canvas.width, this._COORD_INDICATOR_SIZE);
        this._ctx.moveTo(this._COORD_INDICATOR_SIZE, 0);
        this._ctx.lineTo(this._COORD_INDICATOR_SIZE, this._canvas.height);
        this._ctx.moveTo(0, this._canvas.height-this._COORD_INDICATOR_SIZE);
        this._ctx.lineTo(this._canvas.width, this._canvas.height-this._COORD_INDICATOR_SIZE);
        this._ctx.stroke();
    }

    render() {
        this.clear();
        this.render_background();
        this.render_villages();
        //this.render_coord_indicators();
    }
    
    clear() {
        this._ctx.clearRect(0, 0, this._canvas.width, this._canvas.height);
    }
}