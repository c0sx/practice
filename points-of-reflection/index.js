function symmetricPoint(p, q) {
    const [qx, qy] = q;
    const [px, py] = p;

    const p1x = qx * 2 - px;
    const p1y = qy * 2 - py;
    return [p1x, p1y]
}

module.exports = symmetricPoint
