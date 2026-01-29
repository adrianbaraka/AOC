#!/usr/bin/env tsx

import {loadData} from "@utils/utils.ts"

function expand(str: string) : string {
    let b = ""
    // reverse str and replace
    for (let i = str.length - 1; i >= 0; i--) {
        if (str[i] === "0") {
            b += "1"
        } else {
            b += "0"
        }
        //b += str[i]
    }
    return str + "0" + b

}

function calcChecksum(str: string) :string {
    while (true) {
        let checksum = ""
        for (let i = 0; i < str.length - 1; i += 2) {
            if (str[i] === str[i+1]) {
                checksum += "1"
            } else {
                checksum += "0"
            }
        }
        if (checksum.length % 2 === 0) {
            str = checksum
        } else {
            //console.log("here", checksum)
            return checksum
        }
    }
}

function run(str: string, requiredLen: number) {
    // expand the string
    while (str.length < requiredLen) {
        str = expand(str)
    }
    // calculate checksum of first n chars
    const checksum = calcChecksum(str.slice(0, requiredLen))

    return checksum
}

function main() {
    const data = loadData()

    if (data.length < 1) {
        console.error("Incorrect input")
        process.exit(1)
    }
    const start = data[0]
    const len1 = 272
    const len2 = 35651584

    // part 1
    const p1 = run(start, len1)
    console.log(`Part 1: ${p1}`)

    // part 2
    const p2 = run(start, len2)
    console.log(`Part 2: ${p2}`)
}

main()


