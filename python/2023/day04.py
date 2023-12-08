from pathlib import Path

with open(Path(__file__).parent / "inputs/day04.txt") as file:
    data = file.read().splitlines()


def part1() -> int:
    """
    Count the winning numbers and raise add the card score by bit-shifting 1 to the matches-1 exponent. (same as 2 to the power of matches-1)
    """
    total_points = 0

    for row in data:
        _, cards = row.split(": ")
        winning_card, my_card = cards.split("| ")

        winning_numbers = [int(n) for n in winning_card.split(" ") if n]
        my_numbers = [int(n) for n in my_card.split(" ") if n]

        matches = sum([m in winning_numbers for m in my_numbers])

        if matches:
            total_points += 1 << matches - 1

    return total_points


def part2() -> int:
    """
    Instead of scoring points, we're pushing card count down the iteration.
    For each win, use a dictionary to set or update that particular card_id's bonus cards.
    When we get to that card, we're checking the dictionary how many instances of it that we have and using it to count wins.
    """
    total_cards = 0
    wins = {}

    for row in data:
        card_name, cards = row.split(": ")
        card_id = int(card_name.split(" ")[-1])
        winning_card, my_card = cards.split("| ")

        winning_numbers = [int(n) for n in winning_card.split(" ") if n]
        my_numbers = [int(n) for n in my_card.split(" ") if n]

        cards = 1
        if card_id in wins:
            cards += wins[card_id]

        matches = sum([m in winning_numbers for m in my_numbers])
        for i in range(matches):
            winning_card_id = card_id + i + 1

            if winning_card_id in wins:
                wins[winning_card_id] += cards
            else:
                wins[winning_card_id] = cards

        total_cards += cards

    return total_cards


if __name__ == "__main__":
    print("Part 1:", part1())
    print("Part 2:", part2())
