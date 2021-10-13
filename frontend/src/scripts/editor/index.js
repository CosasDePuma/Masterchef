import { Engine, NodeEditor } from "rete";
import PluginConnection from "rete-connection-plugin";
import PluginEvents from "./plugins/events";
import PluginVue from "./plugins/vue-render";
import PluginControls from "./plugins/vue-controls";

import { ingredients } from "@/ingredients";

export async function createEditor(container) {
  // editor
  const editor = new NodeEditor(`misterchef@${Misterchef.version}`, container);
  // plugins
  editor.use(PluginConnection);
  editor.use(PluginVue);
  editor.use(PluginEvents);
  editor.use(PluginControls);
  // engine
  const engine = new Engine(`misterchef@${Misterchef.version}`);
  // ingredients
  ingredients.forEach(function(ingredient) {
    editor.register(ingredient);
    engine.register(ingredient);
  });
  // events
  editor.on("process", async function() {
    await engine.abort();
    editor.nodes = editor.nodes.map(function(node) {
      if (node.vueContext)
        node.vueContext.$el
          .querySelector(".icons")
          .setAttribute("current", "nothing");
      return node;
    });
    await engine.process(editor.toJSON());
  });
  // visualization
  editor.view.resize();
  editor.trigger("process");
}
