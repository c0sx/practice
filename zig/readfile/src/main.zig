const std = @import("std");
const file = @embedFile("./foo.txt");

pub fn main() !void {
    var splits = std.mem.splitSequence(u8, file, "\n");

    while (splits.next()) |line| {
        std.debug.print("{s}\n", .{line});
    }
}
