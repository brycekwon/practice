import defaultTheme from "tailwindcss/defaultTheme";
import { mocha } from "@catppuccin/tailwindcss";

/** @type {import('tailwindcss').Config} */
export default {
    darkMode: ["class"],
    content: [
        "./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}",
    ],
    theme: {
        extend: {
            colors: {
                ...mocha,
            },
            fontFamily: {
                sans: ["Inter", ...defaultTheme.fontFamily.sans],
                serif: ["Lora", ...defaultTheme.fontFamily.serif],
            },
        },
    },
    plugins: [
        require("@tailwindcss/typography"),
    ],
};
