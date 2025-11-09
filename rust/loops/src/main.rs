fn main() {
    let mut counter = 0;

    // loop as an expression which assigns a value to result
    let result = loop {
        counter += 1;

        if counter == 10 {
            break counter * 2;
        }
    };

    println!("The result is {result}");

    // loop over contents of an array
    let a = [10, 20, 30, 40, 50];

    for element in a {
        println!("The value is: {element}");
    }

    // loop over a range
    for number in (1..4).rev() {
        println!("{number}");
    }
}
