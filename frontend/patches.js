const fs = require("fs");
const path = require("path");

// ReteJS
const rete = path.resolve(__dirname, "./node_modules/rete/build/rete.esm.js");
fs.writeFileSync(
  rete,
  fs
    .readFileSync(rete, "utf8")
    .replace(
      "var delta = (wheelDelta ? wheelDelta / 120 : -e.deltaY / 3) * this.intensity;",
      "var delta = (wheelDelta ? wheelDelta / 120 : -e.deltaY / 50) * this.intensity;"
    ) // Fix zoom @mousewheel
);
