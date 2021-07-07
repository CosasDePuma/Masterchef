<template>
  <div class="category">
    <div class="name" @click="hidden = !hidden">{{ name }}</div>
    <ul class="ingredients" v-show="!hidden">
      <ingredient
        v-for="ingredient in ingredients"
        :key="ingredient.name"
        :spec="ingredient"
      />
    </ul>
  </div>
</template>

<script>
import Ingredient from "./Ingredient.vue";

export default {
  name: "category",
  components: { Ingredient },
  props: ["name", "ingredients"],
  data: function() {
    return {
      hidden: true,
    };
  },
  methods: {
    getIngredient: function(event, spec) {
      // TODO: Borrar
      if (!event || !event.dataTransfer) return;
      event.dataTransfer.effectAllowed = "move";
      event.dataTransfer.setData("spec", JSON.stringify(spec));
    },
  },
};
</script>

<style scoped>
.name {
  height: 45px;
  display: flex;
  align-items: center;
  padding: 20px;
  user-select: none;
  border-bottom: 3px solid var(--decorator);
  text-transform: capitalize;
  color: var(--primary);
  font-weight: 700;
  background: var(--backgrounddark);
  cursor: pointer;
}
</style>
