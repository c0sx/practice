const climbStairs = (n) => {
  if (n === 1) {
    return 1;
  }

  let previous = 1;
  let current = 1;

  for (let i = 2; i <= n; i++) {
    let temp = current;
    current = previous + current;
    previous = temp;
  }

  return current;
};
