function ulamSequence(u0, u1, n) {
    const sequence = [u0, u1];

    for (let i = 0; i < n - 2; ++i) {
        const next = findSum2(sequence);
        sequence.push(next);
    }

    return sequence
}

function findSum2(sequence) {
    let rightIterator = sequence.length - 1;
    let leftIterator = 0;
    let i = 0;

    const map = new Map();
    while (leftIterator < rightIterator) {
        iter(map, sequence, i)

        i++;
        leftIterator = i + 1;
        rightIterator = sequence.length - (i + 2)
    }

    const values = [...map.entries()]
        .filter(([key, value]) => value === 1)
        .map(([key]) => key);

    return Math.min(...values);
}

function iter(map, sequence, i) {
    let state = 0;
    const start = i;
    const end = sequence.length - (i + 1);

    while (state < 4) {
        const [leftIndex, rightIndex] = getNextIterators(start, end, state)
        if (leftIndex >= rightIndex) {
            return;
        }

        const left = sequence[leftIndex];
        const right = sequence[rightIndex];
        const sum = left + right;
        if (!sequence.includes(sum)) {
            const counter = map.get(sum) || 0;
            map.set(sum, counter + 1);
        }

        state++;
    }
}

function getNextIterators(left, right, state) {
    switch (state) {
        case 0: return [left, right];
        case 1: return [left + 1, right];
        case 2: return [left, right - 1];
        case 3: return [left + 1, right - 1]
    }
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
