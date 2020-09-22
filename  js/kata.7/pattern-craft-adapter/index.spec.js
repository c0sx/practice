const { Mario, MarioAdapter } = require("./index")

describe("Adapter", () => {
    const mario = new Mario();

    it("Mario does not have attack", () => {
        const marioTest = new Mario();

        expect(typeof marioTest.attack).toEqual("undefined");
    });

    it("MarioAdapter can attack", () => {
        const marioAdapter = new MarioAdapter(mario);
        const target = { health: 33 };

        marioAdapter.attack(target);

        expect(target.health).toEqual(30);
    });
});
