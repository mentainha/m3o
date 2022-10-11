const defaultTheme = require('tailwindcss/defaultTheme')

module.exports = {
  content: [
    './components/**/*.{jsx,tsx}',
    './lib/**/*.{jsx,tsx}',
    './pages/**/*.{jsx,tsx}',
    './providers/**/*.{jsx,tsx}',
  ],
  theme: {
    extend: {},
    fontFamily: {
      sans: ['Poppins', ...defaultTheme.fontFamily.sans],
      mono: [...defaultTheme.fontFamily.mono],
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
}
