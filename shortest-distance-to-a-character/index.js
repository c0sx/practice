function shortesttoChar(input, target) {
    if (target === "" || input === "" || input.indexOf(target) === -1) {
        return [];
    }

    return input.split("").map((char, index) => {
        if (char === target) {
            return 0;
        }

        const left = input.substring(0, index);
        const right = input.substring(index);

        const existsInLeft = left.search(target);
        let leftDistance = -1;
        if (existsInLeft !== -1) {
            leftDistance = 0;
            for (let i = index; i >= 0; --i) {
                if (left[i] === target) {
                    break;
                }

                leftDistance++;
            }
        }

        let rightDistance = right.search(target);
        if (leftDistance === -1) {
            return rightDistance;
        }

        if (rightDistance === -1) {
            return leftDistance;
        }

        return rightDistance > leftDistance ? leftDistance : rightDistance;
    })
}

console.log(shortesttoChar("lovecodewars", "e")) //  === [3, 2, 1, 0, 1, 2, 1, 0, 1, 2, 3, 4]
console.log(shortesttoChar("aaaabbbb", "b")) //  === [4, 3, 2, 1, 0, 0, 0, 0]
console.log(shortesttoChar("", "b")) //  === []
console.log(shortesttoChar("abcde", "")) //  === []
