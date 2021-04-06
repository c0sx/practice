class MarioAdapter {
    constructor(mario) {
        this.mario = mario;
    }

    attack(target) {
        target.health -= this.mario.jumpAttack()
    }
}

class Marine {
    attack(target) {
        target.health -= 6;
    }
}

class Zealot {
    attack(target) {
        target.health -= 8;
    }
}

class Zergling {
    attack(target) {
        target.health -= 5;
    }
}

class Mario {
    jumpAttack() {
        console.log("Mamamia!");
        return 3;
    }
}

module.exports = {
    Mario,
    Zealot,
    Zergling,
    MarioAdapter
}
