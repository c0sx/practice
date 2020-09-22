fun isDivisible(n: Int, x: Int, y: Int): Boolean = n % x == 0 && n % y == 0

println(isDivisible(3, 3, 4)) // false
println(isDivisible(12, 3, 4)) // true
println(isDivisible(8, 3, 4)) // false
println(isDivisible(48, 3, 4)) // true
println(isDivisible(100, 5, 10)) // true
println(isDivisible(100, 5, 3)) // false