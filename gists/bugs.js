class tree {
    constructor(min, max, state) {
        this._value = 0;
        this._leftMin = min;
        this._rightMax = max;
        this._left = null;
        this._right = null;
        this._state = state;

        if (!this._state) {
            const left = this._leftMax - this._leftMin;
            const right = this._rightMax - this._rightMin;
            if (right <= left) {
                this._state = "left";
            }
            else {
                this._state = "right";
            }
        }

        if (this._state === "right") {
            this._value = Math.floor((max + min) / 2);
        }
        else {
            this._value = Math.ceil((max + min) / 2);
        }

        this._leftMax = this._value - 1;
        this._rightMin = this._value + 1;
    }

    slice() {
        if (this._state === "left" && !this._left) {
            this._state = "right";
            return this._left = new tree(this._leftMin, this._leftMax, "right");
        }
        else if (this._state === "right" && !this._right) {
            this._state = "left";
            return this._right = new tree(this._rightMin, this._rightMax, "left");
        }

        if (this._state === "left") {
            this._state = "right";
            return this._left.slice();
        }

        this._state = "left";
        return this._right.slice();
    }

    getRocks() {
        const state = this._state === "right" ? "left" : "right";

        if (state === "right" && this._right) {
            return this._right.getRocks();
        }
        else if (state === "left" && this._left) {
            return this._left.getRocks();
        }

        return [this._leftMax - this._leftMin + 1, this._rightMax - this._rightMin + 1];
    }
}

const rocks = 11;
const bugs = 5;
// 8 1 - 3, 4
// 8 2 - 1, 2
// 8 3 - 1, 1

const t = new tree(1, rocks);
let r = null;
for (let bug = 1; bug < bugs; ++bug) {
    t.slice();
}

console.log(t.getRocks());
