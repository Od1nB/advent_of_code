use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let file = File::open("../input.txt")?;
    let reader = BufReader::new(file);

    let mut factory: Vec<Vec<char>> = Vec::with_capacity(140);
    for line in reader.lines() {
        let line = line?;
        factory.push(line.chars().collect());
    }

    let mut task1 = 0;
    for y in 0..factory.len() {
        for x in 0..factory[y].len() {
            if factory[y][x] != '@' {
                continue;
            }
            let num = adjacent_rolls(&factory, x, y);
            if num < 4 {
                task1 += 1;
            }
        }
    }

    let mut task2 = 0;
    let mut changed = true;
    while changed {
        changed = false;
        for y in 0..factory.len() {
            for x in 0..factory[y].len() {
                if factory[y][x] != '@' {
                    continue;
                }
                let num = adjacent_rolls(&factory, x, y);
                if num < 4 {
                    task2 += 1;
                    factory[y][x] = '.';
                    changed = true;
                }
            }
        }
    }

    println!("task1 {}", task1);
    println!("task2 {}", task2);

    Ok(())
}

fn adjacent_rolls(factory: &[Vec<char>], x: usize, y: usize) -> i32 {
    let mut count = 0;
    let max_y = factory.len();

    let directions = [
        (-1, -1),
        (0, -1),
        (1, -1),
        (-1, 0),
        (1, 0),
        (-1, 1),
        (0, 1),
        (1, 1),
    ];

    for (dx, dy) in directions {
        let new_y = y as i32 + dy;
        let new_x = x as i32 + dx;

        if new_y < 0 || new_y >= max_y as i32 || new_x < 0 {
            continue;
        }

        let row = &factory[new_y as usize];
        if new_x < row.len() as i32 && row[new_x as usize] == '@' {
            count += 1;
        }
    }

    count
}
