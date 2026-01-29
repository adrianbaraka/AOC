#!/usr/bin/env tsx

import { readFileSync } from "node:fs"

function load(): string[] {
    try {
        // Use file path from args, or 0 for stdin
        const input = process.argv[2] || 0; 
        const content = readFileSync(input, "utf-8");
        
        // Split by line first, take the first line, then split by comma
        const firstLine = content.split(/\r?\n/)[0];
        
        if (!firstLine) throw new Error("File is empty");
        
        return firstLine.split(', ').map(s => s.trim());
    } catch (err) {
        console.error("Failed to load input:", err instanceof Error ? err.message : err);
        process.exit(1);
    }
}

interface position {
	row: number
	column: number
	direction: 'N' | 'E' | 'S' | 'W'
}

function processInst(current: position, inst: string | undefined) : position {
	// handle undefined case
	if (!inst) {
		console.error(`Undefined instruction. ${inst}`)
		return current
	}

	const dir = inst[0]
	const steps = Number(inst.slice(1)) // slice from the second element till the end
	switch (current.direction) {
		case 'N':
			if (dir === 'R') {
				current.direction = 'E'
				current.column += steps
			} else {
				current.direction = 'W'
				current.column -= steps
			}
			break;
		case 'E':
			if (dir === 'L') {
				current.direction = 'N'
				current.row += steps
			} else {
				current.direction = 'S'
				current.row -= steps
			}
			break;
		case 'S':
			if (dir === 'L') {
				current.direction = 'E'
				current.column += steps
			} else {
				current.direction = 'W'
				current.column -= steps
			}
			break;
		case "W":
			if (dir === 'R') {
				current.direction = 'N'
				current.row += steps
			} else {
				current.direction = 'S'
				current.row -= steps
			}
			break;

		default:
			console.error(`Unsupported instruction. ${inst}`)
			break;
	}
	//console.debug( inst, current, (Math.abs(current.column) + Math.abs(current.row)))
	return current

}

function main() {
	const data = load()

	let current: position = {row: 0, column: 0, direction: 'N'}

	// set holding all visited points
	const visited = new Set(["0,0"])

	// using anonymous functions
	const calcDist = function() : number {
		// coz starting is 0, 0 just take the abs value of both and add
		return Math.abs(current.column) + Math.abs(current.row)
	}

	// if part 2
	let part2 = false


	//console.log(current)

	for (let i = 0; i < data.length; i++) {
		//current = processInst(current, data[i])
		let inst = data[i]
		if (!inst) {
			console.error("Empty instruction")
			continue
		}

		const dir = inst[0]
		if (!dir) {
			console.error("Empty instruction")
			continue
		}
		const steps = Number(inst.slice(1)) // slice from the second element till the end
		for (let j = 0; j< steps; j++) {
			let face = current.direction
			current = processInst(current, `${dir}1`)
			if (j != steps -1 ){
				current.direction = face
			}

			// if current in visited we've found part 2
			if ( ! part2 && visited.has(`${current.row},${current.column}`) ) {
				console.log(`Part 2: ${calcDist()}`)
				part2 = true
			} else {
				visited.add(`${current.row},${current.column}`)
			}
			//console.debug(visited)
		}
		
	}
	//console.log(current)
	console.log(`Part 1: ${calcDist()}`)
	

}

main()