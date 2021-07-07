import Vue from "vue";
import Node from "./Node.vue";

function createVue(el, vueComponent, vueProps, options = {}) {
  const app = new Vue({
    render: (h) => h(vueComponent, { props: vueProps }),
    ...options,
  });

  const nodeEl = document.createElement("div");

  el.appendChild(nodeEl);
  app.$mount(nodeEl);

  return app;
}

function createNode(
  editor,
  CommonVueComponent,
  { el, node, component, bindSocket, bindControl },
  options
) {
  const vueComponent = component.component || CommonVueComponent || Node;
  const vueProps = {
    ...component.props,
    node,
    editor,
    bindSocket,
    bindControl,
  };
  const app = createVue(el, vueComponent, vueProps, options);

  node.vueContext = app.$children[0];

  return app;
}

function createControl(editor, { el, control }, options) {
  const vueComponent = control.component;
  const vueProps = {
    ...control.props,
    editor,
    getData: control.getData.bind(control),
    putData: control.putData.bind(control),
  };
  const app = createVue(el, vueComponent, vueProps, options);

  control.vueContext = app.$children[0];

  return app;
}

const update = (entity) => {
  return new Promise((res) => {
    if (!entity.vueContext) return res();

    entity.vueContext.$forceUpdate();
    entity.vueContext.$nextTick(res);
  });
};

function install(editor, { component: CommonVueComponent, options }) {
  editor.on(
    "rendernode",
    ({ el, node, component, bindSocket, bindControl }) => {
      if (component.render && component.render !== "vue") return;
      node._vue = createNode(
        editor,
        CommonVueComponent,
        { el, node, component, bindSocket, bindControl },
        options
      );
      node.update = async () => await update(node);
    }
  );

  editor.on("rendercontrol", ({ el, control }) => {
    if (control.render && control.render !== "vue") return;
    control._vue = createControl(editor, { el, control }, options);
    control.update = async () => await update(control);
  });

  editor.on("connectioncreated connectionremoved", (connection) => {
    update(connection.output.node);
    update(connection.input.node);
  });

  editor.on("nodeselected", () => {
    editor.nodes.map(update);
  });
}

export default { name: "vue-render", install };
