'use strict'

const fs = require('fs');
const intcodeProgram = require('./utils/intcodeProgram');

const input = fs.readFileSync("./inputs/day02input.txt").toString().split(",").map(x => parseInt(x));

const operandResolver = (expected) => {
    for (let noun = 0; noun < 100; noun++) {
        for (let verb = 0; verb < 100; verb++) {
            intcodeProgram.init(input);
            intcodeProgram.set(1, noun);
            intcodeProgram.set(2, verb);
            
            intcodeProgram.run(0);
            
            if (intcodeProgram.get(0) === expected) {
                return { noun, verb };
            }
        }   
    }
};

const partOne = (input) => {
    intcodeProgram.init(input)
    
    // Revert to "1202 program alarm" state
    intcodeProgram.set(1, 12);
    intcodeProgram.set(2, 2);

    intcodeProgram.run(0);

    console.log("Part 1: " + intcodeProgram.get(0));
};

const partTwo = (expected) => {
    let result = operandResolver(expected);
    let reportValue = result.noun*100 + result.verb;
    console.log("Part 2: " + reportValue)
};


partOne(input);
partTwo(19690720);