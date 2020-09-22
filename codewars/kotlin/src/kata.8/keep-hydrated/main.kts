import kotlin.math.floor

fun litres(time: Double): Int = floor(time / 2.0).toInt()

println(litres(2.0)) // 1
println(litres(1.4)) // 0
println(litres(12.3)) // 6
println(litres(0.82)) // 0
println(litres(11.8)) // 5
println(litres(1787.0)) // 893
println(litres(0.0)) // 0