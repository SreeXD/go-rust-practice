import './style.css'
import { create_user, User } from 'client-wasm'

fetch('http://localhost:5000/', {
  method: "GET"
}).then(res => res.json())
  .then(json => json.data)
  .then(data => {
    const user: User = create_user(data.name, data.age)
    
    alert(user.name() + " is " + user.age().toString() + " years old")
  })