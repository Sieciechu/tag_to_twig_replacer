<?php
$start = time();
$str = '[subject]Czym jest Lorem Ipsum?[/subject]
[body]Lorem Ipsum jest tekstem stosowanym jako przykładowy wypełniacz w przemyśle poligraficznym.
Został po raz pierwszy użyty w XV w. przez nieznanego drukarza do wypełnienia tekstem próbnej książki. Pięć wieków później zaczął być używany przemyśle elektronicznym, pozostając praktycznie niezmienionym. Spopularyzował się w latach [years]. [century] w. wraz z publikacją arkuszy Letrasetu, zawierających fragmenty Lorem Ipsum, a ostatnio z zawierającym różne wersje Lorem Ipsum oprogramowaniem przeznaczonym do realizacji druków na komputerach osobistych, jak Aldus PageMaker[/body]';

for ($i = 0; $i < 300000; ++$i) {
    $result = run($str);
}
$end = time();
$difference = $end - $start;
var_dump($difference);
var_dump($result);


function run(string $text) : string
{
    $re = '/(?<opening>\[(?<tagName>\w+)\])(?<inner>.*)(?<closing>\[\/(\k<tagName>)\])/ism';
    $result = preg_replace_callback(
        $re,
        function($matches){
            return "{% {$matches['tagName']} %}{$matches['inner']}{% endblock %}";
        },
        $text
    );

    return replace_tag_with_twig_braces($result);
}

function replace_tag_with_twig_braces(string $text) : string
{
    $re = '/(?<tag>\[(?<tagName>\w+)\])/ism';
    return preg_replace_callback(
        $re,
        function($matches) {
            return "{{ {$matches['tagName']} %}}";
        },
        $text
    );
}
