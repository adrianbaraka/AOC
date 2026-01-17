import { readFileSync } from "node:fs"


export function loadData() {
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
