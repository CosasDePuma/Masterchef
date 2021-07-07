<template>
  <aside>
    <div id="ingredients">
      <div class="section">Ingredients</div>
      <div id="search">
        <input type="text" placeholder="Search..." v-model="search" />
      </div>
      <div class="scrolled">
        <ul class="ingredients">
          <ingredient
            v-for="ingredient in found"
            :key="ingredient.name"
            :spec="ingredient"
          />
        </ul>
        <category
          v-for="(ingredients, name) in categories"
          :key="name"
          :name="name"
          :ingredients="ingredients"
        />
      </div>
    </div>
  </aside>
</template>

<script>
import Category from "./Category.vue";
import Ingredient from "./Ingredient.vue";
import { categories } from "@/ingredients";

export default {
  name: "ingredients",
  components: { Category, Ingredient },
  data: function() {
    return { categories, search: "" };
  },
  computed: {
    found: function() {
      const ingredients = [];
      if (this.search !== "") {
        const term = this.search.toLowerCase();
        Object.keys(categories).forEach(function(category) {
          categories[category].forEach(function(ingredient) {
            const names = [ingredient.name, ...ingredient.aliases];
            if (
              names.filter(function(name) {
                return name.toLowerCase().includes(term);
              }).length > 0
            ) {
              ingredients.push(ingredient);
            }
          });
        });
      }
      return ingredients;
    },
  },
};
</script>

<style scoped>
aside {
  grid-area: ingredients;
  width: 100%;
  height: 100%;
  display: grid;
  grid-template-rows: 70px 1fr;
  grid-template-columns: 1fr;
  grid-template-areas:
    "banner"
    "ingredients";
  background: var(--background);
  overflow: hidden;
}
aside > #ingredients {
  width: 100%;
  height: 100%;
}
aside > #ingredients > #search {
  border-bottom: 2px solid var(--decorator);
}
aside > #ingredients > #search > input {
  padding: 15px;
  font-size: 14px;
  border-bottom: 1px solid var(--decorator);
}
aside > #ingredients > .scrolled {
  max-height: calc(100vh - 161px);
  overflow-y: auto;
}
</style>
