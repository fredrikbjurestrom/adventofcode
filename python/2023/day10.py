from pathlib import Path
import time
import numpy as np
import sys

symbol_map = {
    "|": [(-1, 0), (1, 0)],
    "-": [(0, -1), (0, 1)],
    "L": [(-1, 0), (0, 1)],
    "J": [(-1, 0), (0, -1)],
    "7": [(1, 0), (0, -1)],
    "F": [(1, 0), (0, 1)],
    ".": [(-1, 0), (1, 0), (0, -1), (0, 1)],
}


nodes = {}
start_node = (0, 0)
ground_nodes = []

with open(Path(__file__).parent / "inputs/day10_example.txt") as file:
    for y, row in enumerate(file.read().splitlines()):
        for x, symbol in enumerate(row):
            if symbol == "S":
                start_node = (y, x)
                nodes[start_node] = ("S", [])
            elif symbol == ".":
                ground_nodes.append((y, x))
            elif symbol in symbol_map:
                pipes = [tuple(np.add((y, x), pipe)) for pipe in symbol_map[symbol]]
                nodes[(y, x)] = (symbol, pipes)

    for node, links in nodes.items():
        for link in links[1]:
            if link == start_node:
                nodes[link][1].append(node)

    for ground_node in ground_nodes:
        nodes[ground_node] = ".", [
            tuple(np.add(ground_node, pipe)) for pipe in symbol_map["."]
        ]

print(start_node, nodes[start_node], nodes[0, 0])


def _can_escape(graph, node) -> bool:
    visited = set()
    queue = []
    visited.add(node)
    queue.append([node, node])

    while queue:
        test = queue.pop(0)
        print("poped", test)

        prev, cur = test
        print("pop", prev, cur, queue)
        tile, neigbours = graph[cur]

        delta = (0, 0) if not prev else tuple(np.subtract(cur, prev))
        print(prev, cur, delta)

        # if tile == "7" or tile == "F":
        #     neigbours.append(tuple(np.add(s, (-1, 0))))

        # if tile == "J" or tile == "L":
        #     neigbours.append(tuple(np.add(s, (1, 0))))

        for neigbour in neigbours:
            if neigbour not in graph:
                # print(node, "escaped")
                return True

            # print(tile, cur, neigbour, delta)

            if neigbour not in visited:
                visited.add(neigbour)
                print("append", cur, neigbour, queue)
                queue.append([cur, neigbour])
                print("appended queue", queue)

    print(node, "can't escape")
    return False


def part1() -> int:
    moves = 1

    next = nodes[start_node][1][0]
    prev = start_node
    while next != start_node:
        moves += 1
        prev, next = next, nodes[next][1][0] if nodes[next][1][0] != prev else nodes[next][1][1]

    return moves // 2


def part2() -> int:
    escapees = 0

    for ground_node in ground_nodes:
        if not _can_escape(nodes, ground_node):
            escapees += 1

    return escapees


if __name__ == "__main__":
    start = time.time()
    sys.setrecursionlimit(2000)
    print("Part 1:", part1())
    end = time.time()
    print("Took:", end - start)
    start = time.time()
    print("Part 2:", part2())
    end = time.time()
    print("Took:", end - start)
