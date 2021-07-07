<template>
  <select @change="onChange($event)" :value="value" @pointerdown.stop>
    <optgroup :label="group">
      <option
        v-for="(value, i) in values"
        :key="value"
        :selected="index === i"
        :value="value.toLowerCase()"
      >
        {{ value }}
      </option>
    </optgroup>
  </select>
</template>

<script>
import { Control } from "rete";

const component = {
  props: [
    "emitter",
    "ikey",
    "getData",
    "putData",
    "disabled",
    "values",
    "initial",
  ],
  data: function() {
    return { value: "" };
  },
  computed: {
    index: function() {
      if(!this.value.length) this.value = this.values[this.initial].toLowerCase();
      return this.value ? this.values.indexOf(this.value) : this.initial;
    },
    group: function() {
      return (
        this.ikey.charAt(0).toUpperCase() + this.ikey.toLowerCase().slice(1)
      );
    },
  },
  mounted() {
    const data = this.getData(this.ikey);
    this.value = data ? data : this.values[this.initial].toLowerCase();
    this.update();
  },
  methods: {
    parse(value) {
      if (value === undefined) value = "";
      return value;
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
      initial: 0,
      values: [],
      disabled: false,
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
select {
  text-align: center;
  border: 1px solid var(--primary);
  border-radius: 2px;
  color: var(--primarylight);
}
select[disabled] {
  background-color: var(--primarylighter);
}
</style>
