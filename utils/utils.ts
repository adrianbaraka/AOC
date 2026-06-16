import { readFileSync } from "node:fs";
import { createHash } from "node:crypto";

/**
 *
 * @returns An array containing all lines of input passed over stdin or as the first arguement.
 */
export function loadData(): string[] {
	try {
		const input = process.argv[2] || 0;
		const content = readFileSync(input, "utf-8").trim().split(/\r?\n/); //CRLF
		return content;
	} catch (error) {
		console.error(error);
		process.exit(1);
	}
}

// Compute the md5hash in hex of str
export function md5hash(str: string): string {
	const hash = createHash("md5");
	hash.update(str);
	const digest = hash.digest("hex");
	return digest;
}

/**
 *
 * @param n number
 * @returns Factorial of the number
 * @example factorial(5) === 5! === 120
 */
export function factorial(n: number): number {
	if (n <= 1) {
		return 1;
	} else {
		return n * factorial(n - 1);
	}
}

/**
 * Converts a string to a number.
 * Throws an error if the conversion fails.
 */
export function toInt(str: String): number {
	const num = Number(str);
	if (Number.isNaN(num)) {
		throw new Error(`Cannot convert "${str}" to a number.`);
		//console.error(`Cannot convert "${str}" to a number.`);
		//process.exit(1);
	}
	return num;
}

//https://www.geeksforgeeks.org/javascript/javascript-program-to-print-all-permutations-of-given-string/
// TODO are there interators in js?
export function Permutations(str: string) {
	const permutations = [];
	//const factorial = n => n <= 1 ? 1 : n * factorial(n - 1);

	const len = str.length;
	const totalPermutations = factorial(len);

	for (let i = 0; i < totalPermutations; i++) {
		let currentPermutation = "";
		let temp = i;

		const availableChars = str.split("");

		for (let j = len - 1; j >= 0; j--) {
			const index = temp % (j + 1);
			temp = Math.floor(temp / (j + 1));

			currentPermutation += availableChars[index];
			availableChars.splice(index, 1);
		}

		permutations.push(currentPermutation);
	}

	return permutations;
}

/**
 * Asserts that a value is neither null nor undefined.
 * @param val - The value to check.
 * @param message - Optional custom error message.
 * @throws {Error} If val is null or undefined.
 */
export function assertExists<T>(
	val: T,
	message: string = "Value is null or undefined",
): asserts val is NonNullable<T> {
	if (val == null) {
		throw new Error(message);
	}
}

/**
 *
 * @param str
 * @returns An array of the string in bytes
 */
export function stringToBytes(str: string): number[] {
	const res = [];
	for (let i = 0; i < str.length; i++) {
		res.push(str.charCodeAt(i));
	}
	return res;
}
