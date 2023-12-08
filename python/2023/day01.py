from pathlib import Path

with open(Path(__file__).parent / "inputs/day01.txt") as file:
    data = file.read().splitlines()


def part1() -> int:
    """Get numeric positions for each row, concatenate the first and last occurrence and turn it into an int."""
    numbers = []

    for row in data:
        hits = [i for i, c in enumerate(row) if c.isnumeric()]
        numbers.append(int(row[min(hits)] + row[max(hits)]))

    return sum(numbers)


def _ends_with_number(current_row) -> list | None:
    """
    Retrieves the string up until the current index. Checks if it ends with a number, albeit a digit or a word.
    If it's a match, the position and the digit is returned (in string form, since we're concatenating the first and last occurrence).
    """
    pos = len(current_row) - 1

    if current_row[-1:].isnumeric():
        return [pos - 1, current_row[-1:]]

    number_words = {
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9",
    }

    for word_length in (3, 4, 5):
        if (len(current_row) >= word_length and current_row[-word_length:] in number_words):
            return [pos - word_length, number_words[current_row[-word_length:]]]


def part2() -> int:
    """Replace isnumeric call with custom function that returns a list with both the position and the actual value."""
    numbers = []

    for row in data:
        hits = dict([_ends_with_number(row[:i]) for i in range(1, len(row) + 1) if _ends_with_number(row[:i])])
        numbers.append(int(hits[min(hits)] + hits[max(hits)]))

    return sum(numbers)


if __name__ == "__main__":
    print("Part 1:", part1())
    print("Part 2:", part2())
