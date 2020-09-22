const mostFrequentDays = require("./index")

describe("Most Frequent Weekdays", () => {
    it("tests", () => {
        expect(mostFrequentDays(1084)).toEqual(["Tuesday", "Wednesday"]) // , "Should be: ['Tuesday', 'Wednesday']");
        expect(mostFrequentDays(1167)).toEqual(["Sunday"]) // , "Should be: ['Sunday']");
        expect(mostFrequentDays(1216)).toEqual(["Friday", "Saturday"]) //, "Should be: ['Friday', 'Saturday']");
        expect(mostFrequentDays(1492)).toEqual(["Friday", "Saturday"]) //, "Should be: ['Friday', 'Saturday']");
        expect(mostFrequentDays(1770)).toEqual(["Monday"]) //, "Should be: ['Monday']");
        expect(mostFrequentDays(1785)).toEqual(["Saturday"]) //, "Should be: ['Saturday']");
        expect(mostFrequentDays(212)).toEqual(["Wednesday", "Thursday"]) //, "Should be: ['Wednesday', 'Thursday']");
        expect(mostFrequentDays(1901)).toEqual(["Tuesday"]) //, "Should be: ['Tuesday']");
        expect(mostFrequentDays(2135)).toEqual(["Saturday"]) //, "Should be: ['Saturday']");
        expect(mostFrequentDays(3043)).toEqual(["Sunday"]) //, "Should be: ['Sunday']");
        expect(mostFrequentDays(2001)).toEqual(["Monday"]) //, "Should be: ['Monday']");
        expect(mostFrequentDays(3150)).toEqual(["Sunday"]) //, "Should be: ['Sunday']");
        expect(mostFrequentDays(3230)).toEqual(["Tuesday"]) //, "Should be: ['Tuesday']");
        expect(mostFrequentDays(328)).toEqual(["Monday", "Sunday"]) //, "Should be: ['Monday', 'Sunday']");
        expect(mostFrequentDays(2016)).toEqual(["Friday", "Saturday"]) //, "Should be: ['Friday', 'Saturday']");
        expect(mostFrequentDays(334)).toEqual(["Monday"]) //, "Should be: ['Monday']");
        expect(mostFrequentDays(1986)).toEqual(["Wednesday"]) //, "Should be: ['Wednesday']");
        expect(mostFrequentDays(3361)).toEqual(["Thursday"]) //, "Should be: ['Thursday']");
        expect(mostFrequentDays(5863)).toEqual(["Thursday"]) //, "Should be: ['Thursday']");
        expect(mostFrequentDays(1910)).toEqual(["Saturday"]) //, "Should be: ['Saturday']");
        expect(mostFrequentDays(1968)).toEqual(["Monday", "Tuesday"]) //, "Should be: ['Monday', 'Tuesday']");
        expect(mostFrequentDays(7548)).toEqual(["Thursday", "Friday"]) //, "Should be: ['Thursday', 'Friday']");
        expect(mostFrequentDays(8116)).toEqual(["Wednesday", "Thursday"]) //, "Should be: ['Wednesday', 'Thursday']");
        expect(mostFrequentDays(9137)).toEqual(["Friday"]) //, "Should be: ['Friday']");
        expect(mostFrequentDays(1794)).toEqual(["Wednesday"]) //, "Should be: ['Wednesday']");
    });
})
