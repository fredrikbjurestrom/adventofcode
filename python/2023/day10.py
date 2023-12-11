import matplotlib.pyplot as plt
from matplotlib.patches import Polygon
from pathlib import Path
import numpy as np

N = (-1, 0)
S = (1, 0)
E = (0, 1)
W = (0, -1)

symbol_map = {
    "|": [N, S],
    "-": [W, E],
    "L": [N, E],
    "J": [N, W],
    "7": [S, W],
    "F": [S, E],
}


graph = {}
start_node = (0, 0)
ground_nodes = []

with open(Path(__file__).parent / "inputs/day10.txt") as file:
    grid = np.array([[c for c in row] for row in file.read().splitlines()])

    for y, row in enumerate(grid):
        for x, symbol in enumerate(row):
            if symbol == "S":
                start_node = (y, x)
                graph[start_node] = []
            elif symbol == ".":
                ground_nodes.append((y, x))
            elif symbol in symbol_map:
                graph[(y, x)] = [
                    tuple(np.add((y, x), pipe)) for pipe in symbol_map[symbol]
                ]

    for node, links in graph.items():
        for link in links:
            if link == start_node:
                graph[link].append(node)


def part1() -> list[tuple[int, int]]:
    """Traverse the pipes and return the visited nodes. The length is used in part1, the path is used in part2"""
    visited = [start_node]
    next = graph[start_node][0]
    while next != start_node:
        visited.append(next)
        next = graph[next][0] if graph[next][0] != visited[-2] else graph[next][1]

    return visited


def part2(pipes) -> int:
    """Plot the path and count the number of points inside the polygon"""
    pipes.append(start_node)
    xy = np.array([(pipe[1], pipe[0]) for pipe in pipes])
    p = Polygon(xy, closed=True, edgecolor="darkblue", facecolor="lightskyblue")
    edges = set(pipes)
    contained = 0
    for y, x in np.ndindex(grid.shape):
        if (y, x) in edges:
            continue
        if p.contains_point((x, y)):
            contained += 1

    _, ax = plt.subplots()
    ax.add_patch(p)
    ax.set_xlim([0, grid.shape[1] - 1])
    ax.set_ylim([grid.shape[0] - 1, 0])

    plt.show()

    return contained


if __name__ == "__main__":
    pipes = part1()
    print("Part 1:", len(pipes) // 2)
    print("Part 2:", part2(pipes))
