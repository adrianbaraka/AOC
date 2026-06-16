#!/usr/bin/env bun

import { loadData } from "@utils/utils.ts";

const data = loadData()[0];

let i = 0;
let curr = 0;
let score = 0;
let garbage = false;
let garbageChars = 0;

while (i < data.length) {
	const elem = data[i];
	// console.log(elem);
	// console.log(curr)

	// skip
	if (elem === "!") {
		i += 2;
		continue;
	}

	// begin garbage
	if (elem === "<" && !garbage) {
		garbage = true;
		i++;
		continue;
	}

	// end garbage
	if (garbage && elem === ">") {
		garbage = false;
		i++;
		continue;
	}

	// within garbage
	if (garbage) {
		garbageChars++;
		i++;
		continue;
	}

	// group opening
	if (elem === "{") {
		curr++;
		i++;
		continue;
	}

	// group closing
	if (elem === "}") {
		score += curr;
		curr--;
		i++;
		continue;
	}

	// within a group
	i++;
}

console.log(`Part 1: ${score}`);
console.log(`Part 2: ${garbageChars}`);
