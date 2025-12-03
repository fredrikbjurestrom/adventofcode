#!/usr/bin/dotnet run

#:project ./lib/lib.csproj

using static Lib.IO;

var part1Joltage = ReadInput("./inputs/day03.txt").Select(part1).Sum();
var part2Joltage = ReadInput("./inputs/day03.txt").Select(part2).Sum();
Console.WriteLine($"Part 1: {part1Joltage}, Part 2: {part2Joltage}");

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

long part2(string bank)
{
    var batteries = bank.Select((v, index) => new { joltage = int.Parse($"{v}"), index });

    var first = batteries
        .Where(battery => battery.index < bank.Length - 11)
        .OrderByDescending(battery => battery.joltage)
        .ThenBy(battery => battery.index)
        .First();

    var second = batteries
        .Where(battery => battery.index < bank.Length - 10)
        .Where(battery => battery.index > first.index)
        .OrderByDescending(battery => battery.joltage)
        .ThenBy(battery => battery.index)
        .First();

    var third = batteries
        .Where(battery => battery.index < bank.Length - 9)
        .Where(battery => battery.index > second.index)
        .OrderByDescending(battery => battery.joltage)
        .ThenBy(battery => battery.index)
        .First();

    var fourth = batteries
        .Where(battery => battery.index < bank.Length - 8)
        .Where(battery => battery.index > third.index)
        .OrderByDescending(battery => battery.joltage)
        .ThenBy(battery => battery.index)
        .First();

    var fifth = batteries
        .Where(battery => battery.index < bank.Length - 7)
        .Where(battery => battery.index > fourth.index)
        .OrderByDescending(battery => battery.joltage)
        .ThenBy(battery => battery.index)
        .First();

    var sixth = batteries
        .Where(battery => battery.index < bank.Length - 6)
        .Where(battery => battery.index > fifth.index)
        .OrderByDescending(battery => battery.joltage)
        .ThenBy(battery => battery.index)
        .First();

    var seventh = batteries
        .Where(battery => battery.index < bank.Length - 5)
        .Where(battery => battery.index > sixth.index)
        .OrderByDescending(battery => battery.joltage)
        .ThenBy(battery => battery.index)
        .First();

    var eighth = batteries
        .Where(battery => battery.index < bank.Length - 4)
        .Where(battery => battery.index > seventh.index)
        .OrderByDescending(battery => battery.joltage)
        .ThenBy(battery => battery.index)
        .First();

    var ninth = batteries
        .Where(battery => battery.index < bank.Length - 3)
        .Where(battery => battery.index > eighth.index)
        .OrderByDescending(battery => battery.joltage)
        .ThenBy(battery => battery.index)
        .First();

    var tenth = batteries
        .Where(battery => battery.index < bank.Length - 2)
        .Where(battery => battery.index > ninth.index)
        .OrderByDescending(battery => battery.joltage)
        .ThenBy(battery => battery.index)
        .First();

    var eleventh = batteries
        .Where(battery => battery.index < bank.Length - 1)
        .Where(battery => battery.index > tenth.index)
        .OrderByDescending(battery => battery.joltage)
        .ThenBy(battery => battery.index)
        .First();

    var twelveth = batteries
        .Where(battery => battery.index > eleventh.index)
        .OrderByDescending(battery => battery.joltage)
        .ThenBy(battery => battery.index)
        .First();

    return long.Parse(
        $"{first.joltage}{second.joltage}{third.joltage}{fourth.joltage}{fifth.joltage}{sixth.joltage}{seventh.joltage}{eighth.joltage}{ninth.joltage}{tenth.joltage}{eleventh.joltage}{twelveth.joltage}"
    );
}
