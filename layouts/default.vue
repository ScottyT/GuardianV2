<template>
  <v-app dark>
    <app-navigation />
    <main class="v-content__wrap page" :class="this.$route.name == undefined ? 'error' : ''">
      <nuxt />
    </main>
    <footer-nav />
    <login-modal />
  </v-app>
</template>

<script>
  export default {
    data() {
      return {}
    },
    head() {
      const capitalizedName = []
      if (Object.entries(this.$route.params).length > 0) {
        var normalizedName

        normalizedName = this.$route.params.uid
          .replace(/\W+/g, " ")
          .trim()
          .split(" ")

        normalizedName.map((word, i) => {
          word = word.charAt(0).toUpperCase() + word.slice(1)
          capitalizedName.push(word)
        })
      }
      return {
        title: capitalizedName.join(" "),
      }
    },
    async middleware({
      store,
      $prismic
    }) {
      await store.dispatch("fetchMenu", $prismic)
      await store.dispatch("project/fetchProjects", null, {
        root: true
      })
    },
    methods: {
      // ...mapActions(['hideAuthModal'])
    }
  }
</script>
<style lang="scss">
  html {
    @include respond(mobileSmall) {
      font-size: 56.5% !important;
    }

    @include respond(mobileLarge) {
      font-size: 58.5% !important;
    }

    @include respond(tabletLarge) {
      font-size: 62.5% !important;
    }
  }

  body {
    font-size: 18px;
  }

  .v-application--wrap {
    overflow: hidden;
  }

  h1 {
    font-family: $heading-font;
    letter-spacing: 1px;

    @include respond(desktopSmall) {
      font-size: 4em;
    }

    &.extra-large {
      font-size: 2.5em;

      @include respond(mobileLarge) {
        font-size: 3.5em;
      }

      @include respond(tabletLarge) {
        font-size: 5em;
      }

      @include respond(desktopSmall) {
        font-size: 6em;
      }
    }

    &.red-text {
      color: $primary-dark;
    }

    &.small {
      font-size: 2em;

      @include respond(tabletLarge) {
        font-size: 3em;
      }
    }
  }

  h2 {
    font-size: 1.7em;

    @include respond(mobileLarge) {
      font-size: 2em;
    }

    @include respond(tabletLarge) {
      font-size: 2.5em;
    }

    @include respond(desktopSmall) {
      font-size: 3em;
    }
  }

  h3 {
    font-size: 1.2em;

    @include respond(mobileLarge) {
      font-size: 1.5em;
    }

    @include respond(tabletLarge) {
      font-size: 2.027em;
    }
  }

  h4 {
    font-size: 1.424em;
  }

  ul {
    list-style: none;
  }

  input {
    &:focus {
      outline: none !important;
    }
  }

  textarea {
    &:focus {
      outline: none !important;
    }
  }

  button {
    outline: none !important;
  }

  .error {
    padding-top: 100px;
	}
</style>