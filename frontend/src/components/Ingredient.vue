<template>
  <li draggable="true" @dragstart="move" :title="spec.description">
    {{ spec.name }}
  </li>
</template>

<script>
export default {
  name: "ingredient",
  props: ["spec"],
  methods: {
    move: function(event) {
      if (!event || !event.dataTransfer) return;
      event.dataTransfer.allowedEffect = "move";
      event.dataTransfer.setData("spec", JSON.stringify(this.spec));
    },
  },
};
</script>

<style scoped>
li {
  height: 40px;
  display: flex;
  align-items: center;
  padding: 20px;
  user-select: none;
  border-bottom: 1px solid var(--decoratorprimary);
  color: var(--primarylight);
  background: var(--primarylighter);
  cursor: move;
}
li:hover {
  background-color: var(--decoratorprimary);
}
</style>
