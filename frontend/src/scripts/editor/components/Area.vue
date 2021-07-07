<template>
  <textarea
    :value="value"
    :readonly="readonly"
    :placeholder="placeholder"
    spellcheck="false"
    autofocus="false"
    rows="1"
    autocomplete="false"
    @dblclick="resize"
    @input="onChange($event)"
    @pointerdown.stop
  />
</template>

<script>
import { Control } from "rete";

const component = {
  props: ["emitter", "ikey", "getData", "putData", "readonly", "placeholder"],
  data: function() {
    return { value: "" };
  },
  mounted() {
    this.value = this.getData(this.ikey) || "";
    this.update();
  },
  methods: {
    parse(value) {
      if (value === undefined) value = "";
      return value;
    },
    resize() {
      this.$el.setAttribute("style", `height: auto`);
      this.$el.setAttribute(
        "style",
        `height: ${this.$el.scrollHeight < 460 ? this.$el.scrollHeight : 460}px`
      );
    },
    onChange(event) {
      this.value = event.target.value;
      this.update();
    },
    update() {
      if (this.ikey) {
        this.putData(this.ikey, this.value);
      }
    },
  },
};

export class Ctrl extends Control {
  constructor(
    emitter,
    ikey,
    options = {
      readonly: false,
      placeholder: "",
    }
  ) {
    super(ikey);
    this.data.render = "vue";
    this.props = {
      emitter,
      ikey,
      ...options,
    };
    this.component = component;
  }

  setValue(value) {
    const ctx = this.vueContext || this.props;
    ctx.value = value;
    this.update();
  }
}

export default component;
</script>

<style scoped>
textarea {
  border: 1px solid var(--primary);
  border-radius: 2px;
  color: var(--primarylight);
  height: auto;
}
textarea:read-only {
  background-color: var(--primarylighter);
}
</style>
