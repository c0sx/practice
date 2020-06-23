const sortAnimal = require("./index");

describe("Sort my animals", () => {
    it("test", () => {
        const animal = [
            { name: "Cat", numberOfLegs: 4 },
            { name: "Snake", numberOfLegs: 0 },
            { name: "Dog", numberOfLegs: 4 },
            { name: "Pig", numberOfLegs: 4 },
            { name: "Human", numberOfLegs: 2 },
            { name: "Bird", numberOfLegs: 2 }
        ];

        expect(sortAnimal(animal)).toEqual([
            { name: "Snake", numberOfLegs: 0 },
            { name: "Bird", numberOfLegs: 2 },
            { name: "Human", numberOfLegs: 2 },
            { name: "Cat", numberOfLegs: 4 },
            { name: "Dog", numberOfLegs: 4 },
            { name: "Pig", numberOfLegs: 4 }
        ]);
    });
});
