fun areaOrPerimeter(w: Int, h: Int): Int {
    return if (w == h) {
        w * h
    } else {
        (w + h) * 2
    }
}

println(areaOrPerimeter(5, 5)) // 25;
println(areaOrPerimeter(10, 2)) // 24
