/**
 * @param {number[]} digits
 * @return {number[]}
 */
const plusOne = (digits) => {
  let index = digits.length - 1;

  while (index >= 0) {
    let num = digits[index] + 1;
    if (num < 10) {
      digits[index] = num;

      return digits;
    }

    let remainder = num % 10;
    digits[index] = remainder;

    index -= 1;
  }

  digits.unshift(1);

  return digits;
};

console.log(plusOne([1, 2, 3]));
console.log(plusOne([9]));
console.log(plusOne([9, 9, 9]));
