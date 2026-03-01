#!/usr/bin/env bun

import { loadData, toInt } from "@utils/utils.ts";

function run(dataInt: number[], part2: boolean = false): number {
	// create a new copy of dataInt as it is modified here
	let data = [...dataInt];
	let k = 0;
	let steps = 0;
	while (k < data.length) {
		const elem = data[k];
		if (part2 && elem >= 3) {
			data[k] = elem - 1;
		} else {
			data[k] = elem + 1;
		}
		k += elem;
		steps++;
		//console.log(dataInt);
	}
	//console.log(dataInt);
	return steps;
}

function main() {
	const data = loadData();
	const dataInt: number[] = [];

	for (let i = 0; i < data.length; i++) {
		dataInt[i] = toInt(data[i]);
	}

	const p1 = run(dataInt);
	console.log("Part 1: " + p1);

	const p2 = run(dataInt, true);
	console.log("Part 2: " + p2);
}

main();
