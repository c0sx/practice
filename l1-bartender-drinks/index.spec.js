const getDrinkByProfession = require("./index")

describe("L1: Bartender, drinks!", () => {
    it("tests", () => {
        expect(getDrinkByProfession("jabrOni")).toEqual("Patron Tequila") // , "'Jabroni' should map to 'Patron Tequila'");
        expect(getDrinkByProfession("scHOOl counselor")).toEqual("Anything with Alcohol") // , "'School Counselor' should map to 'Anything with alcohol'");
        expect(getDrinkByProfession("prOgramMer")).toEqual("Hipster Craft Beer") //, "'Programmer' should map to 'Hipster Craft Beer'");
        expect(getDrinkByProfession("bike ganG member")).toEqual("Moonshine") // , "'Bike Gang Member' should map to 'Moonshine'");
        expect(getDrinkByProfession("poLiTiCian")).toEqual("Your tax dollars") // , "'Politician' should map to 'Your tax dollars'");
        expect(getDrinkByProfession("rapper")).toEqual("Cristal") // , "'Rapper' should map to 'Cristal'");
        expect(getDrinkByProfession("pundit")).toEqual("Beer") // , "'Pundit' should map to 'Beer'");
        expect(getDrinkByProfession("Pug")).toEqual("Beer") // , "'Pug' should map to 'Beer'");

    });
})
