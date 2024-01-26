import { defineConfig } from "vite";
import { run } from "vite-plugin-run";
import { viteStaticCopy } from "vite-plugin-static-copy";

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [
		run([
			{
				name: "build wasm",
				run: ["go", "generate", "frontend/main.go"],
				pattern: ["*.go", "./**/*.go"],
				startup: true,
			},
		]),
		viteStaticCopy({
			targets: [
				{ src: "app.js", dest: "../frontend_dist/" },
				{ src: "app-worker.js", dest: "../frontend_dist/" },
				{ src: "wasm_exec.js", dest: "../frontend_dist/" },
				{ src: "manifest.webmanifest", dest: "../frontend_dist/" },
				{ src: "web/app.wasm", dest: "../frontend_dist/web/" },
			],
		}),
	],
	build: {
		outDir: "../frontend_dist/",
		emptyOutDir: true,
	},
});
