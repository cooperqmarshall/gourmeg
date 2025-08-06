const colors = require('tailwindcss/colors')

module.exports = {
  content: ["./templates/*.html"],
  theme: {
    extend: {},
    colors: {
        "transparent": "transparent",
        "current": "currentColor",
        "white": colors.white,
        "indigo": colors.indigo,
        "gray": colors.gray,
        "turquoise": "#2dd4bf",
        "jordy-blue": "#93dcfd",
        "mauve": {100: "#d9c2f2", 200: "#d8b4fe", 300: "#c29de9"},
        "carnation-pink": "#fe99cf",
        "seasalt": "#fafafa",
    },
  },
  plugins: [],
}

