#!/usr/bin/env bun
import { assertExists, loadData, toInt } from "@utils/utils.ts";

function main() {
	const data = loadData();
	const regex =
		/^(\w+) (inc|dec) (-?\d+) if (\w+) (>|<|<=|>=|==|!=) (-?\d+)$/;

	// map of the registers
	const registers: Map<string, number> = new Map();
	//Checks if the register exists in the map if it does not initialize it to 0. Else nothing is done.
	const initCheckregister = function (register: string) {
		if (!registers.has(register)) {
			registers.set(register, 0);
		}
	};

	let p2 = Number.MIN_VALUE;

	for (let i = 0; i < data.length; i++) {
		const element = data[i];
		const match = regex.exec(element);
		assertExists(match);

		// init ? check
		initCheckregister(match[1]);
		initCheckregister(match[4]);

		// get the value of first
		const firstRegister = registers.get(match[1]);
		assertExists(firstRegister);

		// get the value of the change
		let change = toInt(match[3]);
		if (match[2] === "dec") {
			change = -change;
		}

		// get value of the if register
		const ifregister = registers.get(match[4]);
		assertExists(ifregister);

		// first use case ever of eval
		const toEval = `${ifregister} ${match[5]} ${match[6]}`;
		const shouldEvaluate: boolean = eval(toEval);

		if (shouldEvaluate) {
			const newVal = firstRegister + change;
			registers.set(match[1], newVal);

			// part 2
			if (newVal > p2) {
				p2 = newVal;
			}
		}
	}

	let max = Number.MIN_VALUE;
	for (const r of registers.values()) {
		if (r > max) {
			max = r;
		}
	}
	console.log("Part 1: " + max);
	console.log("Part 2: " + p2);
}

main();
