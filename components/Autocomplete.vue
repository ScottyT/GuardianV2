<template>
  <div class="autocomplete" v-click-outside="onClickOutside">
    <input type="text" @input="onChange" v-model="search" :class="cssClass" @keydown.tab="setResult(results[arrowCounter])" :placeholder="placeholderText" @keydown.down.prevent="onArrowDown" @keydown.up.prevent="onArrowUp" />
    <ul ref="scrollContainer" v-show="isOpen" class="autocomplete__results">
      <li class="loading" v-if="isLoading">Loading results...</li>
      <li ref="options" v-else @click="setResult(result)" class="autocomplete__result" :class="{ 'is-active' : i === arrowCounter }" v-for="(result, i) in results" :key="i">
        {{result}}
      </li>
    </ul>
  </div>
</template>
<script>
export default {
  name: "Autocomplete",
  props: {
    cssClass: {
      type: String,
      required:false
    },
    placeholderText: {
      type: String,
      required:false
    },
    items: {
      type: Array,
      required: false,
      default: () => []
    },
    isAsync: {
      type: Boolean,
      required:false,
      default:false
    },
    search: {
      type: String,
      required: false
    }
  },
  data() {
    return {
      results: [],
      isOpen: false,
      isLoading: false,
      arrowCounter: -1
    }
  },
  watch: {
    items(value, oldValue) {
      if (value.length !== oldValue.length) {
        this.results = value
        this.isOpen = true
        this.isLoading = false
      }
    }
  },
  methods: {
    onClickOutside() {
      this.isOpen = false
      this.arrowCounter = -1
    },
    onChange() {
      this.$emit('input', this.search)
      if (this.isAsync) {
        this.isLoading = true
      } else {
        this.filterResults()
        this.isOpen = true
      }
    },
    setResult(result) {
      this.search = result
      this.isOpen = false
      this.arrowCounter = -1
      this.$emit('set_result', result)
    },
    filterResults() {
      this.results = this.items.filter(item => item.toLowerCase().indexOf(this.search.toLowerCase()) > -1)
    },
    onArrowDown() {
      if (this.arrowCounter < this.results.length - 1) {
        this.arrowCounter = this.arrowCounter + 1
        this.fixScrolling()
      }
    },
    onArrowUp() {
      if (this.arrowCounter > 0) {
        this.arrowCounter = this.arrowCounter - 1
        this.fixScrolling()
      }
    },
    // onEnter() {
    //   this.search = this.results[this.arrowCounter]
    //   this.isOpen = false
      
    // },
    fixScrolling() {
      const itemHeight = this.$refs.options[this.arrowCounter].clientHeight
      this.$refs.scrollContainer.scrollTop = itemHeight * (this.arrowCounter - 1)
    },
    mounted() {
      document.addEventListener("click", this.onClickOutside)
    },
    destroyed() {
      document.removeEventListener("click", this.onClickOutside)
    }
  }
}
</script>
<style lang="scss">
.autocomplete {
  position:relative;
  width:130px;
  &__results {
    z-index:4;
    position:relative;
    width:500px;
    background-color:$color-white;
    padding:0!important;
    border:1px solid $color-black;
    height:150px;
    overflow:auto;
  }
  &__result {
    padding:4px 20px;
    cursor:pointer;
    &:hover {
      background-color:$primary-dark;
      color:$color-white;
    }
    &.is-active {
      background-color:$primary-dark;
      color:$color-white;
    }
  }
}
</style>