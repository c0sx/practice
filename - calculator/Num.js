class Num {
    constructor(token) {
        this.token = token;
        this.value = token.value
    }

    evaluate() {
        return this.value;
    }
}

module.exports = Num;
