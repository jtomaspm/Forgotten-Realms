import { defineConfig } from "vite";
import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
    build: {
        outDir: "dist",
        manifest: true,
        emptyOutDir: true,
        rollupOptions: {
            input: {
                index: "src/pages/index.ts",
                test: "src/pages/test.ts"
            },
            output: {
                entryFileNames: "js/[name].js",
                assetFileNames: "css/[name].css"
            }
        }
    },
    plugins: [
        tailwindcss(),
    ]
});
