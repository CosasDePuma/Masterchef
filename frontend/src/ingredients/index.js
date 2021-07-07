import { Component } from "rete";

import builder from "./builder";
import worker from "./worker";
import { ingredients as specs } from "@/assets/ingredients";

// Variables
const categories = {};
const ingredients = [];

// Factory
function build(spec) {
  // check clones
  if (spec.clone) {
    const templates = specs.filter((s) => s.id === spec.clone);
    if (!templates.length) return null;
    spec = { ...templates[0], ...spec };
  }
  // create ingredient
  return class extends Component {
    constructor() {
      super(spec.name || "unknown");
      this.previous = {};
    }

    builder(node) {
      builder(node, spec, this.editor);
    }

    async worker(node, inputs, outputs) {
      // state
      const element = this.editor.nodes.find((n) => n.id === node.id);
      element.vueContext.$el
        .querySelector(".icons")
        .setAttribute("current", "busy");
      // checks & worker
      try {
        await worker(spec, node, inputs, outputs, element.controls);
        element.vueContext.$el
          .querySelector(".icons")
          .setAttribute("current", "ok");
      } catch (err) {
        console.error(`Error. Node: ${node.id}`, err);
        element.vueContext.$el
          .querySelector(".icons")
          .setAttribute("current", "err");
      }
    }
  };
}

// Iterators
specs.forEach(function(spec) {
  if (categories[spec.category] === undefined) categories[spec.category] = [];
  categories[spec.category].push(spec);
  const proto = build(spec);
  if (typeof proto === "function") ingredients.push(new proto());
});

export { categories, ingredients };
