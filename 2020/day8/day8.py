import sys
import fileinput
from copy import deepcopy

def exec_command(P, ip, accumulator):
    words = P[ip]
    if words[0] == 'acc':
        accumulator += int(words[1])
        ip += 1
    if words[0] == 'nop':
        ip += 1
    if words[0] == 'jmp':
        ip += int(words[1])
    return (ip, accumulator)

accumulator = 0
ip = 0
seen = set()
lines = list([l.split() for l in fileinput.input()])
while True:
    if ip in seen:
        print("Part 1: {}".format(accumulator))
        break
    seen.add(ip)
    ip, accumulator = exec_command(lines, ip, accumulator)

for change in range(len(lines)):
    instructions = deepcopy(lines)
    if instructions[change][0] == 'nop':
        instructions[change][0] = 'jmp'
    elif instructions[change][0] == 'jmp':
        instructions[change][0] = 'nop'
    else:
        continue
    t = 0
    ip = 0
    acc = 0
    while 0<=ip<len(instructions) and t<1000:
        t += 1
        ip, acc = exec_command(instructions, ip, acc)
    if ip == len(instructions):
        print("Part 2: {}".format(acc))
