from pathlib import Path

seeds = []
mappings = []

with open(Path(__file__).parent / "inputs/day05.txt") as file:
    seeds = list(map(int, file.readline().split(": ")[1].split()))

    mapping_index = -1
    for row in filter(None, map(str.strip, file)):
        if row.endswith(":"):
            mapping_index += 1
            mappings.insert(mapping_index, [])
            continue

        destination, source, range_length = [int(n) for n in row.split(" ")]

        mappings[mapping_index].append(
            (range(source, source + range_length), destination - source)
        )

    mappings = [sorted(ranges, key=lambda x: x[0].stop) for ranges in mappings]


def part1() -> int:
    """Walk the seeds through the mappings and return the smallest mapped seed."""

    min_result = None
    for seed in seeds:
        for mapping in mappings:
            for mapping_range, offset in mapping:
                if seed in mapping_range:
                    seed += offset
                    break

        if not min_result or seed < min_result:
            min_result = seed

    return min_result


def map_range(seed_ranges, mapping_ranges) -> list[range]:
    """
    Compares one seed range with one type of mapping ranges (seed, soil etc).
    Returns new ranges with mapped or unmapped seeds.
    Mapping ranges needs to be sorted for this to work.
    """
    mapped_seed_ranges = []
    for seed_range in seed_ranges:
        seed = seed_range.start
        for mapping_range, offset in mapping_ranges:
            if seed < mapping_range.start:
                mapped_seed_ranges.append(range(seed, mapping_range.start))
                seed = mapping_range.stop
            elif seed in mapping_range:
                new_stop = min(seed_range.stop, mapping_range.stop)
                mapped_seed_ranges.append(range(seed + offset, new_stop + offset))
                seed = new_stop

            if seed >= seed_range.stop:
                break

        if seed < seed_range.stop:
            mapped_seed_ranges.append(range(seed, seed_range.stop))

    return mapped_seed_ranges


def part2() -> int:
    min_result = None
    for seed_range in zip(
        [seeds[i] for i in range(len(seeds)) if i % 2 == 0],
        [seeds[i] for i in range(len(seeds)) if i % 2 == 1],
    ):
        seed_ranges = [range(seed_range[0], seed_range[0] + seed_range[1])]
        for mapping in mappings:
            seed_ranges = map_range(seed_ranges, mapping)

        min_seed = min([r.start for r in seed_ranges])
        if not min_result or min_seed < min_result:
            min_result = min_seed

    return min_result


if __name__ == "__main__":
    print("Part 1:", part1())
    print("Part 2:", part2())
