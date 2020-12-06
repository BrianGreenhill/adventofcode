import * as rd from 'readline';
import * as fs from 'file-system';

function readData(fileName: string): {
  reader: rd.Interface,
  data: Array<{range: number[]; letter: string; password: string}>
}
{
  const data: Array<{range: number[]; letter: string; password: string}> = [];
  const reader = rd.createInterface(fs.createReadStream(fileName));
  reader.on("line", (l: string) => {
    const tokens = l.split(' ')
    // prepare the range in form: "1-3" to [1, 3]
    const range = tokens[0].split('-').map(Number)
    // prepare the letter by removing the colon "a:" to "a"
    const letter = tokens[1].replace(":", "")
    // grab the password, no changes necessary
    const password = tokens[2]
    // add the elements to the data array
    data.push({
      range: range,
      letter: letter,
      password: password
    });
  })
  return {reader: reader, data: data}
}

function part1(fileName: string) {
  let validPasswords = 0
  const input = readData(fileName)

  input.reader.on("close", () => {
    for(const element of input.data) {
      const occurrences = (element.password.match(new RegExp(element.letter, "g")) || []).length
      if (occurrences < element.range[0] || occurrences > element.range[1]) {
        continue
      }
      validPasswords++
    }
      console.log("Part 1: ", validPasswords)
  })
}

function part2(fileName: string) {
  let validPasswords = 0
  const input = readData(fileName)

  input.reader.on("close", () => {
    for(const element of input.data) {
      const min = element.range[0]
      const max = element.range[1]
      const chars = element.password.split("")
      let atPosition = 0
      if (chars[min-1] == element.letter) {
        atPosition++
      }
      if (chars[max-1] == element.letter) {
        atPosition++
      }
      if (atPosition == 1) {
        validPasswords++
      }
      continue
    }
    console.log("Part 2: ", validPasswords)
  })
}

console.log("Advent of Code Day 2")
part1("day2/data")
part2("day2/data")
