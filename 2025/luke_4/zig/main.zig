const std = @import("std");

fn countNeighbors(map: *std.AutoHashMap(usize, []const u8), row: usize, col: usize, target: u8) usize {
    const directions = [_][2]i32{
        .{ -1, -1 }, .{ -1, 0 }, . { -1, 1 }, // top-left, top, top-right
        .{ 0, - 1 },             . { 0, 1  }, // left, right
        .{ 1, - 1 }, .{ 1, 0  },  .{ 1, 1  }, // bottom-left, bottom, bottom-right
    };

    var count: usize = 0;

    for (directions) |dir| {
        const new_row = @as(i32, @intCast(row)) + dir[0];
        const new_col = @as(i32, @intCast(col)) + dir[1];

        if (new_row < 0 or new_col < 0) continue;

        const r = @as(usize, @intCast(new_row));
        const c = @as(usize, @intCast(new_col));

        if (map.get(r)) |line| {
            if (c < line.len and line[c] == target) {
                count += 1;
            }
        }
    }

    return count;
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const file = try std.fs.cwd().openFile("../input.txt", .{});
    defer file.close();

    var map = std.AutoHashMap(usize, []const u8).init(allocator);
    defer map.deinit();

    var read_buf: [4096]u8 = undefined;
    var file_reader = file.reader(&read_buf);

    var row: usize = 0;
    while (try file_reader.interface.takeDelimiter('\n')) |line| {
        const line_copy = try allocator.dupe(u8, line);
        try map.put(row, line_copy);
        row += 1;
    }

    var task1: usize = 0;
    var it = map.iterator();
    while (it.next()) |entry| {
        const row_idx = entry.key_ptr.*;
        const line = entry.value_ptr.*;

        for (line, 0..) |char, col_idx| {
            if (char != '@') {
                continue;
            }
            const count = countNeighbors(&map, row_idx, col_idx, '@');
            if (count < 4) {
                task1 += 1;
            }
        }
    }

    var task2: usize = 0;
    var changed: bool = true;
    while (changed) {
        var newit = map.iterator();
        changed = false;
        while (newit.next()) |entry| {
            const row_idx = entry.key_ptr.*;
            const line = entry.value_ptr.*;

            for (line, 0..) |char, col_idx| {
                if (char != '@') {
                    continue;
                }
                const count = countNeighbors(&map, row_idx, col_idx, '@');
                if (count < 4) {
                    task2 += 1;
                    const mutable_line = @constCast(line);
                    mutable_line[col_idx] = '.';
                    changed = true;
                }
            }
        }
    }

    std.debug.print("task1: {}\n", .{task1});
    std.debug.print("task2: {}\n", .{task2});

    var it2 = map.valueIterator();
    while (it2.next()) |value| {
        allocator.free(value.*);
    }
}
