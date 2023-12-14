from pathlib import Path
import numpy as np

N = [-1, 0]
S = [1, 0]
E = [0, 1]
W = [0, -1]


with open(Path(__file__).parent / "inputs/day14_example.txt") as file:
    panel = np.array([[c for c in row] for row in file.read().split("\n")])
    cube_rocks = set([tuple(pos) for pos in np.argwhere(panel == "#")])
    rounded_rocks = list([tuple(pos) for pos in np.argwhere(panel == "O")])


def out_of_bounds(y, x) -> bool:
    return y < 0 or x < 0 or y >= panel.shape[0] or x >= panel.shape[1]


def tilt(rocks, direction) -> list(tuple[int, int]):
    resting_rocks = set()
    for rock in rocks:
        while True:
            y, x = np.add(rock, direction)
            if out_of_bounds(y, x) or (y, x) in cube_rocks or (y, x) in resting_rocks:
                resting_rocks.add(rock)
                break

            rock = (y, x)

    return list(resting_rocks)


def calculate_load(rocks) -> int:
    return sum([panel.shape[0]-r[0] for r in rocks])


def spin_cycle(rocks, cycles) -> list(tuple[int, int]):
    for _ in range(cycles):
        rocks = tilt(rocks, N)
        rocks = tilt(rocks, W)
        rocks = tilt(rocks, S)
        rocks = tilt(rocks, E)

    return rocks


if __name__ == "__main__":
    print("Part 1:", calculate_load(tilt(rounded_rocks, direction=N)))
    print("Part 2:", calculate_load(
        spin_cycle(rounded_rocks, cycles=1_000_000_000)))
