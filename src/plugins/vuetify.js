import Vue from "vue";
import Vuetify from "vuetify/lib";

Vue.use(Vuetify);

export default new Vuetify({
    theme: {
        options: {
            minifyTheme: function (css) {
                return process.env.NODE_ENV === "production" ?
                    css.replace(/[\r\n|\r|\n]/g, "") :
                    css;
            },
            customProperties: true
        },
        themes: {
            light: {
                primary: {
                    base: "#bf493c"
                },
                secondary: {
                    base: "#3cb2bf",
                    darken: "#2a6564"
                }
            }
        }
    },
    breakpoint: {
        thresholds: {
            xs: 415,
            sm: 768,
            md: 960,
            lg: 1200
        }
    }
});