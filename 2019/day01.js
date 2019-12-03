const fs = require('fs');

<<<<<<< HEAD
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
  return array.filter(x => x !== '').reduce((acc, cur) => {return acc + calculateNeededFuel(cur)}, 0);
}

const partTwo = () => {
  let input = fs.readFileSync("day01input.txt").toString().split("\n");
  return array.filter(x => x !== '').reduce((acc, cur) => {return acc + calculateNeededFuelIteratively(cur)}, 0);
}


console.log("Part 1: " + array.filter(x => x !== '').reduce((acc, cur) => {return acc + calculateNeededFuel(cur)}, 0));
console.log("Part 2: " + array.filter(x => x !== '').reduce((acc, cur) => {return acc + calculateNeededFuelIteratively(cur)}, 0));
=======
var array = fs.readFileSync("./inputs/day01input.txt").toString().split("\n");;
console.log(array.filter(x => x !== '').reduce((acc, cur) => {return acc+(Math.floor(cur/3)-2)}, 0));
>>>>>>> 59f0a81dca69b5ddfa653eae76804fe6b380d6b0
