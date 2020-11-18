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
    </div>
    <div class="favorite-projects">
      <div v-for="(item, i) in favorites" :key="i">
        {{item.name}}
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
    //const user = await store.getters['auth/getUser']
    const projects = await store.state.project.favorites
    return {
      data: projects
    }
  },
  head() {
    return {
      title: 'Favorites'
    }
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
      selectActive: false
    }
  },
  computed: {
    ...mapState({
      favorites: (state) => state.favorites
    })
  },
  methods: {
    sortList(option) {
      console.log(option)
    }
  }
}
</script>
<style lang="scss">
.filters-area {
  padding-bottom:40px;
  &__sort-box {
    border:2px solid $color-black; 
    display:block;
    height:100%;
    &.open {
      .filters-area__sort--options {
        opacity:1;
        visibility: visible;
      }
    }
  }
  &__sort {
    width:130px;
    &--trigger {
      padding:5px;
    }

    &--options {
      opacity:0;
      &.open {
        opacity:1;
        visibility:visible;
      }
    }
  }
}
</style>