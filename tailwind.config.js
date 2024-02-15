/** @type {import('tailwindcss').Config} */
module.exports = {
  important: true,
  content: ["views/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/forms")],
};
