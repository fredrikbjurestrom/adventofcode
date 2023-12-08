from pathlib import Path
import sys

with open(Path(__file__).parent / "inputs/day05.txt") as file:
    data = file.read().splitlines()


def part1() -> int:
    """
    First, get all the seeds (starting number), then iterate through all the maps and save the lowest mapped number. Name of the map is saved to ensure that map can't occur more than once.
    """
    seeds = [int(seed) for seed in data[0].split(': ')[1].split(" ")]

    min_result = sys.maxsize

    for seed in seeds:
        map_name = ""
        performed_maps = set()

        for i in range(2, len(data)):
            if len(data[i]) == 0:
                continue

            if not data[i][0].isnumeric():
                map_name = data[i]
                continue

            if map_name in performed_maps:
                continue

            destination, source, range_length = [int(n) for n in data[i].split(" ")]

            if seed in range(source, source + range_length):
                performed_maps.add(map_name)
                seed += (destination - source)

        if seed < min_result:
            min_result = seed

    return min_result


def part2() -> int:
    """Gah!"""
    seed_ranges = []
    mapping_ranges = []

    map_index = -1
    for i in range(len(data)):
        if i == 0:
            seed_input = [int(seed) for seed in data[0].split(': ')[1].split(" ")]
            for s, seed in enumerate(seed_input):
                if s % 2 == 0:
                    seed_ranges.append(range(seed, seed + seed_input[s + 1]))
            continue

        if len(data[i]) == 0:
            continue

        if not data[i][0].isnumeric():
            map_index += 1
            continue

        destination, source, range_length = [int(n) for n in data[i].split(" ")]

        if len(mapping_ranges) < map_index + 1:
            mapping_ranges.insert(map_index, [(range(source, source + range_length), destination - source)])
        else:
            mapping_ranges[map_index].append((range(source, source + range_length), destination - source))


    def split_range(seed_range, map_ranges):
        # print("seed", seed_range, map_ranges)

        mapped_ranges = []
        for map_range, offset in map_ranges:
            # mapping covers seed completely
            if map_range.start <= seed_range.start and map_range.stop >= seed_range.stop:
                # print("mapping covers seed completely", map_range)
                mapped_ranges.append((seed_range, offset))
            # seed covers mapping completely
            elif seed_range.start <= map_range.start and seed_range.stop >= map_range.stop:
                # print("seed covers mapping completely", map_range)
                mapped_ranges.append((map_range, offset))
            # seed touches start of mapping
            elif seed_range.start <= map_range.start < seed_range.stop:
                # print("seed touches start of mapping", map_range)
                mapped_ranges.append((range(map_range.start, seed_range.stop), offset))
            # seed touches end of mapping
            elif map_range.start <= seed_range.start < map_range.stop:
                # print("seed touches end of mapping", map_range)
                mapped_ranges.append((range(seed_range.start, map_range.stop), offset))

        unmapped_ranges = []
        cur_pos = seed_range.start
        for mapped_range, _ in sorted(mapped_ranges, key=lambda x: x[0].start):
            if cur_pos < mapped_range.start:
                unmapped_ranges.append(range(cur_pos, mapped_range.start))
                cur_pos = mapped_range.stop
            elif seed_range.stop > mapped_range.stop:
                unmapped_ranges.append(range(mapped_range.stop, seed_range.stop))
                cur_pos = seed_range.stop

        if not mapped_ranges:
            unmapped_ranges.append(seed_range)

        # print("returning",
        #       [*[range(mapped_range.start + off, mapped_range.stop + off) for mapped_range, off in mapped_ranges],
        #        *unmapped_ranges])
        return [*[range(mapped_range.start + off, mapped_range.stop + off) for mapped_range, off in mapped_ranges],
                *unmapped_ranges]


    min_result = sys.maxsize


    for seed in seed_ranges:
        queue = [[seed]]

        for map_id, mapping_range in enumerate(mapping_ranges):
            while len(queue[map_id]) > 0:
                cur_seed = queue[map_id].pop()
                split_ranges = split_range(cur_seed, mapping_range)
                
                if len(queue) <= map_id + 1:
                    queue.insert(map_id+1, list(split_ranges))
                else:
                    queue[map_id+1].extend(split_ranges)
                

        # print(queue[-1])
        min_seed = min([seed[0] for seed in queue[-1]])
        if min_seed < min_result:
            min_result = min_seed
            

    return min_result


if __name__ == "__main__":
    print("Part 1:", part1())
    print("Part 2:", part2())
