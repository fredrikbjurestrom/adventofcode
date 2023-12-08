# import sys
from pathlib import Path

with open(Path(__file__).parent / "inputs/day05_example.txt") as file:
    seeds = [int(seed) for seed in file.readline().split(": ")[1].split()]

    map_index = -1
    mapping_ranges = []
    for row in filter(None, map(str.strip, file)):
        if not row[0].isnumeric():
            map_index += 1
            continue

        destination, source, range_length = [int(n) for n in row.split(" ")]

        if len(mapping_ranges) < map_index + 1:
            mapping_ranges.insert(
                map_index,
                [(range(source, source + range_length), destination - source)],
            )
        else:
            mapping_ranges[map_index].append(
                (range(source, source + range_length), destination - source)
            )

print(mapping_ranges)
def part1():
    print("hej")

# def parse_input():

#     seeds = []
#     mappings = []

#     map_index = -1
#     for i in range(len(data)):
#         if i == 0:
#             seed_input = [int(seed) for seed in data[0].split(': ')[1].split(" ")]
#             for s, seed in enumerate(seed_input):
#                 if s % 2 == 0:
#                     seeds.append([seed, seed + seed_input[s + 1] - 1])

#             continue

#         if len(data[i]) == 0:
#             continue

#         if not data[i][0].isnumeric():
#             map_index += 1

#             continue

#         destination, source, range_length = [int(n) for n in data[i].split(" ")]
#         if len(mappings) < map_index + 1:
#             mappings.insert(map_index, [[source, source + range_length - 1, destination - source]])
#         else:
#             mappings[map_index].append([source, source + range_length - 1, destination - source])

#     return seeds, mappings

# def part1():
#     seeds, mappings = parse_input()

#     print(mappings)

# part1()
