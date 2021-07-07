<template>
  <main>
    <div class="section">Editor</div>
    <div id="editor" ref="editor" />
  </main>
</template>

<script>
import { createEditor } from "@/scripts/editor";
export default {
  name: "editor",
  mounted: function() {
    createEditor(this.$refs["editor"]);
  },
};
</script>

<style>
main {
  grid-area: main;
  width: 100%;
  height: 100%;
  border-left: 3px solid var(--decorator);
  border-right: 3px solid var(--decorator);
}
main > #editor {
  grid-area: editor;
  width: 100%;
  height: calc(100% - 50px) !important;
  background: var(--background);
  position: relative;
}
main > #editor svg.connection > path.main-path {
  stroke: var(--primary) !important;
  stroke-width: 2px !important;
}
</style>
