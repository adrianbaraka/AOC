import hashlib

target = '000000'
key = 'iwrupvqb'
candidate = 0

while True:
    plaintext = f"{key}{candidate}"
    hash = hashlib.md5(plaintext.encode('ascii')).hexdigest()
    if hash[:6] == target:
        print(f"Num found: {candidate}")
        break
    candidate = candidate + 1
