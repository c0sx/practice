const MIN = -2147483648; // -Math.pow(2, 31);
const MAX = 2147483647; // Math.pow(2, 31) - 1;

const myAtoi = (s) => {
  let value = 0;
  let sign = 1;
  let index = 0;

  while (s[index] === " ") {
    index += 1;
  }
  
  if (s[index] === "+") {
    index += 1;
  } else if (s[index] === "-") {
    index += 1;
    sign = -1;
  }

  while (index < s.length) {
    const c = s[index];
    if (c < "0" || c > "9") {
      break;
    }

    const digit = Number(c);
    value = value * 10 + digit;

    index += 1;
  }

  value = value * sign;

  if (value < MIN) {
    return MIN;
  } else if (value > MAX) {
    return MAX;
  }

  return value;
};

const assert = (result, expected) => {
  console.log(result, "\t", expected, "\t", result === expected);
};

assert(myAtoi("42"), 42);
assert(myAtoi("-42"), -42);
assert(myAtoi("+42"), 42);
assert(myAtoi("+-42"), 0);
assert(myAtoi("-+42"), 0);
assert(myAtoi("42.24"), 42);
assert(myAtoi("hello 123"), 0);
assert(myAtoi("123 hello"), 123);
assert(myAtoi("  + 413"), 0);
assert(myAtoi("00000-42a1234"), 0);
assert(myAtoi("42.42.42"), 42);
assert(myAtoi(".42.42.42"), 0);
assert(myAtoi(".42"), 0);
