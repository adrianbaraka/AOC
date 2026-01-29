#!/usr/bin/env tsx

import {loadData, md5hash} from "@utils/utils.ts"

const salt = loadData()[0]
const hashes: Map<string, string> = new Map() // store computed hashes here
const hashes2: Map<string, string> = new Map()

function getHash(str: string): string {
    if (hashes.has(str)) {
        const hs = hashes.get(str)
        if (hs) {
            return hs
        }
    }
    // compute
    const h = md5hash(str)
    hashes.set(str, h)
    return h
}

function getHash2(str: string): string {
    const original = str
    if (hashes2.has(str)) {
        const hs = hashes2.get(str)
        if (hs) {
            return hs
        }
    }

    // compute
    let i = 0
    while(true) {
        if (i > 2016) {
            hashes2.set(original, str)
            return str
            //break
        }
        const h = md5hash(str)
        //const h = getHash(str)
        i++
        str = h
    }

}

function contains3(str : string): [boolean, string]  {
    for (let i = 0; i <= str.length - 3; i++) {
        if (str[i] === str[i+1] && str[i+1] === str[i+2]) {
            const ret = str[i] + str[i+1] + str[i+2]
            return [true, ret]
        }

    }
    return [false, ""]
}

// takes the first char of search 5 times 
function contains5(str: string, search: string) : boolean {
    search = search[0].repeat(5)
    const found = str.match(search)

    if (!found) {
        return false
    }
    return true

}

function checkNext1000(i: number, char: string, p2: boolean): boolean {

    for (let j = 0; j < 1000; j++) {
        let idx = i + j
        let str = salt + idx
        if (p2) {
            if (contains5(getHash2(str), char)) {
                return true
            }

        } else {
            if (contains5(getHash(str), char)) {
                return true
            }

        }

    }
    return false
}

let i = 0
let keys = 0
let keys2 = 0
let part1 = false
let part2 = false

while (true) {
    //console.log("\r ", + i)
    if (i % 1000 == 0) {
        process.stdout.write(`\r ${i} `)
    }
    let str = salt + i
    if (!part1) {
        let [ok, char] = contains3(getHash(str))
        if (ok && checkNext1000(i+1, char, false)) {
            keys ++
            //console.log(`Found ${i} keys = ${keys}`)
            if (keys === 64) {
                console.log(`\nPart 1: ${i}`)
                part1 = true
            }
        }
    }

    // part 2
    if (!part2){
        let [ok2, char2] = contains3(getHash2(str))
        if (ok2 && checkNext1000(i+1, char2, true)) {
            keys2 ++

            if (keys2 === 64) {
                console.log("\nPart 2: ", + i)
                part2 = true
            }
        }
    }

    if (part1 && part2) {
        break
    }
    i++
    
}