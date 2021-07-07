<template>
  <input
    :type="type || 'text'"
    :value="value"
    :readonly="readonly"
    :placeholder="placeholder"
    :pattern="regex"
    spellcheck="false"
    autofocus="false"
    autocomplete="false"
    @input="onChange($event)"
    @dblclick.stop
    @pointerdown.stop
  />
</template>

<script>
import { Control } from "rete";
import { validations } from "@/scripts/validator";

const component = {
  props: [
    "emitter",
    "ikey",
    "getData",
    "putData",
    "type",
    "placeholder",
    "pattern",
    "readonly",
  ],
  data: function() {
    return { value: "" };
  },
  computed: {
    regex: function() {
      return this.pattern
        ? this.pattern
        : validations[this.type]
        ? validations[this.type]
        : ".*";
    },
  },
  mounted() {
    const value = this.getData(this.ikey);
    this.value =
      this.type === "number"
        ? value === undefined
          ? 0
          : +value
        : value === undefined
        ? ""
        : value;
    this.update();
  },
  methods: {
    parse(value) {
      if (value === undefined) value = "";
      switch (this.type) {
        case "number":
          value = +value;
          break;
        default:
          break;
      }
      return value;
    },
    onChange(event) {
      this.value = event.target.value;
      this.update();
    },
    update() {
      if (this.ikey) {
        this.putData(this.ikey, this.parse(this.value));
      }
    },
  },
};

export class Ctrl extends Control {
  constructor(
    emitter,
    ikey,
    options = {
      type: "text",
      initial: "",
      readonly: false,
      disabled: false, // TODO: Borrar
      placeholder: "",
      pattern: ".*",
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
input {
  border: 1px solid var(--primary);
  border-radius: 2px;
  text-align: center;
  color: var(--primarylight);
}
input:read-only {
  background-color: var(--primarylighter);
}
</style>
