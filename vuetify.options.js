export default {
  theme: {
    customProperties: true,
    minifyTheme: function (css) {
      return process.env.NODE_ENV === "production" ?
        css.replace(/[\r\n|\r|\n]/g, "") :
        css;
    },
    disable: false,
    themes: {
      light: {
        primary: {
          base: "#a13d2e",
        },
        anchor: "#fff",
        secondary: {
          base: "#777777"
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
