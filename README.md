# clairvoyant
Clairvoyant is a seemingly simple game of memory.

### Audience
3+ players ages 5 and up.

### Game Flow
Player 1 starts with the crystal ball and makes a guess as to which player (could also be themselves) will be holding the ball in some number of turns later. Player 1 enters the guess into the crystal ball, and hands it off to another player of their choosing, who also makes a guess.

Each turn starts with an evaluation if anyone was right about the state of this turn (i.e., player holding the ball on turn N). If no one is declared a winner, the player holding the ball enters a guess and passes it on.

The game ends and the winner is declared when the crystal ball ends up in the hands of some player that was guessed to have the ball on that turn.

### Example Scenario
In this example, the game is played with three players: 1, 2, and 3.
- Player 1 guesses that Player 3 will have the ball in three turns. Player 1 then hands the ball off to Player 2
- Player 2 guesses that Player 1 will have the ball in two turns, then hands the ball to Player 3.
- Player 3 guesses that Player 2 will have the ball in five turns, then hands the ball to Player 1.
- Player 2 wins because they guessed Player 1 will have the ball in two turns, two turns ago.

### Rules
- No player can guess the state the game will be in the next turn (e.g. Player 2 cannot guess Player 1 will be holding the ball in one turn, then hand the ball to Player 1).
