from pathlib import Path

with open(Path(__file__).parent / "inputs/day09.txt") as file:
    values = [[int(n) for n in row.split()] for row in file.read().splitlines()]


def _extrapolate(history: list[int]) -> int:
    """Recursively extrapolate the next value from a list of values."""
    if not any(history) or len(history) < 2:
        return 0

    diff = [history[i] - history[i - 1] for i in range(1, len(history))]
    return history[-1] + _extrapolate(diff)


def extrapolate_values(reverse: bool = False) -> int:
    """
    Sum the extrapolated next value of each row.
    If reverse is True, reverse the rows before extrapolating (getting the previous value instead of the next)
    """
    sum = 0
    for value in values:
        ex = _extrapolate(value[::-1] if reverse else value)
        sum += ex

    return sum


if __name__ == "__main__":
    print("Part 1:", extrapolate_values())
    print("Part 2:", extrapolate_values(reverse=True))
