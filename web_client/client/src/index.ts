import { CreateElement } from "./framework/elements";
import "./global_styles.css";
import { GetGameRouter } from "./pages/router";

const approot = document.getElementById("approot") as HTMLDivElement;

const top_section = CreateElement(
    approot,
    "div",
    [],
    () => {}
) as HTMLDivElement;

const game_screen = CreateElement(
    approot,
    "div",
    [],
    (element) => {
        const router = GetGameRouter(element as HTMLDivElement);
        router.render("map")
    }
) as HTMLDivElement;

const bot_section = CreateElement(
    approot,
    "div",
    [],
    () => {}
) as HTMLDivElement;
