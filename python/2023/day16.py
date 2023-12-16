from pathlib import Path

N = [-1, 0]
S = [1, 0]
E = [0, 1]
W = [0, -1]

grid = dict()
grid_height = 0
grid_width = 0

with open(Path(__file__).parent / "inputs/day16.txt") as file:
    for y, row in enumerate(file.read().splitlines()):
        for x, symbol in enumerate(row):
            grid[(y, x)] = symbol

        grid_width = max(grid_width, x + 1)
    grid_height = max(grid_height, y + 1)


def displace(pos, vec) -> tuple[int, int]:
    return tuple(sum(p) for p in zip(pos, vec))


def out_of_bounds(pos) -> bool:
    return pos[0] < 0 or pos[0] >= grid_height or pos[1] < 0 or pos[1] >= grid_width


def beam(pos, direction, energized=dict()) -> set[tuple[int, int]]:
    """
    Beam traverses the grid of mirrors, stopping when out of bounds or entering the same grid from the same direction (kept by energized dict).
    Splitting of beams are handled through recursive calls.
    """
    next = displace(pos, direction)

    while not out_of_bounds(next):
        symbol = grid[next]
        if next in energized:
            if tuple(direction) in energized[next]:
                break
            else:
                energized[next].append(tuple(direction))
        else:
            energized[next] = [tuple(direction)]

        if direction in (E, W):
            if symbol == "/":
                direction = N if direction == E else S
            elif symbol == "\\":
                direction = S if direction == E else N
            elif symbol == "|":
                energized.update(beam(next, N, energized))
                direction = S

            next = displace(next, direction)

        elif direction in (N, S):
            if symbol == "/":
                direction = E if direction == N else W
            elif symbol == "\\":
                direction = W if direction == N else E
            elif symbol == "-":
                energized.update(beam(next, W, energized))
                direction = E

            next = displace(next, direction)

    return energized


def find_max_illumination():
    """Just brute force all the angles, beam function is efficient enough"""
    max_illumination = 0

    for x in range(grid_width):
        max_illumination = max(len(beam((-1, x), S, dict())), max_illumination)
        max_illumination = max(len(beam((grid_height, x), N, dict())), max_illumination)

    for y in range(grid_height):
        max_illumination = max(len(beam((y, -1), E, dict())), max_illumination)
        max_illumination = max(len(beam((y, grid_height), W, dict())), max_illumination)

    return max_illumination


if __name__ == "__main__":
    print("Part 1:", len(beam((0, -1), E)))
    print("Part 2:", find_max_illumination())
