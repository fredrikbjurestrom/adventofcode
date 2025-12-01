#!/usr/bin/dotnet run

#:project ./lib/lib.csproj

using static Lib.IO;

var result = ReadInput("./inputs/day01.txt")
    .Aggregate(
        (dial: 50, part1: 0, part2: 0),
        (acc, next) =>
        {
            var newPos = acc.dial + (int.Parse(next[1..]) * (next[0] == 'L' ? -1 : 1));
            var part1 = acc.part1 + (newPos % 100 == 0 ? 1 : 0);
            var part2 = acc.part2 + (Math.Abs(newPos) >= 100 ? Math.Abs(newPos) / 100 : 0);

            if (acc.dial != 0 && newPos <= 0)
                part2 += 1;

            var dial = (newPos % 100) >= 0 ? (newPos % 100) : 100 + (newPos % 100);

            return (dial, part1, part2);
        }
    );

Console.WriteLine($"Part 1: {result.part1}, Part 2: {result.part2}");
