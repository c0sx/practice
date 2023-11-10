class Graph {
  constructor() {
    this.nodes = new Map();
  }

  link(from, to) {
    const nodes = this.nodes.get(from) ?? [];
    nodes.push(to);

    this.nodes.set(from, nodes);
  }

  traverse() {
    const root = [...this.nodes.keys()].find(
      (key) => this.nodes.get(key).length === 1
    );

    const list = new Set();

    this._traverse(root, list);

    return list;
  }

  _traverse(key, set) {
    if (set.has(key)) {
      return;
    }

    set.add(key);

    const adjacents = this.nodes.get(key);
    for (const one of adjacents) {
      this._traverse(one, set);
    }
  }
}

const restoreArray = (list) => {
  const graph = new Graph();

  for (const [left, right] of list) {
    graph.link(left, right);
    graph.link(right, left);
  }

  const result = graph.traverse();

  return [...result];
};
