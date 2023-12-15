from pathlib import Path

N = [-1, 0]
S = [1, 0]
E = [0, 1]
W = [0, -1]

cubed_rocks = set()
rounded_rocks = list()

panel_width = 0
panel_height = 0

with open(Path(__file__).parent / "inputs/day14.txt") as file:
    for y, row in enumerate(file.read().splitlines()):
        for x, cell in enumerate(row):
            if cell == "#":
                cubed_rocks.add((y, x))
            elif cell == "O":
                rounded_rocks.append((y, x))

        panel_width = max(panel_width, x + 1)
    panel_height = max(panel_height, y + 1)


def displace(pos, vec) -> tuple[int, int]:
    return tuple(sum(p) for p in zip(pos, vec))


def obstructed(rock, tilted_rocks) -> bool:
    return not (
        -1 < rock[0] < panel_height
        and -1 < rock[1] < panel_width
        and rock not in cubed_rocks
        and rock not in tilted_rocks
    )


def calculate_load(rocks) -> int:
    return sum([panel_height - r[0] for r in rocks])


def tilt(rocks, direction) -> list(tuple[int, int]):
    """
    Tilt the panel in a given direction, making the rounded rocks to roll.
    Ensure that we iterate over the rocks in the reverse order for S and E directions.
    """
    tilted_rocks = set()
    sorted_rocks = rocks if direction == N or direction == W else reversed(rocks)

    for rock in sorted_rocks:
        while not obstructed(displace(rock, direction), tilted_rocks):
            rock = displace(rock, direction)
        tilted_rocks.add(rock)

    return sorted(list(tilted_rocks))


def display_panel(rocks):
    """
    Print the panel based on current rounded rock placements.
    Used for debugging and poor man's hashing function.
    """
    rock_dict = set([tuple(pos) for pos in rocks])
    rows = []
    for y in range(panel_height):
        row = ""
        for x in range(panel_width):
            if (y, x) in rock_dict:
                row += "O"
            elif (y, x) in cubed_rocks:
                row += "#"
            else:
                row += "."
        rows.append(row)

    return "\n".join(rows)


def cycle(rocks, cycles) -> list(tuple[int, int]):
    """
    Find the final rock placements after a given number of cycles.
    Early exit when a reapeating panel placement is found.
    Then we skip the remaining cycles by checking the modulo of the cycle length.
    """
    seen = {display_panel(rocks): 0}
    seq_length = seq_start = 0
    for i in range(1, cycles + 1):
        rocks = tilt(rocks, N)
        rocks = tilt(rocks, W)
        rocks = tilt(rocks, S)
        rocks = tilt(rocks, E)

        panel = display_panel(rocks)
        if panel in seen:
            seq_start = seen[panel]
            seq_length = i - seq_start
            break

        seen[panel] = i

    for panel, i in seen.items():
        if i == (cycles - seq_start) % seq_length + seq_start:
            final_rocks = list()
            for y, row in enumerate(panel.splitlines()):
                for x, cell in enumerate(row):
                    if cell == "O":
                        final_rocks.append((y, x))

            return final_rocks


if __name__ == "__main__":
    print("Part 1:", calculate_load(tilt(rounded_rocks, direction=N)))
    print("Part 2:", calculate_load(cycle(rounded_rocks, cycles=1_000_000_000)))
