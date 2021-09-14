const lengthOfLongestSubstring = function(s) {
    const chunk = [];
    let chars = new Set();
    let result = 0;

    for (let i = 0; i < s.length; ++i) {
        const c = s.charAt(i);

        if (chars.has(c) === false) {
            chars.add(c);
            chunk.push(c);
            result = Math.max(chars.size, result);
            continue;
        }

        while (true) {
            const removed = chunk.shift();
            if (removed === c) {
                chunk.push(c);
                chars = new Set(chunk);
                break;
            }
        }
    }

    return result;
};

console.log(lengthOfLongestSubstring("abcabcbb"));
console.log(lengthOfLongestSubstring(" "));
