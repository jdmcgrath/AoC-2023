use std::{fs, collections::HashMap};

struct Element {
    left: String,
    right: String,
}

fn main() {
    let contents = fs::read_to_string("./input.txt")
        .expect("Should have been able to read the file");

    let parts: Vec<&str> = contents.split("\n\n").collect(); 
    let instructions = parse_instructions(parts[0].to_string());
    let elements = parse_elements(parts[1].to_string());

    let mut current_position = String::from("AAA");
    let end_position = String::from("ZZZ");

    let mut attempts = 0;
    // Use a loop to keep running the instructions until we reach the end position
    while current_position != end_position {
        for instruction in instructions.iter() {
            attempts += 1;
            println!("Attempt number: {}", attempts);
            current_position = follow_instruction(instruction, &elements, &current_position).unwrap_or_else(|| {
                panic!("Invalid position encountered: {}", current_position)
            });

            if current_position == end_position {
                break;  // Exit the loop if we reach the end position
            }
        }
    }

    println!("Reached the end position: {}", current_position);
}


fn parse_instructions(instructions: String) -> Vec<char> {
    instructions.chars().collect()
}

fn parse_elements(elements: String) -> HashMap<String, Element> {
    let mut elements_map = HashMap::new();

    // Iterate over each line in the elements string
    for line in elements.lines() {
        // Split the line at ` = ` to get the key and the value part
        let parts: Vec<&str> = line.split(" = ").collect();
        if parts.len() == 2 {
            let key = parts[0];
            let values = parts[1].trim_start_matches('(').trim_end_matches(')');
            let values_parts: Vec<&str> = values.split(", ").collect();

            // Check if we have exactly two values after splitting
            if values_parts.len() == 2 {
                elements_map.insert(
                    key.to_string(),
                    Element {
                        left: values_parts[0].to_string(),
                        right: values_parts[1].to_string(),
                    },
                );
            }
        }
    }

    elements_map
}

fn follow_instruction(instruction: &char, paths: &HashMap<String, Element>, current_position: &str) -> Option<String> {
    if let Some(current_element) = paths.get(current_position) {
        if *instruction == 'L' {
            Some(current_element.left.clone())
        } else {
            Some(current_element.right.clone())
        }
    } else {
        None
    }
}

