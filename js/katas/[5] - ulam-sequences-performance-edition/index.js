function ulamSequence(u0, u1, n) {
    const used = new Set()
    let left = [u0, u1];
    let right = [u0 + u1];
    let remains = [];

    for (let i = 0; i < n - 2; ++i) {
        [left, remains, right] = getSums(left, remains, right, used);
        const kek = 5;
    }

    return left;
}

const getSums = (left, remains, right, used) => {
    const sums = [];

    // нужно не складывать уже складываемые рание слагаемые
    for (let i = 0; i < right.length; ++i) {
        const rightTerm = right[i];
        if (rightTerm === undefined) {
            continue;
        }

        left.forEach(leftTerm => {
            const sum = leftTerm + rightTerm;
            if (used.has(sum)) {
                const sumsIndex = sums.indexOf(sum);
                if (sumsIndex !== -1) {
                    sums.splice(sumsIndex, 1)
                }

                const rightIndex = right.indexOf(sum);
                if (rightIndex !== -1) {
                    right[rightIndex] = undefined
                    const kek = 10;
                }

                return;
            }

            // если сумма уже существует в правом срезе, удалить это значение из среза
            const rightIndex = right.indexOf(sum);
            if (rightIndex !== -1) {
                right[rightIndex] = undefined;
                const kek = rightIndex;
                used.add(sum)
            }

            // если сумма уже существует в собранных суммах, то удалить ее оттуда
            // (а удалять вроде бы нельзя потому что в каком-то кейсе она повторно появится)
            const sumsIndex = sums.indexOf(sum);
            if (sumsIndex === -1 && rightIndex === -1) {
                sums.push(sum)
            }
            else if (sumsIndex !== -1) {
                sums.splice(sumsIndex, 1)
                used.add(sum)
            }

            // удалить так же из остатков
            const remainsIndex = remains.indexOf(sum);
            if (remainsIndex !== -1) {
                remains.splice(remainsIndex, 1)
                used.add(sum)

                const index = sums.indexOf(sum);
                sums.splice(index, 1);
            }
        });
    }

    const total = [...remains, ...right, ...sums].sort((a, b) => a - b);
    const first = total.shift();
    left.push(first);

    const remainsIndex = remains.indexOf(first);
    if (remainsIndex !== -1) {
        remains.splice(remainsIndex, 1);
    }

    const rightIndex = right.indexOf(first);
    if (rightIndex !== -1) {
        right.splice(rightIndex, 1);
    }

    return [left, right.filter(one => one !== undefined), sums]
}

module.exports = ulamSequence
