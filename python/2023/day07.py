from pathlib import Path
from collections import Counter

with open(Path(__file__).parent / "inputs/day07.txt") as file:
    data = file.read().splitlines()

_card_value = {
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9,
    "T": 10,
    "J": 11,
    "Q": 12,
    "K": 13,
    "A": 14,
}


def _score_hand(hand, wildcard=None) -> int:
    """
    Score the hand by type and card value.
    Wildcards and have a value of 1 and counts as any card in the purpose of type scoring.
    """
    hand_counter = Counter("".join(filter(lambda x: x[0] != wildcard, hand)))
    most_common_cards = hand_counter.most_common()
    wildcards = len(list(filter(lambda x: x == wildcard, hand)))

    type_score = 0
    if wildcards == 5 or most_common_cards[0][1] + wildcards == 5:
        type_score = 6
    elif most_common_cards[0][1] + wildcards == 4:
        type_score = 5
    elif most_common_cards[0][1] + most_common_cards[1][1] + wildcards == 5:
        type_score = 4
    elif most_common_cards[0][1] + wildcards == 3:
        type_score = 3
    elif most_common_cards[0][1] + most_common_cards[1][1] + wildcards == 4:
        type_score = 2
    elif most_common_cards[0][1] + wildcards == 2:
        type_score = 1

    card_score = 0
    for i, card in enumerate(hand):
        card_value = _card_value[card] if card != wildcard else 1
        card_score += card_value * (10 ** (8 - (i * 2)))

    return type_score * (10**10) + card_score


def calculate_winnings(wildcard=None) -> int:
    scored_hands = []
    for row in data:
        cards, bet = row.split(" ")

        scored_hands.append((cards, _score_hand(cards, wildcard), int(bet)))

    winnings = 0
    for rank, ranked_hand in enumerate(sorted(scored_hands, key=lambda x: x[1])):
        winnings += (rank + 1) * ranked_hand[2]

    return winnings


if __name__ == "__main__":
    print("Part 1:", calculate_winnings())
    print("Part 2:", calculate_winnings(wildcard="J"))
