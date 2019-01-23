package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

const ITERATIONS int = 3000000

func main() {

	var str = `[subject]Czym jest Lorem Ipsum?[/subject]
[body]Lorem Ipsum jest tekstem stosowanym jako przykładowy wypełniacz w przemyśle poligraficznym.
Został po raz pierwszy użyty w XV w. przez nieznanego drukarza do wypełnienia tekstem próbnej książki. Pięć wieków później zaczął być używany przemyśle elektronicznym, pozostając praktycznie niezmienionym. Spopularyzował się w latach [years]. [century] w. wraz z publikacją arkuszy Letrasetu, zawierających fragmenty Lorem Ipsum, a ostatnio z zawierającym różne wersje Lorem Ipsum oprogramowaniem przeznaczonym do realizacji druków na komputerach osobistych, jak Aldus PageMaker[/body]`

	var start = time.Now()


	//var result string
	//for i := 0; i < ITERATIONS; i++ {
	//	result = run1(str)
	//}

	c := make(chan string)
	var GOROUTINES_NUMBER = 5000
	for i := 0; i < GOROUTINES_NUMBER; i++ {
		go run3(str, c, GOROUTINES_NUMBER)
	}
	var result = <- c

	fmt.Printf("%v\n", result)
	fmt.Printf("Run took %v seconds\n", time.Now().Unix()-start.Unix())
}

func run3(text string, c chan string, goroutinesNumber int) {
	var result string
	size := ITERATIONS / goroutinesNumber
	for i := 0; i < size; i++ {
		//result = replaceTagWithTwigBlockReplacerVersion(text)
		result = replaceTagWithTwigBlockRegexpVersion(text)
		result = replaceTagWithTwigBraces(result)
	}

	c <- result

}

func run1(text string) string {

	var result = replaceTagWithTwigBlockReplacerVersion(text)
	return replaceTagWithTwigBraces(result)
}

func run2(text string) string {
	var result = replaceTagWithTwigBlockRegexpVersion(text)
	return replaceTagWithTwigBraces(result)
}

func replaceTagWithTwigBlockRegexpVersion(text string) string {
	var re = regexp.MustCompile(`(?ism)(?P<opening>\[(?P<tag>\w+)\])(?P<inner>.*?)(?P<closing>\[\/(?P<tag2>\w+)\])`)
	return re.ReplaceAllString(text, "{% block ${tag} %}${inner}{% endblock %}")
}

func replaceTagWithTwigBlockReplacerVersion(text string) string {

	var re = regexp.MustCompile(`(?ism)\[\/(?P<tag>\w+)\]`)
	tags := re.FindAllString(text, -1)

	for i, tag := range tags {
		tags[i] = strings.Trim(tag, "[/]")
	}

	var result = text

	for _, tag := range tags {

		result = strings.NewReplacer(
			fmt.Sprintf("[%s]", tag), fmt.Sprintf("{%% %s block %%}", tag),
			fmt.Sprintf("[/%s]", tag), "{% endblock %}").Replace(result)
	}

	return result
}

func replaceTagWithTwigBraces(text string) string {
	var re = regexp.MustCompile(`(?ism)\[(?P<tag>\w+)\]`)
	return re.ReplaceAllString(text, "{{ ${tag} }}")
}
