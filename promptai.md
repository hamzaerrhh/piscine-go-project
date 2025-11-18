You will write a complete Go project that implements a text-processing tool with the following rules:

Goal:
The program must take two arguments: an input text file and an output text file.
It reads the text, applies a series of transformations, and writes the corrected text.

The solution MUST be structured using good Go practices, modular files, and clear comments.
Everything must be fully documented and auditor-friendly.

Project Requirements:
1. The project MUST be written in Go.
2. The code MUST follow good practices.
3. The code MUST be separated into multiple files:
   - main.go
   - processor.go
   - modifiers.go
   - numbers.go
   - punctuation.go
   - quotes.go
   - articles.go
   - utils.go
4. Each file must contain clean, readable, commented code.
5. You MUST implement unit tests where appropriate.
6. NO external packages are allowed except standard Go packages.

Tool Behavior (very important):
Apply these transformations to the input text:

1. Number conversions:
   - A word followed by "(hex)" must be converted from hexadecimal to decimal.
   - A word followed by "(bin)" must be converted from binary to decimal.

2. Case transformations:
   - "(up)" → uppercase the previous word
   - "(low)" → lowercase the previous word
   - "(cap)" → capitalize the previous word
   - If the tag has a counter like "(up, 3)" then apply the transformation to the previous N words.

3. Punctuation rules:
   - Punctuation marks . , ! ? : ; must attach to the previous word and have a space after.
   - Except for grouped punctuation like "..." or "!?" which must stay together.
   - Example: "hello ,world !!" → "hello, world!!"

4. Quotes rules:
   - Apostrophe quotes ' ' must attach directly to the words inside them.
   - If multiple words are inside the quotes, the quotes move to the edges of the group.
   - Example: "' awesome '" → "'awesome'"
   - Example: "' I am the best '" → "'I am the best'"

5. Article correction:
   - Replace "a" with "an" when the next word begins with a vowel (a, e, i, o, u) or h.
   - Example: "a amazing" → "an amazing"
   - Do NOT replace in words like "a unicorn" (consonant sound).

6. All these transformations must work together on the same line.

Input/Output:
- The program must be executable using "go run . input.txt output.txt".
- The output file must contain the corrected text.

Deliverables:
Produce the full Go project code, separated into the following files:

main.go
processor.go
modifiers.go
numbers.go
punctuation.go
quotes.go
articles.go
utils.go
(optional) processor_test.go

Each file must be output in its own Markdown code block.
All functions must include detailed comments explaining their behavior.

Use Style 3:
Ultra-clean code, extremely clear, highly commented, and auditor-friendly.

Now generate the full multi-file Go project.
