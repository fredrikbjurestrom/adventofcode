#!/usr/bin/dotnet run

#:project ./lib/lib.csproj

using static Lib.IO;

(int dial, int part1, int part2) agg = (50, 0, 0);

var part1 = ReadInput("./inputs/day01.txt")
    .Aggregate(
        agg,
        (cur, next) =>
        {
            var newPosition = agg.dial + (int.Parse(next[1..]) * (next[0] == 'L' ? -1 : 1));
            agg.part1 += newPosition % 100 == 0 ? 1 : 0;

            if (Math.Abs(newPosition) >= 100)
                agg.part2 += Math.Abs(newPosition) / 100;

            if ((agg.dial != 0 && newPosition <= 0))
                agg.part2 += 1;

            agg.dial = (newPosition % 100) >= 0 ? (newPosition % 100) : 100 + (newPosition % 100);

            return agg;
        }
    );

Console.WriteLine(agg);
