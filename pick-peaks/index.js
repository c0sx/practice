function pickPeaks(arr) {
    // скипнуть края опускающиеся вначале и поднимающиеся в конце
    const sliced = slice(arr)
    const sliced2 = slice(sliced.reverse()).reverse()

    // разбить на чанки
    // ищем когда индекс когда начнется подъем
    const extremum = [];
    let isMax = false;
    const index = sliced2.findIndex((one, index, original) => {
        const next = original[index + 1];
        if (next < one && !isMax) {
            extremum.push(index)
            isMax = true;
            return false;
        }

        return next > one && isMax;
    });

    // найти максимум в чанке
    return sliced2
}

function slice(arr) {
    let i;
    for (i = 1; i < arr.length; ++i) {
        const prev = arr[i - 1];
        const current = arr[i];
        if (current > prev) {
            break;
        }
    }

    return arr.slice(i - 1);
}

module.exports = pickPeaks;

// 1, 2, 3, 6, 4, 1, 2, 3, 2, 1
// [1, 2, 3, 5, 4, 1], [2, 3, 2, 1]

// 3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3
// 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3
// [2, 3, 6, 4, 1], [2, 3, 2, 1]
