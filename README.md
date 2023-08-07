# hashtrace
_HashTrace_ is a utility designed to guess the hash based on provided content fragments.

## Description
Sometimes, we're handed a piece of the puzzle: a partial hash or scattered content fragments without context.
HashTrace bridges this information gap, helping to derive the potential plaintext that these fragments might represent.

### Use Case Example: 
Imagine a Secret: the plaintext is derived from the command:
```bash
echo "john|102|john@example.xx" | md5sum
```
This produces the hash `e84be6efffb4b33f8dabbc725e7e1d78`.

Given known input data: `102`, `John`, and` john@example.xx` (along with the hash or its fragment), 
HashTrace assists in unveiling the mystery behind the order and format in which these data segments were originally hashed.

## Usage
```bash
./hashtrace -hash='e6efffb4b33f8' -data=102 -data=John --data=john@example.xx
```

```bash
................................................................................
>> e6efffb4b33f8
>> 102,John,john@example.xx
>> Separators ["," "@" "." "|" ";" ""]
................................................................................
2023/08/07 14:29:25   2895 needles for ["102" "John" "john@example.xx"]
=========================================================================
2023/08/07 14:29:25 FOUND AT 2023-08-07 14:29:25.974850844 +0300 EEST m=+0.043376867
PLAIN:  "john|102|john@example.xx\n"
HASH:    e6efffb4b33f8
FULL:    e84be6efffb4b33f8dabbc725e7e1d78
=========================================================================

```

### Supported Hash Algorithms
- `MD5`
- `SHA1`
- `SHA256`
- `SHA512`
- `SHA224`
- `SHA384`
- `SHA3-256`
- `SHA3-512`
- `SHA512-224`
- `SHA512-256`