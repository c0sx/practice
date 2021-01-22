const validate = (sequence) => {
    const left = ["(", "[", "{"];
    const right = [")", "]", "}"];
    const stack = [];

    const brackets = sequence.split("");
    for (let i = 0; i < brackets.length; ++i) {
        const bracket = brackets[i];
        if (left.includes(bracket)) {
            stack.push(bracket)
            continue;
        }

        const index = right.indexOf(bracket);
        if (index !== -1) {
            const pair = left[index];
            const last = stack.pop();
            if (pair !== last) {
                return false;
            }
        }
    }

    return stack.length === 0;
}

const isValid = validate("[{{{}}}]{[]}]");
console.log(isValid);
