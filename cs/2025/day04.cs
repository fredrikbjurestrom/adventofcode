#!/usr/bin/dotnet run

#:project ./lib/lib.csproj

using static Lib.IO;

var grid = ReadInput("./inputs/day04.txt").Select(row => row.ToArray()).ToArray();

var part1 = grid.SelectMany(
        (line, y) => line.Where((pos, x) => pos == '@' && accessibleByForklift(grid, y, x))
    )
    .Count();

var part2 = 0;
while (true)
{
    var papers = grid.SelectMany(
            (line, y) =>
                line.Where(
                    (pos, x) =>
                    {
                        if (pos != '@')
                            return false;

                        if (accessibleByForklift(grid, y, x))
                        {
                            grid[y][x] = '.';
                            return true;
                        }

                        return false;
                    }
                )
        )
        .Count();

    if (papers == 0)
        break;

    part2 += papers;
}

Console.WriteLine($"Part 1: {part1}, Part 2: {part2}");

bool accessibleByForklift(char[][] grid, int y, int x)
{
    var papers = 0;
    var neighbours = Enumerable.Range(-1, 3).ToArray();

    foreach (var ymov in neighbours)
    {
        var neighbourY = y + ymov;
        if (neighbourY < 0 || neighbourY >= grid.Length)
            continue;

        foreach (var xmov in neighbours.Where(x => ymov != 0 || x != 0))
        {
            var neighbourX = x + xmov;
            if (neighbourX < 0 || neighbourX >= grid[neighbourY].Length)
                continue;

            if (grid[neighbourY][neighbourX] == '@')
                papers += 1;
        }
    }

    return papers < 4;
}
