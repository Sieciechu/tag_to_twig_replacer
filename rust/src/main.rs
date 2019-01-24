#[macro_use]
extern crate lazy_static;
extern crate regex;

use regex::Captures;
use regex::Regex;
use std::thread;

const NUMBER_OF_ITERATIONS: i32 = 300000;
const THREADS_NUMBER: i32 = 4;

fn main() {
    let start = std::time::SystemTime::now();

    let text = "[subject]Czym jest Lorem Ipsum?[/subject]
[body]Lorem Ipsum jest tekstem stosowanym jako przykładowy wypełniacz w przemyśle poligraficznym.
Został po raz pierwszy użyty w XV w. przez nieznanego drukarza do wypełnienia tekstem próbnej książki. Pięć wieków później zaczął być używany przemyśle elektronicznym, pozostając praktycznie niezmienionym. Spopularyzował się w latach [years]. [century] w. wraz z publikacją arkuszy Letrasetu, zawierających fragmenty Lorem Ipsum, a ostatnio z zawierającym różne wersje Lorem Ipsum oprogramowaniem przeznaczonym do realizacji druków na komputerach osobistych, jak Aldus PageMaker[/body]
";
    let mut result: String = String::new();

    let mut handles = vec![];
    for _ in 0..THREADS_NUMBER {
        let handle = thread::spawn(move || {
            let mut result = String::new();
            let size = NUMBER_OF_ITERATIONS / THREADS_NUMBER;
            for _ in 0..size {
                result = run(text);
            }
            result
        });
        handles.push(handle);
    }

    for handle in handles {
        result = handle.join().unwrap();
    }

    let end = std::time::SystemTime::now();
    println!("It took time: {:?}", end.duration_since(start));
    dbg!(result);
}

fn run(text: &str) -> String {
    let result = replace_tag_with_twig_block(&text);
    return replace_tag_with_twig_braces(&result);
}

fn replace_tag_with_twig_block(text: &str) -> String {
    lazy_static! {
        static ref REGEX: Regex = Regex::new(r"(?ism)\[/(?P<tag>\w+)\]").unwrap();
    }

    let tags = REGEX
        .captures_iter(text)
        .map(|captures: Captures| captures.name("tag").unwrap().as_str());

    let mut result: String = String::from(text);
    for tag in tags {
        let new_opening_tag = format!("{{% block {} %}}", tag);
        result = result
            .replace(&format!("[{}]", tag), &new_opening_tag)
            .replace(&format!("[/{}]", tag), "{% endblock %}");
    }

    result

    ////      Below was 1st solution, but it makes running from 2s to ~150s for 300000 iterations
    //        let regex= format!(r"(?ism)(\[{tag}\])(?P<inner>.*)(\[/{tag}\])", tag=searched_block_name);
    //        let re = Regex::new(&regex).unwrap();
    //        let result = re.replace(text, |caps: &Captures| {
    //            format!("{}{}{{% blockend %}}", opening, &caps["inner"])
    //        });
    //
    //        result.to_string()
}

fn replace_tag_with_twig_braces(text: &str) -> String {
    lazy_static! {
        static ref RE: Regex = Regex::new(r"(?ism)(\[(?P<tagName>\w+)\])").unwrap();
    }

    let result = RE.replace_all(text, |caps: &Captures| {
        format!("{{{{ {} }}}}", &caps["tagName"])
    });

    result.to_string()
}
