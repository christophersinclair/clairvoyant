# clairvoyant
Clairvoyant is a seemingly simple game of memory.

### Audience
3 - 5 players ages 5 and up.

### Game Flow
Player RED starts with the crystal ball and makes a guess as to which player (could also be themselves) will be holding the ball in some number of turns later. Player RED enters the guess into the crystal ball, and hands it off to another player of their choosing, who also makes a guess.

Each turn starts with an evaluation if anyone was right about the state of this turn (i.e., player holding the ball on turn N). If no one is declared a winner, the player holding the ball enters a guess and passes it on.

The game ends and the winner is declared when the crystal ball ends up in the hands of some player that was guessed to have the ball on that turn.

### Example Scenario
In this example, the game is played with three players: RED, BLUE, and GREEN.
- Player RED guesses that Player GREEN will have the ball in three turns. Player RED then hands the ball off to Player BLUE.
- Player BLUE guesses that Player RED will have the ball in two turns, then hands the ball to Player GREEN.
- Player GREEN guesses that Player BLUE will have the ball in five turns, then hands the ball to Player RED.
- Player BLUE wins because they guessed Player RED will have the ball in two turns, two turns ago.

### Rules
- No player can guess the state the game will be in the next turn (e.g. Player BLUE cannot guess Player RED will be holding the ball in one turn, then hand the ball to Player RED).
