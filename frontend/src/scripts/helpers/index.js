// HTTP request

const headers = new Headers({
  "Content-Type": "application/json",
  "X-Powered-By": "Masterchef!",
});

const config = {
  mode: "cors",
  credentials: "omit",
  cache: "no-cache",
  redirect: "follow",
  keepalive: false,
  referrerPolicy: "no-referrer",
};

export async function call(url, data) {
  return await fetch(url, {
    ...config,
    method: "POST",
    headers: headers,
    body: JSON.stringify(data),
  });
}

// Focus on nodes

export function focus(editor) {
  // helpers
  const min = (arr) => (arr.length === 0 ? 0 : Math.min(...arr));
  const max = (arr) => (arr.length === 0 ? 0 : Math.max(...arr));
  // boundings
  const left = min(editor.nodes.map((node) => node.position[0]));
  const top = min(editor.nodes.map((node) => node.position[1]));
  const right = max(
    editor.nodes.map(
      (node) => node.position[0] + editor.view.nodes.get(node).el.clientWidth
    )
  );
  const bottom = max(
    editor.nodes.map(
      (node) => node.position[1] + editor.view.nodes.get(node).el.clientHeight
    )
  );
  // calculation
  const [x, y] = [(left + right) / 2, (top + bottom) / 2];
  const [w, h] = [
    editor.view.container.clientWidth,
    editor.view.container.clientHeight,
  ];
  const { area } = editor.view;
  const [kw, kh] = [w / Math.abs(left - right), h / Math.abs(top - bottom)];
  const k = Math.min(kh * 0.9, kw * 0.9, 1);
  area.transform.x = area.container.clientWidth / 2 - x * k;
  area.transform.y = area.container.clientHeight / 2 - y * k;
  area.zoom(k, 0, 0);
  area.update();
}

// Save to file

export function save(editor) {
  focus(editor);
  if (!editor.saving) {
    editor.saving = true;
    const blob = new Blob([JSON.stringify(editor.toJSON())]);
    const uri = URL.createObjectURL(blob);
    const element = document.createElement("a");
    element.setAttribute("download", "masterchef.recipe");
    element.setAttribute("href", uri);
    element.click();
    URL.revokeObjectURL(uri);
    editor.saving = false;
  }
}
