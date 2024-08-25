use std::fs;
// This is one comment
// With 2 lines

fn main() -> std::io::Result<()> {
    // This is another one
    for entry in fs::read_dir(".")? {
        let entry = entry?; // OK?
        println!("{}", entry.path().display());
    }
    Ok(()) // Done
}
