'use strict'

class IntcodeProgram {
    constructor(memory) {
        this._memory = memory.slice();
    }

    init(memory) {
        this._memory = memory.slice();
    }

    run(address) {
        let instruction = this._memory[address];
        
        if (instruction === 99 || instruction == null)
        {
            return;
        }
        else if (instruction !== 1 && instruction !== 2) {
            this.run(address+1);
        }

        let para1address = this._memory[address+1];
        let para2address = this._memory[address+2];
        let para3address = this._memory[address+3];

        if (instruction === 1) {
            this._memory[para3address] = this._memory[para1address] + this._memory[para2address];
        }
        else if (instruction === 2) {
            this._memory[para3address] = this._memory[para1address] * this._memory[para2address];
        }

        this.run(address + 4);
    }

    get(address)
    {
        return this._memory[address];
    }

    set(address, value)
    {        
        this._memory[address] = value;
    }
};

module.exports = IntcodeProgram;