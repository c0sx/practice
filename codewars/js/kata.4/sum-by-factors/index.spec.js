const sumOfDivided = require("./index")

describe("Sum by Factors", () => {
    it("test", () => {
        expect(sumOfDivided([12, 15])).toEqual([ [2, 12], [3, 27], [5, 15] ]);
        expect(sumOfDivided([15, 21, 24, 30, 45])).toEqual([ [2, 54], [3, 135], [5, 90], [7, 21] ]);
        expect(sumOfDivided([-29804, -4209, -28265, -72769, -31744]))
            .toEqual([
                [ 2, -61548 ],
                [ 3, -4209 ],
                [ 5, -28265 ],
                [ 23, -4209 ],
                [ 31, -31744 ],
                [ 53, -72769 ],
                [ 61, -4209 ],
                [ 1373, -72769 ],
                [ 5653, -28265 ],
                [ 7451, -29804 ]
            ])
    });
})
