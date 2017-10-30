# Unique Hash String

Inspired by [Enigma](https://en.wikipedia.org/wiki/Enigma_machine), this is a variant of base64 algorithm targeting encode integers, after encoding every 6bits, the rotor (base64 table) shifts by the summary of all encoded 6bits modulus 64, to make sure a tiny difference could affect every byte of the calculated hash string.

## Run Code

- C++

  ```bash
  g++ -Wall -O2 -std=c++11 -o UniqueHashString UniqueHashString.cpp
  ./UniqueHashString
  ```

- golang

  ```bash
  go build UniqueHashString.go
  ./UniqueHashString
  ```

## Validation

It is impossible to validate all the possible input, for example 64bit integers. The encode and decode functions are tested for several large range of integers, there is no collision (decode error) at the moment.

- `[0, 2^32)`
- `[16345678912345678900, 16345678912345678900 + 2^32)`

## More Optimization

- The base64 table could be further obfuscated, it is base64url table after 10,000 random swap.
- Reverse base64 was using unordered map, array is around 3-4x faster than map.
- The encode() of golang was very slow, the profiling shows half of time was spend on runtime.mallocgc(), it is running twice as faster by using array with length.

## Samples

```text
Integer              -> encode()    -> decode()
16345678912345678900 -> M6aZ1qFgD1J -> 16345678912345678900
16345678912345678901 -> kLe0wbhTtwN -> 16345678912345678901
16345678912345678902 -> HUXGAnfV3Az -> 16345678912345678902
16345678912345678903 -> x-8aMr19qM7 -> 16345678912345678903
16345678912345678904 -> QBpekdwZbk4 -> 16345678912345678904
16345678912345678905 -> _CuXHFA0nH6 -> 16345678912345678905
16345678912345678906 -> 2cR8xhMGrxL -> 16345678912345678906
16345678912345678907 -> joKpQfkadQU -> 16345678912345678907
16345678912345678908 -> YlOu_1HeF_- -> 16345678912345678908
16345678912345678909 -> mI5R2wxXh2B -> 16345678912345678909
16345678912345678910 -> SyPKjAQ8fjC -> 16345678912345678910
16345678912345678911 -> JgEOYM_p1Yc -> 16345678912345678911
16345678912345678912 -> NVsPSHjRASl -> 16345678912345678912
16345678912345678913 -> z9vEJxYKMJI -> 16345678912345678913
16345678912345678914 -> 7ZWiNQmOkNy -> 16345678912345678914
16345678912345678915 -> 40Dsz_S5Hzg -> 16345678912345678915
16345678912345678916 -> 6Gtv72JPx7T -> 16345678912345678916
16345678912345678917 -> La3W4jNEQ4V -> 16345678912345678917
16345678912345678918 -> UeqD6Yzi_69 -> 16345678912345678918
16345678912345678919 -> -XbtLm7s2LZ -> 16345678912345678919
16345678912345678920 -> B8n3US4vjU0 -> 16345678912345678920
16345678912345678921 -> Cprq-J6WY-G -> 16345678912345678921
16345678912345678922 -> cudbBNLDmBa -> 16345678912345678922
16345678912345678923 -> oRFnCzUtSCe -> 16345678912345678923
16345678912345678924 -> lKhrc7-3JcX -> 16345678912345678924
16345678912345678925 -> IOfdo4BqNo8 -> 16345678912345678925
16345678912345678926 -> y51Fl6Cbzlp -> 16345678912345678926
16345678912345678927 -> gPwhILcn7Iu -> 16345678912345678927
16345678912345678928 -> TEAfyUor4yR -> 16345678912345678928
16345678912345678929 -> ViM1g-ld6gK -> 16345678912345678929
16345678912345678930 -> 9skwTBIFLTO -> 16345678912345678930
```
