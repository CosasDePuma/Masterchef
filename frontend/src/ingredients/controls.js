import Ctrl from "@/scripts/editor/components";

/**
 * Creates a new control
 * @param {object} editor Current visual editor
 * @param {object} node Current node
 * @param {object} specs Control specification
 * @returns the new control
 */
export function newControl(
  editor,
  node,
  specs = { name: "", type: "", options: {} }
) {
  if (!specs.options) specs.options = {};
  if (specs.options.default !== undefined)
    node.data[specs.name] = specs.options.default;
  switch (specs.type) {
    case "area":
      return new Ctrl.Area(editor, specs.name, specs.options);
    case "check":
      return new Ctrl.Check(editor, specs.name, specs.options);
    case "input":
      return new Ctrl.Input(editor, specs.name, specs.options);
    case "options":
      return new Ctrl.Options(editor, specs.name, specs.options);
    default:
      break;
  }
}
