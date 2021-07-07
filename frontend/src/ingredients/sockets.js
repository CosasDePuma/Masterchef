import { Socket } from "rete";

const sockets = {
  any: new Socket("any"),
};

/**
 * Creates nor returns a socket
 * @param {string} name Socket name/type
 * @returns the socket
 */
export function newSocket(name) {
  // create
  if (sockets[name] === undefined) {
    sockets[name] = new Socket(name);
  }
  // process
  if (name.includes(",")) {
    name
      .split(",")
      .map((n) => n.trim().toLowerCase())
      .filter((n) => n !== "")
      .forEach(function(n) {
        newSocket(n);
        sockets[n].combineWith(sockets[name]);
      });
  }
  // register
  sockets[name].combineWith(sockets.any);
  // return
  return sockets[name];
}
