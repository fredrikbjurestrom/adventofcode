from pathlib import Path

with open(Path(__file__).parent / "inputs/day03.txt") as file:
    data = file.read().splitlines()


def _has_adjacent_symbol(digit_x, digit_y) -> bool:
    for adjacent_y in range(digit_y - 1, digit_y + 2):
        if adjacent_y < 0 or adjacent_y >= len(data):
            continue

        for adjacent_x in range(digit_x - 1, digit_x + 2):
            if adjacent_x < 0 or adjacent_x >= len(data):
                continue

            neighbour = data[adjacent_y][adjacent_x]
            if not neighbour.isnumeric() and neighbour != ".":
                return True

    return False


def part1() -> int:
    """
    Iterate the input and for every digit, concat into a number and check if any surrounding points have a symbol in them.
    Keep doing it in the same row until a non digit character (or line break) appears, then add the concatenated value if at least one digit neighbours a symbol.
    """
    parts_adjacent_to_symbol = []

    for y, row in enumerate(data):
        value = ""
        is_symbol_adjacent = False

        for x, cell in enumerate(row):
            if cell.isnumeric():
                value += cell
                is_symbol_adjacent = _has_adjacent_symbol(x, y) or is_symbol_adjacent
            elif value:
                if is_symbol_adjacent:
                    parts_adjacent_to_symbol.append(int(value))
                value = ""
                is_symbol_adjacent = False

        if value and is_symbol_adjacent:
            parts_adjacent_to_symbol.append(int(value))

    return sum(parts_adjacent_to_symbol)


def _get_neighbouring_gear(digit_x, digit_y) -> tuple[int, int] | None:
    for adjacent_y in range(digit_y - 1, digit_y + 2):
        if adjacent_y < 0 or adjacent_y >= len(data):
            continue

        for adjacent_x in range(digit_x - 1, digit_x + 2):
            if adjacent_x < 0 or adjacent_x >= len(data):
                continue

            if data[adjacent_y][adjacent_x] == "*":
                return adjacent_x, adjacent_y

    return None


def _append_dict(dictionary, keys, value_to_add) -> None:
    for key in keys:
        if key in dictionary:
            dictionary[key].append(value_to_add)
        else:
            dictionary[key] = [value_to_add]


def part2() -> int:
    """
    Replace the adjacency function to one which only identifies gears (*) and returns its position.
    Then save each identified gear into a dictionary which keep tracks of all the "gear prospect" numbers.
    After that, we're filtering the gears with exactly two neighbours and sums their product.
    """
    gear_candidates = {}

    for y, row in enumerate(data):
        value = ""
        adjacent_gears = set()

        for x, cell in enumerate(row):
            if cell.isnumeric():
                value += cell
                gear_candidate = _get_neighbouring_gear(x, y)
                if gear_candidate:
                    adjacent_gears.add(gear_candidate)
            elif value:
                if adjacent_gears:
                    _append_dict(gear_candidates, adjacent_gears, int(value))
                value = ""
                adjacent_gears = set()

        if value and adjacent_gears:
            _append_dict(gear_candidates, adjacent_gears, int(value))

    gears = [
        gear_candidates[pos][0] * gear_candidates[pos][1]
        for pos in gear_candidates
        if len(gear_candidates[pos]) == 2
    ]

    return sum(gears)


if __name__ == "__main__":
    print("Part 1:", part1())
    print("Part 2:", part2())
