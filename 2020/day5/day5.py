def convert_input_to_binary(x):
    if x == "R": return 1
    if x == "B": return 1
    return 0

seat_ids = []
with open("input") as f:
    input = f.readlines()

for val in input:
    binarylist = map(convert_input_to_binary, val.strip())
    row = int(''.join(str(e) for e in binarylist[:7]), 2)
    col = int(''.join(str(e) for e in binarylist[-3:]), 2)
    seat_id = row * 8 + col
    seat_ids.append(seat_id)

seat_ids.sort()
possibles = []
for i, seat_id in enumerate(seat_ids):
    if i == 0:
        continue
    if i == len(seat_ids) -1:
        break
    if seat_id + 1 != seat_ids[i+1]:
        possibles.append(seat_ids[i])
    if seat_id - 1 != seat_ids[i-1]:
        possibles.append(seat_ids[i])

if possibles[0] + 1 in seat_ids:
    print("wrong!")
    exit

ans = possibles[0] + 1
print("{} is your seat number".format(ans))
