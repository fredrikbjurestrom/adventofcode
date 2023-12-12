from pathlib import Path
import numpy as np
import itertools


with open(Path(__file__).parent / "inputs/day11.txt") as file:
    grid = np.array([[c for c in row] for row in file.read().splitlines()])

def expand_space(expansion):
    space = grid.copy()
    insert_y = 0
    for match in np.all(space[:, 0:] == ".", axis=0):
        if match:
            for _ in range(1, expansion):
                space = np.insert(space, insert_y, ".", axis=1)
                insert_y += 1

        insert_y += 1

    insert_x = 0
    for match in np.all(space[:, 0:] == ".", axis=1):
        if match:
          for _ in range(1, expansion):
            space = np.insert(space, insert_x, ".", axis=0)
            insert_x += 1
        insert_x += 1

    galaxies = [tuple(g) for g in np.argwhere(space == "#")]

    return galaxies


    
def part1(galaxies) -> int:
    distance = 0
    for from_galaxy, to_galaxy in itertools.combinations(galaxies, 2):
        ydiff, xdiff = np.subtract(from_galaxy, to_galaxy)
        distance += abs(ydiff) + abs(xdiff)

    return distance


if __name__ == "__main__":
    print("Part 1:", part1(expand_space(2)))
    # print("Part 2:", part1(expand_space(1_000_000)))
