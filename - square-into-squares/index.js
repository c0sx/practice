// до тех пор пока существует квадрат, значение которого выше остатка от числа, продолжаем добавлять их в массив
// проверка суммы на каждой итерации и подбор ближайших квадратов

function decompose(digit) {
    const result = [];
    const square = digit ** 2;

    const root = digit - 1;
    result.push(root);

    let iter = [root, square]
    do {
        iter = iteration(iter[1], iter[0]);
        const [root] = iter;
        if (root < 1) {
            return null;
        }

        result.push(root);
    }
    while (assertResult(result, digit) === false)

    return result.reverse();
}

function iteration(square, root) {
    const remains1 = square - root ** 2;
    const root1 = getNextSquareRoot(remains1);
    return [root1, remains1];
}

function getNextSquareRoot(remains) {
    let middle = Math.ceil(remains * 2 / 3);
    while (Math.trunc(Math.sqrt(middle)) !== Math.sqrt(middle)) {
        middle++
    }

    return Math.sqrt(middle);
}

function assertResult(result, target) {
    const square = target ** 2;
    const sum = result.reduce((acc, curr) => acc + curr ** 2, 0);
    return square === sum;
}

console.log(decompose(85986)); // [1, 3, 6, 23, 414, 85985]
console.log(decompose(50)); // [1, 3, 5, 8, 49] 1 + 9 + 25 + 64 + 2401 = 2500
console.log(decompose(11)); // [1, 2, 4, 10] 1 + 4 + 16 + 100 = 121
console.log(decompose(7)); //  [2, 3, 6]
console.log(decompose(4)) // "Nothing"
console.log(decompose(2)) // "Nothing"



