from pathlib import Path
import numpy as np
import itertools

with open(Path(__file__).parent / "inputs/day12_example.txt") as file:
    condition_records = np.array([row.split() for row in file.read().splitlines()])


def valid_record(prospect, groups):
    prospect_character_groups = [len(count) for count in prospect.split(".") if count]
    contiguous_groups = list(map(int, groups.split(",")))

    # if prospect_character_groups == contiguous_groups:
    #     print(prospect, prospect_character_groups, contiguous_groups)

    return prospect_character_groups == contiguous_groups


def count_arrangements(springs, groups):
    unknown_positions = [pos for pos, char in enumerate(springs) if char == "?"]

    combinations = 0
    for combination in itertools.product(".#", repeat=len(unknown_positions)):
        prospect = springs[:]
        for ci, c in enumerate(combination):
            offset = unknown_positions[ci]
            prospect = prospect[:offset] + c + prospect[offset + 1:]

        if valid_record(prospect, groups):
            combinations += 1

    # print("DONE", springs, unknown_positions, groups, combinations)
    return combinations


if __name__ == "__main__":
    print("Part 1:", sum([count_arrangements(*record) for record in condition_records]))
