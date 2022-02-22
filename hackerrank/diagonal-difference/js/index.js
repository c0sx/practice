const diagonalDifference = (arr) => {
    const n = arr.length;
    let left = 0;
    let right = 0;

    for (let i = 0; i < n; ++i) {
        const leftIndex = i;
        const rightIndex = n - i - 1;
        left += arr[i][leftIndex];
        right += arr[i][rightIndex];
    }

    return Math.abs(left - right);
}
