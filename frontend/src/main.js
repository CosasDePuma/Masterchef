import "regenerator-runtime/runtime";
import Vue from "vue";

import App from "./App.vue";
import { version } from "../package.json";

global.Masterchef = { version };

Vue.config.productionTip = false;
new Vue({
  render: (h) => h(App),
}).$mount("#root");
