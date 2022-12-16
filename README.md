# rpgforge-go
Simple dice roll simulator, translating "human readable" dice definition like "3d6" to set of rolls.

## Roll definitions
Default roll definition `3d6+1` means:

 Roll 3 times, 6 sides dice, add 1 to final sum.

 ### Exploding dice
 Exploding dice are when rolls the maximum on a die,  then can roll that die again and the new die roll result is added to the original. 

 `3d6!` means - roll 3 times, 6 sides dice, reroll every "6" adding new result to final sum.

## Usage
Simple throw.
```
result := ThrowDices("3d6+1").Sum()
```
## Initialization
To get more random results, use the Initialize function once (and only once, since Go 1.19), preferably at the beginning of the code.