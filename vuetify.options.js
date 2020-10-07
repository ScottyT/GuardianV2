export default {
  theme: {
    customProperties: true,
    minifyTheme: function (css) {
      return process.env.NODE_ENV === "production" ?
        css.replace(/[\r\n|\r|\n]/g, "") :
        css;
    },
    disable: true,
    themes: {
      light: {
        primary: {
          base: "#bf493c",
          anchor: "#000"
        },
        secondary: {
          base: "#3cb2bf",
          darken: "#2a6564"
        }
      },
      dark: {
        primary: {
          anchor: "#fff"
        }
      }
    },
  }
}
