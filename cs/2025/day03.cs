#!/usr/bin/dotnet run

#:project ./lib/lib.csproj

using static Lib.IO;

var part1Joltage = ReadInput("./inputs/day03.txt").Select(part1).Sum();

Console.WriteLine($"Part 1: {part1Joltage}");

int part1(string bank)
{
    var batteries = bank.Select((v, index) => new { joltage = int.Parse($"{v}"), index });

    var first = batteries
        .Where(battery => battery.index < bank.Length - 1)
        .OrderByDescending(battery => battery.joltage)
        .ThenBy(battery => battery.index)
        .First();

    var second = batteries
        .Where(battery => battery.index > first.index)
        .OrderByDescending(battery => battery.joltage)
        .First();

    return int.Parse($"{first.joltage}{second.joltage}");
}
