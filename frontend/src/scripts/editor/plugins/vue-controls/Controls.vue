<template>
  <div class="controls">
    <button class="control save" @click="save">
      <icon-save :size="22" />
    </button>

    <button class="control load" @click="load">
      <icon-load :size="22" />
    </button>

    <button class="control bake" @click="play">
      <img
        src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAGH0lEQVR4AbWXA3Rr6xLH/2GNi/SZ17Zt27Zt27Zts3Z7eWxW54R142xn3szLaq6arIOe/1q/Tr/Ze7DnS761g3x67LHHKhobG79sbm72MSS0tLRQa2vr7xF/Br6noaFhDvN5TU3N2QDsyCMbJpa9urr6nNLS0tfXW289sIW7sBCGrkPXDUaHYZqyht3phNvlhtvtZNyw2xzsV6FpGuLxOAYGBnDttdcWAVAnLLT11lvj97AKi4qKXt91113xt7//HYVc3DQMmFzUsiyk02lYjFgSSP7P4HDYueEycDwcLhdEgUCghMFEOOXPBCqSZPJE/KSZQkTIjs1m+w1KKokYT0ZVVaiayusUP32sq6ent+Pmm29+HUAaOeQcGRnBBHL8upjdbucnc0gjMgVEoxHEo1EupskaxcXFsk2oqvLA46lCV9dSHHTQQecDGGbGmGTOBpBDUtQwDLGZvRwcRDQShq7pqKhcC2UVFSiDDcVcuKigAAWFBSjirRKkUVaAGWA0Jr2yDVgzZ858kIAbeMgutlhvvf9ivf+uB7vTBUVRkGKUlAIbN2jn6TjYL7DGbZJRkFP5G1Buuummt9jWMWU/Tp1dHxgYgy84gjTJwxBskO0RbBAysmUaYrJaxQZUJsSMMZ6vrt8RG2/p4b2uhLuoEk5XGU+ilOuVwNfXh+XLfEhbBlykw9RjsCzCZKrq1oPXpefOX49eunRzeuPqnejNa/eih07bii7e3UMvXr0n/fDOxdTxxoV08d7/opv3X5vxkMRhBbRCs5I9VlQLpRXrYP68LgyPjuHP//Tg/qdvxFkXnoCNNvsvNtz03zjy+H0RSViyE6uwBblFkZg+o7jUtf3pd16L01UN0DRYKQWpRArxWBIq++Sg2nizf+AjLc0Hl9UtcZM1ASMY1r9RUirARZBMQovGkYgnkWIffy2zR7Sm6lIWIzGzBoA5WRNQp/n1mg3+VHBfwutDcUkZD0CKytFscE+CmcXltKeae5IfStxkNaAHo8YwoVTr6woVbLrV+rBM6w+FTctCyDuAotKKKf5INChxk7UFaUYbjqc+7qj/IXscZ4vz/9JQ2jLh6/HBPxz7CkCKsSarAZHy9rSRx+OxaHw0nJIGBFgGw9tgpU1EYil4A/1Lvl0Yqc+OfxLlYP569kbwzai+g+bW3EI9TbfRgurraPYXV1L9K2fTw9fuTzce+udqAB7GhsnU88+/UPb5V9UPNX79Pi1quovaPriavnz+fHrv0dPplXuOo2duPpzuv3xPqv38Deub6voHn3v++TIGK0Leiw8//HB5fX395+3t7bS0q4/GBrpoceMdNK/2dvrxo6up/rUL6f1HTqZX7zmaHrlqT1o05wfq7l1O7e0dJHESzyAfEzsBR21t7UNS2Ov10riWtD5Ec6tvpNnf3kLTv7yefv74Sup4+0Jqff0can3tDJpZfSeNy8dxEi95JB9WQi5++RyYN28eiUzDolgiSbF4ipYv+o6mfX4j/fTRJTTl0ytp6ufXMFfTlE8uoxlf3Uy+JT/wfYrcL3Ekkjz8grpQ8mIF5OTR/bhkyRISRSJRGhkdyzAWobFIgiIJg/oWdND3H15JLa+dTj98dJWsxS/X5b5sjMSLJJ/kXZFzp7ypqYlEQ8MjNDg0/FvYNzQyRsNjMRqJpGg0qoqVtfg5ZpTv+22c5BFJXsmPPLI9/vjjB8yaNYtE/kA/BUMDORikYP8QhcYZyNDPZHyD2Xslj0jySv58X1HXu+++e1dvby8pqkpen598/sAK8fOUKTRl6tRfMYULB+Wa5KFEMkk+n48kv9TJdRI6S0pKti8vL4eSTGXf/0XEltc52WGHHbDddttB2H777Xi9o8QymTh5d5S8/Pa8fb7PQUVdXV2/YRjk5c67e/uob7mXXnn1NbHU07dMfBPS1dMr0NL/2z7qXZaNy+Tx+kjE+UP5PgfryG870ZKlXdTV3UO9XPRlTiS2m5Ozb0XIxL2SiZO15BN1dnaS1EEOrfvZZ599n/nxKbRSe1sbdXR0iKU2XreuAG2/i8te47x8vsQBeJBDZcxmzF7MfmuAfZltmMp8LydljIepWkNU/v5b8D+p18eex/V4oQAAAABJRU5ErkJggg=="
        alt="Cooker"
      />
      <span>Bake!</span>
    </button>

    <button class="control focus" @click="focus">
      <icon-focus :size="22" />
    </button>

    <button class="control delete" @click="remove">
      <icon-delete :size="22" />
    </button>
  </div>
</template>

<script>
import { focus, save } from "@/scripts/helpers";

export default {
  name: "controls",
  props: ["editor"],
  data: function() {
    return { k: 0.2 }; // TODO: Borrar
  },
  methods: {
    save: function() {
      save(this.editor);
    },
    focus: function() {
      focus(this.editor);
    },
    play: function() {
      this.editor.trigger("process");
    },
    load: function() {
      const editor = this.editor;
      const focus = this.focus;
      const element = document.createElement("input");
      element.setAttribute("type", "file");
      element.setAttribute("accept", ".recipe");
      element.addEventListener("change", function(event) {
        const files = event.target.files;
        if (files.length === 0) return;
        const reader = new FileReader();
        reader.addEventListener("load", async function() {
          await editor.fromJSON(JSON.parse(reader.result));
          focus();
        });
        reader.readAsText(files[0]);
      });
      element.click();
    },
    remove: function() {
      while (this.editor.nodes.length > 0) {
        this.editor.nodes.forEach((node) => this.editor.removeNode(node));
      }
    },
  },
};
</script>

<style scoped>
.controls {
  position: absolute;
  bottom: 20px;
  left: 0;
  right: 0;
  width: 100px;
  height: 50px;
  margin: 0 auto;
  display: flex;
  justify-content: center;
  align-items: flex-end;
}
.controls > button {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 10px 0px;
  margin: 0 5px;
  color: var(--decoratordark);
  background: var(--background);
  border: 1px solid transparent;
  box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.21);
  border-radius: 5px;
  cursor: pointer;
  font-size: 18px;
}
.controls > button:active {
  filter: brightness(97%);
}
.controls > button:hover {
  color: var(--primary);
  border-color: var(--primary);
}
.controls > button.bake {
  padding: 13px 35px;
}
.controls > button.bake:hover {
  color: #28a745;
  border-color: #28a745;
}
.controls > button.delete:hover {
  color: tomato;
  border-color: tomato;
}
.controls > button > span {
  margin: 0 10px 0 20px;
}
.controls > button > .material-design-icon {
  margin: 0 15px;
}
</style>
