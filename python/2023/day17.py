from pathlib import Path
from collections import defaultdict
from queue import PriorityQueue
from typing import Iterator, NewType
import sys

Point = NewType("Point", tuple[int, int])
Vector = NewType("Vector", list[int, int])

directions: list[Vector] = {
    "N": [-1, 0],
    "S": [1, 0],
    "E": [0, 1],
    "W": [0, -1],
}


def displace(pos, vec) -> tuple[int, int]:
    return tuple(sum(p) for p in zip(pos, vec))


class Graph:
    def __init__(self, fileName):
        self.weights: dict[Point, int] = {}
        self.target: Point = (0, 0)
        with open(Path(__file__).parent / f"inputs/{fileName}") as file:
            for y, row in enumerate(file.read().splitlines()):
                for x, heat_loss in enumerate(row):
                    self.weights[(y, x)] = int(heat_loss)

            self.target = (y, x)

    def neighbours(
        self, current: Point, current_dir: Vector, momentum: int, rules: tuple[int, int]
    ) -> Iterator[tuple[Point, Vector, int]]:
        min_moves, max_moves = rules
        for dir in directions.values():
            y, x = displace(current, dir)
            if (y, x) not in self.weights:
                continue
            elif displace(dir, current_dir) == (0, 0):
                continue
            elif current_dir[0] == dir[0] and current_dir[1] == dir[1]:
                if momentum == max_moves:
                    continue
                yield (y, x), dir, momentum + 1
            else:
                if min_moves <= momentum:
                    yield (y, x), dir, 1


def dijkstra(graph: Graph, start: Point, rules: tuple[int, int]):
    pq = PriorityQueue()
    dist = defaultdict(lambda: sys.maxsize)
    pq.put((0, (start, directions["E"], 0)))
    pq.put((0, (start, directions["S"], 0)))

    while not pq.empty():
        weight, (current, direction, momentum) = pq.get()

        for neighbour, new_direction, new_momentum in graph.neighbours(
            current, direction, momentum, rules
        ):
            new_weight = weight + graph.weights[neighbour]
            if new_weight < dist[(neighbour, current, new_momentum)]:
                dist[(neighbour, current, new_momentum)] = new_weight
                pq.put((new_weight, (neighbour, new_direction, new_momentum)))

    return min([v for d, v in dist.items() if d[0] == graph.target])


if __name__ == "__main__":
    graph = Graph("day17.txt")
    print("Part 1:", dijkstra(graph, (0, 0), (-1, 3)))
    print("Part 2:", dijkstra(graph, (0, 0), (4, 10)))
