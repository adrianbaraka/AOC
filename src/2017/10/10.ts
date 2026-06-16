#!/usr/bin/env bun
import { assertExists, loadData, toInt, stringToBytes } from "@utils/utils.ts";

function round(
	lengths: string[],
	elements: number[],
	currentPosition: number = 0,
	skipSize: number = 0,
) {
	//console.log("current position = ", currentPosition, "skip size  = ", skipSize);

	const numElements = elements.length;
	for (let i = 0; i < lengths.length; i++) {
		//console.log(elements[currentPosition])
		const num = toInt(lengths[i]);

		// reversed elements
		const affected: Map<number, number> = new Map();
		let iter = 0;
		for (let j = num - 1; j >= 0; j--) {
			const index = (currentPosition + j) % numElements;
			//console.log(currentPosition + iter);
			affected.set(
				(currentPosition + iter) % numElements,
				elements[index],
			);
			iter++;
		}
		//console.log(affected);
		//const newElements: number[] = [];
		for (let k = 0; k < numElements; k++) {
			if (affected.has(k)) {
				const val = affected.get(k);
				assertExists(val);
				elements[k] = val;
			}
		}
		currentPosition = (currentPosition + num + skipSize) % numElements;
		skipSize++;
		//console.log(elements)
	}
	//console.log(elements[0] * elements[1]);
	return { currentPosition, skipSize };
}

function part1(data: string, elements: number[]) {
	const d1 = data.split(/,\s?/);
	round(d1, elements);
	console.log(`Part 1: ${elements[0] * elements[1]}`);
}

function part2(data: string, elements: number[]) {
	const toAdd = ["17", "31", "73", "47", "23"];
	const d2: string[] = [];

	//let q = 0
	stringToBytes(data).forEach((byte) => {
		//console.log(byte)
		//q++
		d2.push(byte.toString());
	});
	//console.log(q);

	d2.push(...toAdd);
	//console.log(data2.length);

	let rounds = 64;
	let curr = 0;
	let skip = 0;

	// 64 rounds
	while (rounds > 0) {
		let { currentPosition, skipSize } = round(d2, elements, curr, skip);

		curr = currentPosition;
		skip = skipSize;
		// elements is modified as it is passed by reference
		// elems = elements;

		rounds--;
	}
	//console.log(elems)
	//const d2 = [65, 27, 9,1, 4, 3, 40, 50 ,91,7, 6, 0, 2,5, 68, 22]
	//step3(d2)
	printHash(elements);
}

function printHash(sparseHash: number[]) {
	if (sparseHash.length != 256)
		throw new Error("Sparse hash is not 256 chars");

	let hash = "";
	for (let i = 0; i < sparseHash.length; i += 16) {
		let num = sparseHash[i];
		for (let j = i + 1; j < i + 16; j++) {
			num = num ^ sparseHash[j];
			//console.log(sparseHash[j]);
		}
		//console.log(num.toString(16))
		hash += num.toString(16).padStart(2, "0");
	}
	if (hash.length != 32) throw new Error ("Error hash is not 32 chars")

	console.log(`Part 2: ${hash}`);
}

function main() {
	const data = loadData()[0];
	const numElements = 256;
	let elems: number[] = [];
	for (let i = 0; i < numElements; i++) {
		elems.push(i);
	}

	part1(data, [...elems]);
	part2(data, [...elems]);
}

main();
