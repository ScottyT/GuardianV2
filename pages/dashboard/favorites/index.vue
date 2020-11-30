<template>
  <div class="user-area user-area__favorites">
    <div class="filters-area">
      <div class="filters-area__sort">
        <label for="sorting">Sort By:</label>
        <div class="filters-area__sort-box" @click="selectActive = !selectActive" :class="{ open: selectActive }">
          <div class="filters-area__sort--trigger">
            <span>{{ selectedOption.text }}</span>
            <font-awesome-icon class="arrow" :icon="['fas', 'angle-down']" />
          </div>
          <div class="filters-area__sort--options">
            <span class="filters-area__sort--option" v-for="option in selectOptions" :key="option.value" @click="sortList(option)">
              {{option.text}}
            </span>
          </div>
        </div>
      </div>
      <div class="filters-area__filter">
        <label>Filter By:</label>
        <div class="filters-area__filter-box">
          <div class="filters-area__filter-box--group" v-for="fav in filterOptions" :key="fav.value">
            <label :for="fav.value">{{fav.text}}</label>
            <input type="checkbox" :id="fav.value" v-model="selectedOptions" class="filters-area__filter-box--checkbox" :value="fav.text" />
          </div>
        </div>
      </div>
    </div>
    <div class="favorite-projects">
      <div v-for="(item, i) in sortedFavs" :key="i">
        {{item.name}} - {{item.type}}
      </div>
    </div>
  </div>
</template>
<script>
import { createNamespacedHelpers } from "vuex";
const { mapState } = createNamespacedHelpers("project");
export default {
  layout: 'dashboard',
  async asyncData({ store, params }) {
    const projects = await store.state.project.favorites
    return {
      favs: projects
    }
  },
  head() {
    return {
      title: 'Favorites'
    }
  },
  async middleware({ store }) {
    const user = store.state.auth.user
    await store.dispatch("fetchUserFavs", user)
  },
  data() {
    return {
      selectOptions: [
        { value: 'nameAsc', text: 'Name, A-Z' },
        { value: 'nameDesc', text: 'Name, Z-A' },
        { value: 'typeAsc', text: 'Type, Ascending' },
        { value: 'typeDesc', text: 'Type, Decsending' }
      ],
      selectedOption: {
        value: 'nameAsc',
        text: 'Name, A-Z'
      },
      filterOptions: [
        { value: 'roofing', text: 'Roofing' },
        { value: 'gutters', text: 'Gutters' },
        { value: 'siding', text: 'Siding' },
        { value: 'flooring', text: 'Flooring' },
        { value: 'drywall', text: 'Drywall' },
        { value: 'painting', text: 'Painting' }
      ],
      selectedOptions:[],
      selectActive: false,
      filteredFavs: [],
      projectFavs: [],
      sorted: false
    }
  },
  computed: {
    ...mapState({
      favorites: (state) => state.favorites
    }),
    sortedFavs() {
      switch (this.sorted) {
        case false:
          return this.favorites
        case true:
          return this.filteredFavs
      }
    }
  },
  methods: {
    async getFavs() {
      await this.$store.dispatch("fetchUserFavs")
    },
    async sortList(option) {
      console.log(option)
      this.selectedOption = option
      
      
      await this.$axios.$post("/sorting", {
        option: option.value,
        project: this.favorites
      }).then((res) => {
        console.log(res)
        this.sorted = true
       // this.favorites = []
        //this.$store.dispatch("project/sortedFavs", res.projects, { root: true })
        this.filteredFavs = res.project
        this.favorites = newArr
      }).catch((err) => {
        console.log(err)
      })
    }
  },
  created() {
    //this.getFavs()
   // this.projectFavs = 
  }
}
</script>
<style lang="scss">
.filters-area {
  padding-bottom:40px;
  display:grid;
  grid-template-rows:100px;
  grid-template-columns:auto 1fr;
  column-gap:30px;
  &__sort-box {
    border:2px solid $color-black; 
    display:block;
    height:0;
    transition: height .3s ease-in;
    background-color:$color-white;
    z-index:2;
    position:relative;
    &.open {
      height:217px;
      transition: height .3s ease-in;
      .filters-area__sort--options {
        opacity:1;
        visibility: visible;
        cursor:pointer;
      }
    }
  }
  &__sort {
    width:185px;
    &--trigger {
      padding:5px;
      border-bottom:1px solid $color-black;
      cursor:pointer;
    }

    &--options {
      opacity:0;
      display:flex;
      flex-direction:column;
      padding:0 5px;
      transition:all .2s ease-in;
      &.open {
        opacity:1;
        visibility:visible;
      }
    }
  }
  &__filter {
    max-width:600px;
    width:100%;
  }
  &__filter-box {
    height:50px;
    display:flex;
    align-items:center;
    justify-content:space-between;
    border-top:1px solid $color-black;
    border-bottom:1px solid $color-black;
  }
}
</style>