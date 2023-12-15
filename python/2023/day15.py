from pathlib import Path


with open(Path(__file__).parent / "inputs/day15.txt") as file:
    sequence = [step for step in file.read().split(",")]


def holiday_hash(step) -> int:
    value = 0
    for c in step:
        value += ord(c)
        value *= 17
        value %= 256

    return value


def HASHMAP(sequence) -> int:
    boxes = [[] for _ in range(256)]
    for step in sequence:
        if "=" in step:
            label, value = step.split("=")
            value = int(value)
            existing = False
            for i, v in enumerate(boxes[holiday_hash(label)]):
                if v[0] == label:
                    boxes[holiday_hash(label)][i] = (label, value)
                    existing = True

            if not existing:
                boxes[holiday_hash(label)].append((label, value))

        else:
            label = step.split("-")[0]
            for i, v in enumerate(boxes[holiday_hash(label)]):
                if v[0] == label:
                    del boxes[holiday_hash(label)][i]
                    break

    return sum(
        [
            sum([(bi + 1) * (si + 1) * slot[1] for si, slot in enumerate(box)])
            for bi, box in enumerate(boxes)
        ]
    )


if __name__ == "__main__":
    print("Part 1:", sum([holiday_hash(step) for step in sequence]))
    print("Part 2:", HASHMAP(sequence))
