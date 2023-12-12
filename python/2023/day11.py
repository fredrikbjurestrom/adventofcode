from pathlib import Path
import numpy as np
import itertools

with open(Path(__file__).parent / "inputs/day11.txt") as file:
    grid = np.array([[c for c in row] for row in file.read().splitlines()])
    galaxies = np.array([g for g in np.argwhere(grid == "#")])


def expand_space(unexpanded_space, expansion):
    """
    First we find all space rows and columns (all ".":s)
    Then we iterate each galaxy and increment the x and y value by an incrementing offset of the expansion
    """
    space_columns = np.all(unexpanded_space[:, 0:] == ".", axis=0)
    x_offset = 0
    for space_x in np.arange(unexpanded_space.shape[1])[space_columns]:
        for galaxy in galaxies:
            if galaxy[1] > space_x + x_offset:
                galaxy[1] += expansion - 1
        x_offset += expansion - 1

    space_rows = np.all(unexpanded_space[:, 0:] == ".", axis=1)
    y_offset = 0
    for space_y in np.arange(unexpanded_space.shape[0])[space_rows]:
        for galaxy in galaxies:
            if galaxy[0] > space_y + y_offset:
                galaxy[0] += expansion - 1
        y_offset += expansion - 1

    return galaxies


def calculate_distance(expansion: int) -> int:
    """Manhattan distance between two points"""
    expanded_galaxies = expand_space(grid, expansion)

    distance = 0
    for from_galaxy, to_galaxy in itertools.combinations(expanded_galaxies, 2):
        ydiff, xdiff = np.subtract(from_galaxy, to_galaxy)
        distance += abs(ydiff) + abs(xdiff)

    return distance


if __name__ == "__main__":
    print("Part 1:", calculate_distance(expansion=2))
    print("Part 2:", calculate_distance(expansion=1_000_000))
