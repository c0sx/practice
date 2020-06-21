const disemvowel = require("./index")

describe("Disemvowel Trolls", () => {
    it("tests", () => {
        expect(disemvowel("This website is for losers LOL!")).toEqual("Ths wbst s fr lsrs LL!");
        expect(disemvowel("No offense but,\nYour writing is among the worst I've ever read")).toEqual("N ffns bt,\nYr wrtng s mng th wrst 'v vr rd");
        expect(disemvowel("What are you, a communist?")).toEqual("Wht r y,  cmmnst?");
    });
})
