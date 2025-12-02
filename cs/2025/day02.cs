#!/usr/bin/dotnet run

#:project ./lib/lib.csproj

using static Lib.IO;

var result = ReadInput("./inputs/day02-test.txt")
    .SelectMany(x => x.Split(','))
    .Select(x => new { from = long.Parse(x.Split('-')[0]), to = long.Parse(x.Split('-')[1]) })
    .SelectMany(x => expand(x.from, x.to))
    .Select(id => new { part1 = part1(id), part2 = part2(id) })
    .GroupBy(x => 1)
    .Select(g => new { part1 = g.Sum(x => x.part1), part2 = g.Sum(x => x.part2) })
    .First();

Console.WriteLine($"Part 1: {result.part1}, Part 2: {result.part2}");

IEnumerable<long> expand(long from, long to)
{
    for (var id = from; id <= to; id++)
    {
        yield return id;
    }
}

long part1(long id)
{
    var s = $"{id}";
    if (s.Length % 2 != 0)
        return 0;

    return s[..(s.Length / 2)] == s[(s.Length / 2)..] ? id : 0;
}

long part2(long id)
{
    return 0;
}
