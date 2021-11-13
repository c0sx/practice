const paul = require("./index");

describe("Paul's Misery", () => {
    it("Testing for fixed tests", () => {
        expect(paul(['life', 'eating', 'life'])).toEqual('Super happy!');
        expect(paul(['life', 'Petes kata', 'Petes kata', 'Petes kata', 'eating'])).toEqual('Super happy!');
        expect(paul(['Petes kata', 'Petes kata', 'eating', 'Petes kata', 'Petes kata', 'eating'])).toEqual('Happy!');
    });
});
