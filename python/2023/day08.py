from pathlib import Path
import math

with open(Path(__file__).parent / "inputs/day08.txt") as file:
    instructions = file.readline().strip()

    network = {}
    for row in filter(None, map(str.strip, file)):
        node, links = row.split(" = ")
        network[node] = tuple(links.strip("()").split(", "))

instruction_map = {
    "L": 0,
    "R": 1,
}


def part1(start, end, max_steps=None) -> int:
    """
    Returns the number of steps required to reach the end node from the start node.
    """
    steps = 0
    position = start
    while True:
        i = steps % len(instructions)
        position = network[position][instruction_map[instructions[i]]]

        if max_steps and steps >= max_steps:
            return None

        steps += 1
        if position == end:
            return steps


def part2(start_positions, end_positions) -> int:
    """
    Returns the lcm (least common multiple) number of steps required to reach the first end node for all start nodes.
    part1() is called with a max_steps parameter to prevent looking to deep, we only care about the closest end node.
    """
    steps = []
    for start in start_positions:
        for end in end_positions:
            step_count = part1(start, end, 100_000)
            if step_count:
                steps.append(step_count)

    return math.lcm(*steps)


if __name__ == "__main__":
    print("Part 1:", part1("AAA", "ZZZ"))

    start_positions = list(filter(lambda x: x[2] == "A", network.keys()))
    end_positions = list(filter(lambda x: x[2] == "Z", network.keys()))
    print("Part 2:", part2(start_positions, end_positions))
