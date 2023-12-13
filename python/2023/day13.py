from pathlib import Path
import numpy as np

with open(Path(__file__).parent / "inputs/day13.txt") as file:
    patterns = [[row for row in pattern.split("\n")] for pattern in file.read().split("\n\n")]


def is_match(pattern1, pattern2) -> bool:
    return "".join(pattern1) == "".join(pattern2)


def is_perfect_mirror(arr, start_pos):
    length = 0
    pos = start_pos
    while pos > 0:
        length += 1
        pos -= 1
        mirrored = pos + 1 + (length * 2)
        if mirrored >= arr.shape[0]:
            return True

        if not is_match(arr[pos], arr[mirrored]):
            return False

    return True


def part1():
    sum = 0
    for pattern in patterns:
        arr = np.array([list(row) for row in pattern])

        for y in range(arr.shape[0] - 1):
            if is_match(arr[y], arr[y + 1]):
                if is_perfect_mirror(arr, y):
                    sum += (y + 1) * 100
                    break

        for x in range(arr.shape[1] - 1):
            if is_match(arr[:, x], arr[:, x + 1]):
                if is_perfect_mirror(np.swapaxes(arr, 0, 1), x):
                    sum += x + 1
                    break

    return sum


if __name__ == "__main__":
    print("Part 1:", part1())
