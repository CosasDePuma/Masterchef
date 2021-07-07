import { save } from "@/scripts/helpers";

const zoomStep = 0.1;

function install(editor) {
  // prevent zoom on dblclick
  editor.on("zoom", function({ source }) {
    return source !== "dblclick";
  });

  // keyboard
  editor.on("keydown", function(event) {
    if (event.code === "Delete") {
      //delete nodes
      editor.selected.each(function(node) {
        editor.removeNode(node);
      });
    } else if (event.ctrlKey) {
      // zooms
      if (event.key === "+") {
        event.preventDefault();
        editor.view.area.zoom(editor.view.area.transform.k + zoomStep);
      } else if (event.key === "-") {
        event.preventDefault();
        editor.view.area.zoom(editor.view.area.transform.k - zoomStep);
      }
      // save
      else if (event.code === "KeyS") {
        event.preventDefault();
        save(editor);
      }
    }
  });

  // drag & drop
  editor.view.container.addEventListener("dragover", function(event) {
    event.preventDefault();
  });

  editor.view.container.addEventListener("drop", async function(event) {
    if (!event.dataTransfer) return;
    const spec = event.dataTransfer.getData("spec");
    if (!spec) return;
    const data = JSON.parse(spec);
    const ingredient = editor.components.get(data.name);
    if (!ingredient) return;
    editor.view.area.pointermove(event);
    const node = await ingredient.createNode();
    node.position = [editor.view.area.mouse.x, editor.view.area.mouse.y];
    editor.addNode(node);
  });
}

export default { name: "events", install };
