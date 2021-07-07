<template>
  <div class="theme chef">
    <div id="app">
      <Navbar />
      <Ingredients />
      <Editor />
    </div>
  </div>
</template>

<script>
import "@/scripts/icons";
import Navbar from "@/components/Navbar.vue";
import Ingredients from "@/components/Ingredients.vue";
import Editor from "@/components/Editor.vue";

export default {
  name: "App",
  components: { Navbar, Ingredients, Editor },
};
</script>

<style>
*,
*::before,
*::after {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

input,
select,
textarea {
  width: 100%;
  height: 100%;
  border: none;
  outline: none;
  padding: 5px 10px;
  background: var(--background);
}

.chef {
  --primary: #1976d2;
  --primarylight: #3a87ad;
  --primarylighter: #d9edf7;
  --decorator: #eeeeee;
  --decoratorprimary: #d5e9f3;
  --decoratordark: #5f5f5f;
  --background: #ffffff;
  --backgrounddark: #fafafa;
}

#app {
  width: 100%;
  height: 100vh;
  display: grid;
  grid-template-rows: 60px 1fr;
  grid-template-columns: 230px 1fr 230px;
  grid-template-areas:
    "navbar navbar navbar"
    "ingredients main main";
  background: var(--background);
  font-size: 16px;
  font-family: ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont,
    "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif,
    "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
}

#app .section {
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  font-weight: 700;
  user-select: none;
  text-align: center;
  height: 50px;
  color: var(--decoratordark);
  background: var(--backgrounddark);
  border-bottom: 3px solid var(--decorator);
}
</style>
