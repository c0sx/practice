function pickPeaks(arr) {
    const chunks = splitToChunks(arr);
    let iter = 0;

    return chunks.reduce((acc, current, index) => {
        if (index === 0 && !firstChunk(current)) {
            iter += current.length
            return acc;
        } else if (index === chunks.length - 1 && !lastChunk(current)) {
            iter += current.length
            return acc;
        }

        const peak = Math.max(...current);
        const currentIndex = current.findIndex(one => one === peak);
        const pos = currentIndex + iter;

        iter += current.length;
        acc.pos.push(pos);
        acc.peaks.push(peak);
        return acc;
    }, { pos: [], peaks: [] });
}

function firstChunk(chunk) {
    const peak = Math.max(...chunk);
    const index = chunk.indexOf(peak);
    return [0, chunk.length - 1].includes(index) === false;
}

function lastChunk(chunk) {
    const result = chunk.every((value, index, original) => {
        if (index + 1 > original.length - 1) {
            return true;
        }

        const next = original[index + 1];
        return value <= next;
    });

    return result === false;
}

function splitToChunks(arr) {
    const chunks = [];

    while(arr.length) {
        let isMax = false;
        const index = arr.findIndex((one, index, original) => {
            const next = original[index + 1];
            if (next < one && !isMax) {
                isMax = true;
                return false;
            }

            return next > one && isMax;
        });

        if (index === -1) {
            chunks.push([...arr])
            arr.splice(-arr.length);
        }
        else {
            chunks.push(arr.slice(0, index + 1));
            arr.splice(0, index + 1);
        }
    }

    return chunks;
}

module.exports = pickPeaks;
