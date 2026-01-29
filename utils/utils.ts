import { readFileSync } from "node:fs"
import { createHash } from "node:crypto"


export function loadData() : string[] {
    // Returns an array containing all lines of input passed over stdin or as the first arguement.
    try {
        const input = process.argv[2] || 0
        const content = readFileSync(input, "utf-8").trim().split(/\r?\n/) //CRLF
        return content
    } catch (error) {
        console.error(error)
        process.exit(1)
    }
}

// Compute the md5hash in hex of str
export function md5hash(str : string) :string {
    const hash = createHash("md5")
    hash.update(str)
    const digest = hash.digest("hex")
    return digest
}

export function factorial(n : number): number {
    if (n <= 1) {
        return 1
    } else {
        return n * factorial(n -1)
    }
}

//https://www.geeksforgeeks.org/javascript/javascript-program-to-print-all-permutations-of-given-string/
// TODO are there interators in js?
export function Permutations(str: string) {
    const permutations = [];
    //const factorial = n => n <= 1 ? 1 : n * factorial(n - 1);

    const len = str.length;
    const totalPermutations = factorial(len);

    for (let i = 0; i < totalPermutations; i++) {
        let currentPermutation = '';
        let temp = i;

        const availableChars = str.split('');

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
