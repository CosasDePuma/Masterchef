import validate, { validations } from "@/scripts/validator";
import { call } from "@/scripts/helpers";
import { endpoint } from "@/assets/ingredients";

/**
 * Checks if the node should be processed based on the inputs
 * @param {object} spec Specification of the input nodes
 * @param {object} node Current node
 * @param {object} inputs Node inputs
 * @returns true if the node has the correct inputs
 */
function checkInputs(spec, node, inputs) {
  for (let i = 0; i < spec.length; i++) {
    const { name, type, optional } = spec[i];
    // connected
    if (!inputs[name].length) {
      // control
      if (node.data && node.data[name] !== undefined)
        inputs[name].push(node.data[name]);
      // default value
      else if (spec[i].default !== undefined)
        value = inputs[name].push(spec[i].default);
      // input needed
      else if (!optional) return false;
    }
    // format
    if (type.toLowerCase().endsWith("s") && !Array.isArray(inputs[name][0])) {
      inputs[name][0] = [inputs[name][0]];
    }

    // value
    if (Array.isArray(inputs[name][0])) {
      for (const value of inputs[name][0]) {
        if (!validate(type, value)) return false;
      }
    } else if (!validate(type, inputs[name][0])) return false;
  }
  return true;
}

/**
 * Node worker
 * @param {object} spec Node specification
 * @param {object} node Current node
 * @param {object} inputs Inputs
 * @param {object} outputs Outputs
 * @param {object} controls Controls
 * @returns nothing
 */
export default async function worker(spec, node, inputs, outputs, controls) {
  const previous = outputs;
  // inputs
  if (spec.inputs && !checkInputs(spec.inputs, node, inputs)) return;
  const flatInputs = {};
  Object.keys(inputs).forEach(function(input) {
    flatInputs[input] = inputs[input].length ? inputs[input][0] : undefined;
  });
  // worker
  await (async ({ endpoint, node, inputs, outputs, controls, call, regex }) =>
    await eval(`(async () => { ${spec.code} })()`))({
    // params
    node,
    inputs: flatInputs,
    outputs,
    controls,
    // globals
    call,
    endpoint,
    regex: validations,
  });
  // outputs
  if (previous !== outputs) this.editor.trigger("process");
}
