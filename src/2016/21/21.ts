#!/usr/bin/env tsx

import {loadData, Permutations} from "@utils/utils.ts"


function swapPosition(s : string, x : number, y: number): string {
    let sx = s[x]
    let sy = s[y]

    let news = ""
    for (let i = 0; i < s.length; i ++) {
        if (i === x) {
            news += sy
            continue
        }

        if (i === y) {
            news += sx
            continue
        }

        news += s[i] 
    }
    return news
}

function swapLetter(s: string, x: string, y: string) : string {
    let news = ""
    for (let i = 0; i < s.length; i++) {
        if (s[i] === x) {
            news += y
            continue
        }
        if (s[i] === y) {
            news += x
            continue
        }
        news += s[i]
    }
    return news
}

function rotate(s: string, direction: string, steps: number) : string {
    const  sa = []

    for (let i = 0; i < s.length; i++) {
        let newPos = i
        if (direction === "right"){
            newPos = (i + steps) % s.length
        } else {
            newPos = i - steps
            if (newPos < 0) {
                newPos += s.length
            }
        }
        sa[newPos] = s[i]
    }
    return sa.join("")

}

function rotatePosition(s: string, x: string) :string{
    // find first position that x occurs
    const index = s.indexOf(x)
    //console.log(index)
    if (index < 0) {
        // Don't know what to do if not found
        return s
    }
    let times = 1 + index
    if (index >= 4) {
        times ++
    }

    return rotate(s, "right", times)
}

function reverse(s: string, x: number, y: number) :string{
    //console.log(`Received ${s}, ${x}, ${y}`)
    let xy = s.substring(x, y+1)
    let final = s.substring(0, x)

    for (let i = xy.length-1; i >= 0; i--) {
        final += xy[i]
    }
    final += s.substring(y+1)

    return final
}

function move(s: string, x: number, y: number) :string {
    // get element at pos x
    const elem = s[x]
    let str = ""
    let next = false
    for(let i = 0; i < s.length; i++) {
        if (i === x) {
            next = true
        }
        if (i === y) {
            next = false
            str += elem
            if (x < y) {
                continue
            }
        }
        if (next) {
            str += s.charAt(i+1)
        } else {
            str += s.charAt(i)
        }
    }
    return str

    //console.log(str)

}

function scramble(str : string, instructions: string[]) :string{
    for (let i = 0; i < instructions.length; i++) {
        const inst = instructions[i]
        // case 1
        let m = inst.match(/^swap position (\d+) with position (\d+)$/)
        if (m) {
            str = swapPosition(str, Number(m[1]), Number(m[2]))
            continue
        }

        // case 2
        m = inst.match(/^swap letter (\w+) with letter (\w+)$/)
        if (m) {
            str = swapLetter(str, m[1], m[2])
            continue
        }

        // case 3
        m = inst.match(/reverse positions (\d+) through (\d+)/)
        if (m) {
            str = reverse(str, Number(m[1]), Number(m[2]))
            continue
        }

        // case 4
        m = inst.match(/rotate (left|right) (\d+) step/)
        if (m) {
            str = rotate(str, m[1], Number(m[2]))
            continue
        }

        // case 5
        m = inst.match(/move position (\d+) to position (\d+)/)
        if (m) {
            str = move(str, Number(m[1]), Number(m[2]))
            continue
        }

        // case 6
        m = inst.match(/rotate based on position of letter (\w+)/)
        if (m) {
            str = rotatePosition(str, m[1])
            continue
        }

        // if reached here unsupported regex
        console.error("Unsupported line ", + inst)
        break
    }
    return str
    //console.log(str)

}

function main() {
    const instructions = loadData()
    const str = "abcdefgh"
    const p1 = scramble(str, instructions)

    console.log(`Part 1: ${p1}`)

    // only 40320 permutations can just brute force it
    const perms = Permutations(str)
    const target = "fbgdceah"
    for (let i = 0; i < perms.length; i++) {
        let scr = scramble(perms[i], instructions)
        if (scr === target) {
            console.log(`Part 2: ${perms[i]}`)
            break
        }
    }

}

main()
