import { Input, Output } from "rete";

import { newSocket } from "./sockets";
import { newControl } from "./controls";

/**
 * Creates the inputs
 * @param {object} node Current node
 * @param {object} inputs Specification of the inputs
 * @param {object} editor Current visual editor
 */
function setInputs(node, inputs, editor) {
  inputs.forEach(function(input) {
    // create
    const element = new Input(input.name, input.name, newSocket(input.type));
    // inline controls
    if (input.control) {
      const ctrl = newControl(editor, node, {
        ...input.control,
        name: input.name,
      });
      if (ctrl) element.addControl(ctrl);
    }
    // append
    node.addInput(element);
  });
}

/**
 * Creates the outputs
 * @param {object} node Current node
 * @param {object} outputs Specification of the outputs
 */
function setOutputs(node, outputs) {
  outputs.forEach(function(output) {
    node.addOutput(
      new Output(output.name, output.name, newSocket(output.type))
    );
  });
}

/**
 * Creates the controls
 * @param {object} node Current node
 * @param {object} controls Specification of the controls
 * @param {object} editor Current visual editor
 */
function setControls(node, controls, editor) {
  controls.forEach(function(ctrl) {
    const element = newControl(editor, node, { ...ctrl });
    if (element) node.addControl(element);
  });
}

/**
 * Node builder
 * @param {object} node Current node
 * @param {object} spec Specification of the node
 * @param {object} editor Current visual editor
 * @returns
 */
export default function builder(node, spec, editor) {
  if (spec.inputs) setInputs(node, spec.inputs, editor);
  if (spec.outputs) setOutputs(node, spec.outputs);
  if (spec.controls) setControls(node, spec.controls, editor);
  return node;
}
