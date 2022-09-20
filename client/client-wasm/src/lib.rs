use wasm_bindgen::prelude::*;

#[wasm_bindgen]
pub struct User {
    name: String,
    age: i32 
}

#[wasm_bindgen]
impl User {
    pub fn new(name: String, age: i32) -> User {
        User {
            name: name,
            age: age
        }
    }

    pub fn name(&self) -> String {
        String::clone(&self.name)
    }

    pub fn age(&self) -> i32 {
        self.age
    }
}

#[wasm_bindgen]
pub fn create_user(name: String, age: i32) -> User {
    User::new(name, age)
}