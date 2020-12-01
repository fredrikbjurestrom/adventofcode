const fs = require('fs');

var input = fs.readFileSync("./inputs/day01input.txt").toString().split("\n");;

const calculateNeededFuel = (mass) => {
  return Math.floor(mass/3)-2;
};

const calculateNeededFuelIteratively = (mass) => {
    var totalFuel = 0;
    var fuel = calculateNeededFuel(mass);
    while(fuel > 0)
    {
        totalFuel += fuel;
        fuel = calculateNeededFuel(fuel);
    }
    return totalFuel;
}

const partOne = () => {
  let input = fs.readFileSync("day01input.txt").toString().split("\n");
  return input.filter(x => x !== '').reduce((acc, cur) => {return acc + calculateNeededFuel(cur)}, 0);
}

const partTwo = () => {
  let input = fs.readFileSync("day01input.txt").toString().split("\n");
  return input.filter(x => x !== '').reduce((acc, cur) => {return acc + calculateNeededFuelIteratively(cur)}, 0);
}



console.log("Part 1: " + input.filter(x => x !== '').reduce((acc, cur) => {return acc + calculateNeededFuel(cur)}, 0));
console.log("Part 2: " + input.filter(x => x !== '').reduce((acc, cur) => {return acc + calculateNeededFuelIteratively(cur)}, 0));
