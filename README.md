# Wordle First Guess Optimizer

Welcome to the Wordle First Guess Optimizer repository! If you're an enthusiast of the Wordle game and have been pondering over the best statistical starting word, you've found the right tool.

## Overview

The Wordle First Guess Optimizer analyzes a dictionary of all 5-letter words and ranks them based on the popularity of each letter in each position. Using these probabilities, each word is assigned a "distance". Words with the lowest distance are statistically the most optimal starting guesses for Wordle.

## Key Features

1. **Letter Position Probabilities:** Analysis of the likelihood of each letter appearing at each of the five positions in a word.
2. **Word Ranking:** Rank the entire dictionary based on the computed "distance", providing a list of top words optimal as starting guesses.
3. **Quick and Non-Interactive:** Simply run the script to get the top 10 word recommendations instantly.

## How to Use

1. **Installation:**
    ```bash
    git clone https://github.com/your-github-username/wordle-first-guess-optimizer.git
    cd wordle-first-guess-optimizer
    ```

2. **Compile (if you haven't compiled it yet):**
    ```bash
    go build optimizer.go
    ```

3. **Run the Script:**
    ```bash
    ./optimizer
    ```

    Upon execution, the script will print the top 10 recommended starting words for Wordle based on the analyzed dictionary.

## Built with

- [Go](https://golang.org/)

## Contributions

Your contributions are always welcome! Feel free to fork this repo, create a new branch, make changes, and then submit a pull request.

## Credits

Inspiration for this project comes from the popular game, Wordle. All rights belong to their respective owners. This project is not affiliated with or endorsed by the creators of Wordle.

## License

This project is licensed under the MIT License. Refer to the `LICENSE` file for more details.

## Feedback & Support

If you have any feedback or suggestions, kindly open an issue on this GitHub repository. Enjoy your Wordle games with a strategic edge!

