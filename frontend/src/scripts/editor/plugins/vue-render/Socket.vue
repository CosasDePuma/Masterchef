<template>
  <div
    class="socket"
    :class="classes"
    :title="socket.name.replaceAll(',', ', ')"
  />
</template>

<script>
import filter from "./filter";

export default {
  name: "socket",
  props: ["type", "socket"],
  computed: {
    classes: function() {
      return filter([this.type, this.socket.name]);
    },
  },
};
</script>

<style scoped>
.socket {
  display: inline-block;
  width: 10px;
  height: 10px;
  border-radius: 50%;
  vertical-align: middle;
  background: var(--primary);
  cursor: pointer;
  z-index: 2;
}
.socket.output {
  margin: 0 -4px 0 6px;
}
.socket.input {
  margin: 0 6px 0 -4px;
}
.socket.error {
  background-color: tomato;
}
</style>
