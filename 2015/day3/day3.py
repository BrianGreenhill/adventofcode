import fileinput

def get_visited(directions):
    x = 0
    y = 0
    VISITED = set()
    for i in range(len(directions)):
        if directions[i] == '>':
            x += 1
        if directions[i] == '<':
            x -= 1
        if directions[i] == '^':
            y += 1
        if directions[i] == 'v':
            y -= 1
        coords = (x,y)
        if coords not in VISITED:
            VISITED.add(coords)
    return VISITED

p1 = 0
p2 = 0

lines = [list(l) for l in fileinput.input()][0]

#part 1
visited = get_visited(lines)

#part 2
santa = get_visited(lines[::2])
robosanta = get_visited(lines[1::2])

p1 = len(visited) + 1
p2 = len(santa.union(robosanta))

print(p1)
print(p2)
