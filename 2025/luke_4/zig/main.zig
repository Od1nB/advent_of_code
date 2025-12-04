const std = @import("std");

fn countNeighbors(factory: [][]u8, row: usize, col: usize, target: u8) usize {
    const directions = [_][2]i32{
        .{ -1, -1 }, .{ -1, 0 }, .{ -1, 1 },
        .{ 0, -1 },  .{ 0, 1 },  .{ 1, -1 },
        .{ 1, 0 },   .{ 1, 1 },
    };

    var count: usize = 0;

    for (directions) |dir| {
        const new_row = @as(i32, @intCast(row)) + dir[0];
        const new_col = @as(i32, @intCast(col)) + dir[1];

        if (new_row < 0 or new_col < 0) continue;
        if (new_row >= factory.len) continue;

        const r = @as(usize, @intCast(new_row));
        const c = @as(usize, @intCast(new_col));

        if (c < factory[r].len and factory[r][c] == target) {
            count += 1;
        }
    }

    return count;
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const content = try std.fs.cwd().readFileAlloc(allocator, "../input.txt", 1024 * 1024);
    defer allocator.free(content);

    var factory = try std.ArrayList([]u8).initCapacity(allocator, 140);
    defer {
        for (factory.items) |line| {
            allocator.free(line);
        }
        factory.deinit(allocator);
    }

    var lines = std.mem.splitScalar(u8, content, '\n');
    while (lines.next()) |line| {
        if (line.len == 0) continue;
        const line_copy = try allocator.dupe(u8, line);
        try factory.append(allocator, line_copy);
    }

    var task1: usize = 0;
    for (factory.items, 0..) |line, row_idx| {
        for (line, 0..) |char, col_idx| {
            if (char != '@') {
                continue;
            }
            const count = countNeighbors(factory.items, row_idx, col_idx, '@');
            if (count < 4) {
                task1 += 1;
            }
        }
    }

    var task2: usize = 0;
    var changed: bool = true;
    while (changed) {
        changed = false;
        for (factory.items, 0..) |line, row_idx| {
            for (line, 0..) |char, col_idx| {
                if (char != '@') {
                    continue;
                }
                const count = countNeighbors(factory.items, row_idx, col_idx, '@');
                if (count < 4) {
                    task2 += 1;
                    factory.items[row_idx][col_idx] = '.';
                    changed = true;
                }
            }
        }
    }

    std.debug.print("task1: {}\n", .{task1});
    std.debug.print("task2: {}\n", .{task2});
}
