const grid = require("./grid")
// найти самую верхнюю строку, в которой нужно произвести изменения
// зафиксировать строки которые уже собраны правильно


module.exports = input => {
    // пройтись по всем строкам сверху вниз и пометить клетки как зафиксированные
    let marked = mark(input)
    for (let i = 0; i < marked.length; ++i) {
        if (rowIsFinished(marked[i], i)) {
            marked[i].forEach(cell => cell.finished = true);
            marked = grid.rotate(marked);
        }
    }

    const findNextRow = () => {
        const index = marked.findIndex((row, index) => {
            const factor = row.length * index + 1;
            const expected = row.map((row, index) => index + factor);
            return row.some((cell, index) => cell.value !== expected[index]);
        })

        const factor = input[index].length * index + 1;
        return input[index].map((cell, index) => index + factor)
    }

    return {
        marked,
        findNextRow
    }
}


const mark = (input) => {
    return input.map(row => {
        return row.map(cell => {
            return {
                value: cell,
                finished: false
            }
        })
    })
}

const rowIsFinished = (row, index) => {
    const factor = row.length * index + 1;
    const expected = Array.from({ length: row.length }).map((one, index) => index + factor)
    return row.every((cell, index) => cell.value === expected[index])
}
