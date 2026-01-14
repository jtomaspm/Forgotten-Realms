import type { IComponent } from "../../framework/component";
import { CreateElement } from "../../framework/elements";
import { GameMapRender } from "./render";

export class MapPage implements IComponent {
    _root: HTMLElement;
    constructor(root: HTMLElement) {
        this._root = root;
    }
    root(): HTMLElement {
        return this._root
    }
    render(): void {
        const canvas = CreateElement(
            this._root,
            "canvas",
            ["border-8", "m-8"],
            () => {}
        ) as HTMLCanvasElement;
        canvas.width = 1000;
        canvas.height = 1000;
        const ctx = canvas.getContext("2d")!;

        const map = new GameMapRender(canvas, ctx, {x: 500, y: 500});
        let i = 0;
        setInterval(() => {
            const seed = Math.random() - 0.2 * 10
            const seed2 = Math.random() - 0.5 * 10
            map.update_center({x: 500 + (i % 100)/30 * seed, y: 500 + (i % 100)/30*seed2});
            map.render();
            i++;
        }, 20);
        
        map.render();
    }
}