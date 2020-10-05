// # 1
Bomb.diffuse(42);

// # 2
Bomb.diffuse(50);
Bomb.diffuse(1);
Bomb.diffuse(2);
Bomb.diffuse(3);
Bomb.diffuse(4);

// # 3
Bomb.diffuse(BombKey);

// # 4
const diffuseTheBomb = () => true;
Bomb.diffuse();

// # 5
Bomb.diffuse(3.14159);

// # 6
const now = new Date();
now.setFullYear(now.getFullYear() - 4)
Bomb.diffuse(now);

// # 7
const key = {
    key: 43
};

Bomb.diffuse(Object.freeze(key));

// # 8
class A {
    constructor() {
        this.value = 9;
    }

    toString() {
        const a = this.value;
        this.value = 12;
        return a;
    }
}

const a = new A();

Bomb.diffuse(a)

// # 9
let k = 0.5
Math.random = () => {
    const a = k;
    k = 1;
    return a;
}

Bomb.diffuse(42)

// # 10
const value = "yes i enjoy";
const encoded = Buffer.from(value, "ascii").toString("base64");
const decoded = Buffer.from(encoded, "base64").toString("ascii");

Array.prototype.toString = value => 14

Bomb.diffuse(encoded)
