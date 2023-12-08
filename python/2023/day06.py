from pathlib import Path
import numpy as np

with open(Path(__file__).parent / "inputs/day06.txt") as file:
    data = file.read().splitlines()


def part1() -> int:
    """Naïve approach, surely in for a surprise in part 2"""
    times = [int(d) for d in data[0].split(":")[1].split(" ") if d]
    distances = [int(d) for d in data[1].split(":")[1].split(" ") if d]

    ways_to_win = []
    for i in range(len(times)):    
        wins = 0
        for n in range(times[i]):
            travel_distance = n * (times[i]-n)
            
            if travel_distance > distances[i]:
                wins += 1
                
        ways_to_win.append(wins)
        
    return np.prod(np.array(ways_to_win))


def part2() -> int:
    """One race instead of many. Is this really day 6? Ok, it takes 10s, but I'm not going to bother."""
    time = int(data[0].split(":")[1].replace(" ",""))
    distance = int(data[1].split(":")[1].replace(" ",""))

    wins = 0
    for n in range(time):
        travel_distance = n * (time-n)

        if travel_distance > distance:
            wins += 1
            
    return wins


if __name__ == "__main__":
    print("Part 1:", part1())
    print("Part 2:", part2())
