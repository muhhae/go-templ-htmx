/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "**/*.templ"
  ],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["dim", "nord"],
  },
  plugins: [require("daisyui")],
}
