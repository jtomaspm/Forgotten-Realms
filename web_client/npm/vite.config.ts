import { defineConfig } from "vite";
import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
    build: {
        outDir: "dist",
        manifest: true,
        emptyOutDir: true,
        rollupOptions: {
            input: {
                home: "src/pages/index.ts",
                dashboard: "src/pages/test.ts"
            },
            output: {
                entryFileNames: "[name].js",
                assetFileNames: "[name].css"
            }
        }
    },
    plugins: [
        tailwindcss(),
    ]
});
