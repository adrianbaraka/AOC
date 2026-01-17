#!/usr/bin/env tsx

import { loadData } from '../../../utils/utils.ts'
import { readFileSync } from 'node:fs'
const data = loadData()

interface current {
    row: number
    column: number
}

function handleDir(curr: current, dir: string, len: number, cols: number) :[current, number]{
    let tcurr = {... curr} // apparently passing this to the function is by reference
    switch (dir) {
        case 'U':
            tcurr.row -= 1
            break;
        case 'R':
            tcurr.column += 1
            break
        case 'D':
            tcurr.row += 1
            break
        case 'L':
            tcurr.column -= 1
            break
        default:
            console.error(`Unsupported function. ${dir}`)
            break;
    }
    // check if current is valid
    if (tcurr.row >= 0 && tcurr.row < (len / cols) && tcurr.column >= 0 && tcurr.column < cols) {
        let idx = (tcurr.row * cols) + tcurr.column
        if (idx >= 0 && idx < len) {
            //idx ++
            //console.log("here ", tcurr)
            return [tcurr, idx]
        }
    }
    // if fails
    //console.log("failed ", original)
    return [curr, (curr.row * cols) + curr.column]
}

function getKeypad() : string[]{
    // create an arrray of the keypad
    let keypad : string[] = []
    try {
        const cont = readFileSync("keypad.text", "utf-8").trim().split(/\n/)
        for(let i=0; i < cont.length; i++) {
            const line = cont[i]
            if (!line) {
                continue
            }
            for (let j=0; j< line?.length; j++) {
                let char = line[j]
                if (!char) {
                    continue
                }
                keypad.push(char)
                
            }

        }
        //console.log(keypad)
        return keypad
        //console.log(cont)
    } catch (error) {
        console.error(error)
        process.exit(1)
    }
}


function main() {
    const keypad = getKeypad()
    // both start at 5
    let cur : current = {row: 1, column: 1}
    let cur2 : current = {row: 2, column: 0} 
    let pass = ""
    let pass2 = ""

    for (const line of data) {
        if (!line) continue
            for (let i = 0; i < line.length; i++) {
                const dir = line[i]
                if (!dir) continue

                // part 1
                const [next, idx] = handleDir(cur, dir, 9, 3)
                cur = next
                if (i === line.length - 1) {
                    pass += `${idx+1}`
                }

                // part 2
                const old = (cur2.row * 5) + cur2.column
                let [next2, idx2] = handleDir(cur2, dir, 25, 5)
                // in keypad.txt replace blanks with the '-' character easy visibility
                if (keypad[idx2] !== '-') {
                    cur2 = next2
                } else {
                    idx2 = old
                }

                if (i === line.length -1) {
                    pass2 += `${keypad[idx2]}`
                }
            }
    }

    console.log("Part 1: " + pass)
    console.log("Part 2: " + pass2)
}

main()

