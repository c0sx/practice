const std = @import("std");

pub fn main() !void {
    var n: usize = 1;

    while (n <= 100) : (n += 1) {
        if (n % 3 == 0 and n % 5 == 0) {
            std.debug.print("fizzbuzz\n", .{});
        } else if (n % 3 == 0) {
            std.debug.print("fizz\n", .{});
        } else if (n % 5 == 0) {
            std.debug.print("buzz\n", .{});
        }
    }
}
