"use strict";

import Vue from "vue";
import Controls from "./Controls.vue";

function install(editor) {
  const wrapper = document.createElement("div");
  wrapper.id = "controls";
  editor.view.container.appendChild(wrapper);

  const component = Vue.extend(Controls);
  new component({ propsData: { editor } }).$mount("#controls");
}

export default { name: "vue-controls", install };
