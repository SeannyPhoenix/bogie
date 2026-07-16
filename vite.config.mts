import { defineConfig } from "vite";

export default defineConfig({
  server: {
    open: "/static/index.html",
  },
  build: {
    rollupOptions: {
      input: "static/index.html",
    },
  },
});
