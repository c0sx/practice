class Graph {
  constructor() {
    this.nodes = new Map();
    this.edges = new Set();
  }

  link(from, to) {
    const nodes = this.nodes.get(from) ?? new Set();
    nodes.add(to);

    if (nodes.size < 2) {
      this.edges.add(from);
    } else {
      this.edges.delete(from);
    }

    this.nodes.set(from, nodes);
  }

  traverse() {
    const list = new Set();

    this.edges.forEach((key) => this._traverse(key, list));

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
