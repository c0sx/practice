function validParentheses(parens){
    const queue = [];
    const chars = parens.split("");

    for (let i = 0; i < chars.length; ++i) {
        const char = chars[i];
        if (char === "(") {
            queue.push(char);
            continue;
        }

        const last = queue.length > 0 ? queue[queue.length - 1] : undefined;
        if (char === ")" && last !== "(") {
            return false;
        }

        queue.splice(queue.length - 1);
    }

    return queue.length === 0;
}

module.exports = validParentheses;
