# Become a Developer Test Task

Вам потрібно розробити алгоритм програми, яка повинна виконувати наступне:

програма приймає на вхід довільний текст і знаходить в кожному слові цього тексту найперший символ, який більше НЕ повторюється в аналізуємому слові
далі із отриманого набору символів програма повинна вибрати перший унікальний (тобто той, який більше не зустручається в наборі) і повернути його.
Наприклад, якщо програма отримує на вхід текст нижче:

```
The Tao gave birth to machine language.
Machine language gave birth to the assembler.
The assembler gave birth to the compiler.
Now there are ten thousand languages.
Each language has its purpose, however humble.
Each language expresses the Yin and Yang of software.
Each language has its place within the Tao.
But do not program in COBOL if you can avoid it.
-- Geoffrey James, "The Tao of Programming"
```

то повинна повернути вона символ "m".

## Solution

The solution is as follows:
- read the text
- normalize it into Unicode NFC normal form
- go over all contiguous sequences of Unicode category L (Letter) characters and
  find the first unique character in each of them, storing it in an array
- find the first unique character in the resulting array

The general idea for the algorithm that finds the first unique character is as follows:
- iterate over the characters and update a dictionary of type
  ```character -> {position: int, is_unique: bool}```,
  where ```position``` is the index of the character in the word.
- read the dictionary and find the character for which the position is lowest and is_unique is true.

The Go version is much more efficient, because it doesn't process the whole text
multiple times and doesn't use regular expressions. I've used a library
for wrapping ```os.Stdin``` into a reader that does Unicode normalization as it
reads, and implemented a custom ```ScanStrippedWords``` function for
```bufio.Scanner```, which splits the text into words that consist of contiguous
Unicode Letter (L) characters as is scans.

The JavaScript version runs on each ```textarea#text``` update, uses
```String.prototype.normalize``` to normalize the input, then uses a regular
expressions to split it into words, and runs the algorithm.

## How to compile and run the solution

### Go

```bash
# Clone this repository
git clone https://github.com/emilia-miki/become-a-developer-test-task
cd become-a-developer-test-task/go # Enter the program's directory
go install # Install dependencies
go build # Build the program
./go # Run it
# Enter text here
# Press ^D on an empty line to finish typing and get the result, or pipe a file:
# cat my_text.txt | ./go
```

### JavaScript

- Clone this repository: ```git clone https://github.com/emilia-miki/become-a-developer-test-task```
- Open the ```index.html``` file with your default browser:
  - Linux: ```xdg-open become-a-developer-test-task/js/index.html```
  - MacOS: ```open become-a-developer-test-task/js/index.html```
  - Windows: do it in File Explorer i guess (i haven't used windows for many years)
