go run . "" | cat -e

go run . "\n" | cat -e

go run . "Hello\n" | cat -e

go run . "hello" | cat -e

go run . "HeLlO" | cat -e

go run . "Hello There" | cat -e

go run . "1Hello 2There" | cat -e

go run . "{Hello There}" | cat -e

go run . "Hello\nThere" | cat -e

go run . "Hello\n\nThere" | cat -e