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
        for (let i = index; i < sequence.length; ++i) {
            const one = sequence[i]

            if (one === current) {
                continue;
            }

            const sum = current + one;
            const pair = [current, one].sort((a, b) => a - b);

            const values = map.get(sum) || [];
            const index = values.findIndex(([v1, v2]) => {
                const [p1, p2] = pair;
                return v1 === p1 && p2 === v2;
            });

            if (index !== -1) {
                return;
            }

            values.push(pair);
            map.set(sum, values)
        }

        return map;
    }, new Map());

    const values = [...sums.entries()]
        .reduce((acc, [key, value]) => {
            if (value.length === 1 && !sequence.includes(key)) {
                acc.push(key)
            }

            return acc;
        }, []);
    return Math.min(...values);
}

module.exports = ulamSequence
