#!/usr/bin/env bun

import { loadData, toInt } from "@utils/utils.ts";

function main() {
	const data = loadData();
	let checksum = 0;
	let checksum2 = 0;

	for (let i = 0; i < data.length; i++) {
		// split on whitespace
		let nums = data[i].trim().split(/\s+/);
		let max = Number.MIN_VALUE;
		let min = Number.MAX_VALUE;

		for (let j = 0; j < nums.length; j++) {
			const num = toInt(nums[j]);
			// part 1
			{
				if (num > max) {
					max = num;
				}
				if (num < min) {
					min = num;
				}
			}
			// part 2
			{
				for (let k = 0; k < nums.length; k++) {
					if (k === j) {
						continue;
					}
					const num2 = toInt(nums[k]);
					// check if it is divisible
					if (num % num2 === 0) {
						//console.debug(`num=${num} num2 = ${num2}`)
						checksum2 += num / num2;
					}
				}
			}
		}

		checksum += max - min;
	}
	console.log("Part 1: " + checksum);
	console.log("Part 2: " + checksum2);
}

main();
