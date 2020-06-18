class BinOp {
    constructor(left, op, right) {
        this.left = left;
        this.op = op;
        this.token = op;
        this.right = right;
    }
}

module.exports = BinOp;
