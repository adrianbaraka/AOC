#!/usr/bin/env bun

import { loadData } from '@utils/utils.ts'
const data = loadData()

if (!data) {
    console.error('No data available')
    process.exit(1)
}



const width = data[0]?.length ?? 0
let message = ""
let message2 = ""

for (let i = 0; i < width; i++) {
    const freq = new Map()
    for (const line of data) {
        const char = line[i]
        freq.set(char, (freq.get(char) ?? 0) + 1)
    }
    let mostFreq = ""
    let countmostFreq = 0
    let leastFreq = ""
    let countLeastFreq = Infinity
    for (const key of freq.keys()) {
        //console.log(key)
        if (freq.get(key) > countmostFreq) {
            mostFreq = key
            countmostFreq = freq.get(key)
        }

        if (freq.get(key) < countLeastFreq) {
            leastFreq = key
            countLeastFreq = freq.get(key)
        }
    }
    //console.log(mostFreq)
    message += mostFreq
    message2 += leastFreq

}

console.log(`Part 1: ${message}`)
console.log(`Part 2: ${message2}`)