<template>
  <div class="node" :class="classes">
    <div class="title">
      <span>{{ node.name }}</span>
      <div class="icons" :current="node.status">
        <icon-loading class="loading" :size="14" />
        <icon-ok class="ok" :size="14" />
        <icon-err class="err" :size="14" />
      </div>
    </div>

    <div class="sockets">
      <div class="input">
        <div v-for="input in inputs()" :key="input.key">
          <socket v-socket:input="input" type="input" :socket="input.socket" />
          <div class="title" v-show="!input.showControl()">
            <span>{{ input.name }}</span>
          </div>
          <div
            class="input-control"
            v-show="input.showControl()"
            v-control="input.control"
          />
        </div>
      </div>

      <div class="output">
        <div v-for="output in outputs()" :key="output.key">
          <div class="title">
            <span>{{ output.name }}</span>
          </div>
          <socket
            v-socket:output="output"
            type="output"
            :socket="output.socket"
          />
        </div>
      </div>
    </div>

    <div class="controls">
      <div
        class="control"
        v-for="control in controls()"
        :key="control.key"
        v-control="control"
      />
    </div>
  </div>
</template>

<script>
import filter from "./filter";
import Socket from "./Socket.vue";

export default {
  components: { Socket },
  props: ["node", "editor", "bindSocket", "bindControl"],
  computed: {
    classes: function() {
      return filter([this.selected(), this.node.name]);
    },
  },
  methods: {
    inputs() {
      return Array.from(this.node.inputs.values());
    },
    outputs() {
      return Array.from(this.node.outputs.values());
    },
    controls() {
      return Array.from(this.node.controls.values());
    },
    selected() {
      return this.editor.selected.contains(this.node) ? "selected" : "";
    },
  },
  directives: {
    socket: {
      bind(el, binding, vnode) {
        vnode.context.bindSocket(el, binding.arg, binding.value);
      },
      update(el, binding, vnode) {
        vnode.context.bindSocket(el, binding.arg, binding.value);
      },
    },
    control: {
      bind(el, binding, vnode) {
        if (!binding.value) return;

        vnode.context.bindControl(el, binding.value);
      },
    },
  },
};
</script>

<style scoped>
@keyframes rotate {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(359deg);
  }
}

.node {
  position: relative;
  min-width: 261px;
  height: auto;
  padding-bottom: 6px;
  background: var(--background);
  border: 1px solid var(--primarylight);
  border-radius: 3px;
  user-select: none;
}
.node:hover,
.node.selected {
  box-shadow: 0 2px 6px -1px rgba(0, 0, 0, 0.1),
    0 2px 4px -1px rgba(0, 0, 0, 0.06);
}
.node > .title {
  padding: 8px;
  color: var(--background);
  background: var(--primary);
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.node > .title > .icons > span {
  display: none;
}
.node > .title > .icons[current="ok"] > .ok,
.node > .title > .icons[current="err"] > .err,
.node > .title > .icons[current="busy"] > .loading {
  display: block;
}
.node > .title > .icons[current="busy"] > .loading {
  animation: rotate 2s infinite ease-in-out;
}

.node > .sockets {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 4px;
  color: var(--primary);
}
.node > .sockets > .output {
  margin-left: 15px;
  text-align: right;
}
.node > .sockets > .input {
  text-align: left;
}
.node > .sockets > .output > div,
.node > .sockets > .input > div {
  display: flex;
  align-items: center;
  overflow: hidden;
}
.node > .sockets > .output > div {
  justify-content: flex-end;
}
.node > .sockets > .output > .title,
.node > .sockets > .input > .title {
  display: inline-block;
  vertical-align: middle;
  margin: 6px;
  font-size: 14px;
  line-height: 24px;
  color: var(--primary);
}
.node > .sockets > .input .input-control {
  display: inline-block;
  vertical-align: middle;
  z-index: 1;
  width: 120px;
  height: 20px;
}
.node > .controls > .control {
  padding: 1px 18px;
}
.node > .controls > .control:first-of-type {
  padding-top: 10px;
}
.node > .controls > .control:last-of-type {
  padding-bottom: 2px;
}
</style>
