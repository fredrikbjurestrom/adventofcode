import sys

data = [line for line in open("./inputs/day05.txt", "r").read().splitlines()]

seeds = []
mappings = []

map_index = -1
for i in range(len(data)):
    if i == 0:
        seed_input = [int(seed) for seed in data[0].split(': ')[1].split(" ")]
        for s, seed in enumerate(seed_input):
            if s % 2 == 0:
                seeds.append([seed, seed + seed_input[s + 1] - 1])

        continue

    if len(data[i]) == 0:
        continue

    if not data[i][0].isnumeric():
        map_index += 1

        continue

    destination, source, range_length = [int(n) for n in data[i].split(" ")]
    if len(mappings) < map_index + 1:
        mappings.insert(map_index, [[source, source + range_length - 1, destination - source]])
    else:
        mappings[map_index].append([source, source + range_length - 1, destination - source])


def split_range(seed_range, mapping_ranges):
    seed_start, seed_end = seed_range
    print("seed", seed_start, seed_end)

    mapped_ranges = []
    for mapping_start, mapping_end, offset in mapping_ranges:
        print("mapping", mapping_start, mapping_end, offset)

        # mapping covers completely
        if mapping_start <= seed_start and mapping_end >= seed_end:
            print("mapping cover completely")
            mapped_ranges.append([seed_start, seed_end, offset])
        # seed touches start of mapping
        elif seed_start < mapping_start <= seed_end:
            print("seed touches start of mapping")
            # mapping covers end of seed
            if mapping_end >= seed_end:
                print("mapping covers end of seed")
                mapped_ranges.append([mapping_start, seed_end, offset])
            else:
                print("mapping does not cover end of seed")
                mapped_ranges.append([mapping_start, mapping_end, offset])
        # seed touches end of mapping
        elif mapping_start < seed_start <= mapping_end:
            print("seed touches end of mapping")
            mapped_ranges.append([seed_start, mapping_end, offset])

    unmapped_start = seed_start
    unmapped_end = seed_end
    unmapped_ranges = []
    for mapped_start, mapped_end, _ in sorted(mapped_ranges, key=lambda x: x[0]):
        print("hey", unmapped_start, unmapped_end)
        print("mapped", mapped_start, mapped_end)

        if unmapped_start == unmapped_end:
            continue

        if unmapped_end < mapped_start or unmapped_start > mapped_end:
            print("pre or post map")
            unmapped_ranges.append([unmapped_start, unmapped_end])
            unmapped_start = unmapped_end
        elif unmapped_start <= mapped_start:
            print("inside 1")
            unmapped_ranges.append([unmapped_start, mapped_start - 1])
            unmapped_start = mapped_start
        elif unmapped_start >= mapped_start and mapped_end < unmapped_end:
            print("inside 2")
            unmapped_ranges.append([mapped_end + 1, unmapped_end])
            unmapped_start = mapped_start
        else:
            unmapped_start = unmapped_end

    if not unmapped_ranges:
        unmapped_ranges.append([seed_start, seed_end])

    print("ny", [*[[start + offset, end + offset] for start, end, offset in mapped_ranges], *unmapped_ranges])


print(seeds)
print(mappings)
split_range([81, 94], [[15, 51, -15], [52, 53, -15], [0, 14, 39]])
