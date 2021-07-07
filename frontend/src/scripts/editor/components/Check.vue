<template>
  <div class="checkbox">
    <div class="slider">
      <input
        :id="ikey.concat('-chk')"
        type="checkbox"
        value="null"
        :checked="value"
        @change="onChange"
      />
      <label :for="ikey.concat('-chk')"></label>
    </div>
    <span>{{ name }}</span>
  </div>
</template>

<script>
import { Control } from "rete";

const component = {
  props: ["emitter", "ikey", "getData", "putData", "label", "disabled"],
  data: function() {
    return { value: "" };
  },
  computed: {
    name: function() {
      return this.label && this.label.length ? this.label : this.ikey;
    },
  },
  mounted() {
    this.value = this.getData(this.ikey) || false;
    this.update();
  },
  methods: {
    onChange(event) {
      this.value = event.target.checked;
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
      label: "",
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
.checkbox {
  display: flex;
  align-items: center;
  justify-content: flex-start;
}
.slider {
  width: 12px;
  height: 12px;
  position: relative;
  margin: 5px;
  margin-right: 5px;
  background: transparent;
  border-radius: 2px;
}
.slider > label {
  width: 10px;
  height: 10px;
  position: absolute;
  top: 0px;
  left: 0px;
  cursor: pointer;
  background: var(--background);
  border: 1px solid var(--primary);
  border-radius: 2px;
}
.slider > label::after {
  content: "";
  width: 8px;
  height: 8px;
  position: absolute;
  top: 0px;
  left: 0px;
  background: var(--primarylight);
  border-radius: 1px;
  opacity: 0;
}
.slider > input[type="checkbox"] {
  visibility: hidden;
}
.slider > input[type="checkbox"]:checked + label:after {
  opacity: 1;
}
span {
  color: var(--primary);
  font-size: 14px;
  padding-bottom: 2px;
}
</style>
