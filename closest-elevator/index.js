function elevator(left, right, call){
    return (Math.abs(left-call) < Math.abs(right-call)) ? 'left' : 'right';
}

console.log(elevator(0, 1, 0)); // => "left"
console.log(elevator(0, 1, 1)); // => "right"
console.log(elevator(0, 1, 2)); // => "right"
console.log(elevator(0, 0, 0)); // => "right"
console.log(elevator(0, 2, 1)); // => "right"
