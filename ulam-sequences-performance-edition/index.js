function ulamSequence(u0, u1, n) {
    const sequence = [u0, u1];

    for (let i = 0; i < n - 2; ++i) {
        const next = findMinSumInSequence(sequence);
        sequence.push(next);
    }

    return sequence;
}

function findMinSumInSequence(sequence) {
    const sums = sequence.reduce((map, current, index, sequence) => {

        for (let i = index + 1; i < sequence.length; ++i) {
            const one = sequence[i]
            const sum = current + one;
            if (sequence.includes(sum)) {
                continue;
            }

            const counter = map.get(sum) || 0;
            map.set(sum, counter + 1)
        }

        return map;
    }, new Map());

    const values = [...sums.entries()]
        .reduce((acc, [key, value]) => {
            if (value === 1) {
                acc.push(key)
            }


            return acc;
        }, []);

    return Math.min(...values);
}

module.exports = ulamSequence
