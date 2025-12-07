#!/usr/bin/dotnet run

#:project ./lib/lib.csproj

using static Lib.IO;

List<(long From, long To)> ranges = new();
var rangeDefinition = true;
var result = ReadInput("./inputs/day05.txt")
    .Where(line =>
    {
        if (line.Length == 0)
        {
            rangeDefinition = false;
            return false;
        }

        if (rangeDefinition)
        {
            var fromAndTo = line.Split('-');
            ranges.Add((long.Parse(fromAndTo[0]), long.Parse(fromAndTo[1])));
            return false;
        }

        var ingredient = long.Parse(line);
        foreach (var range in ranges)
        {
            if (ingredient >= range.From && ingredient <= range.To)
                return true;
        }

        return false;
    });

var part1 = result.Count();

var exlusiveRanges = ranges
    .OrderBy(x => x.From)
    .ThenBy(x => x.To)
    .Aggregate(
        new List<(long From, long To)>(),
        (acc, next) =>
        {
            var prev = acc.LastOrDefault();

            if (next.To > prev.To)
            {
                acc.Add((prev.To >= next.From ? prev.To + 1 : next.From, next.To));
            }

            return acc;
        }
    );

var part2 = exlusiveRanges.Select(x => x.To - x.From + 1).Sum();

Console.WriteLine($"Part 1: {part1}, Part 2: {part2}");
