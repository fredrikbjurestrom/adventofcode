const fs = require('fs');

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

var array = fs.readFileSync("day01input.txt").toString().split("\n");


console.log("Part 1: " + array.filter(x => x !== '').reduce((acc, cur) => {return acc + calculateNeededFuel(cur)}, 0));
console.log("Part 2: " + array.filter(x => x !== '').reduce((acc, cur) => {return acc + calculateNeededFuelIteratively(cur)}, 0));