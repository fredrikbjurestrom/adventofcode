from pathlib import Path

with open(Path(__file__).parent / "inputs/day02.txt") as file:
    data = file.read().splitlines()


def part1() -> int:
    """Parse the rows and check if game is below limit."""
    limit = {"red": 12, "green": 13, "blue": 14}
    valid_games = []

    for row in data:
        valid = True
        game, details = row.split(": ")
        game_id = int(game.split(" ")[1])

        for grabs in details.split("; "):
            for grab in grabs.split(", "):
                qty, color = grab.split(" ")
                if int(qty) > limit[color]:
                    valid = False

        if valid:
            valid_games.append(game_id)

    return sum(valid_games)


def part2() -> int:
    """Move limit dict inside row loop and repurpose it to store max seen value."""
    game_powers = []

    for row in data:
        _, details = row.split(": ")

        min_scores = {"red": 0, "green": 0, "blue": 0}

        for grabs in details.split("; "):
            for grab in grabs.split(", "):
                qty, color = grab.split(" ")
                if int(qty) > min_scores[color]:
                    min_scores[color] = int(qty)

        game_powers.append(min_scores["red"] * min_scores["green"] * min_scores["blue"])

    return sum(game_powers)


if __name__ == "__main__":
    print("Part 1:", part1())
    print("Part 2:", part2())
