//import React from 'react';
import React, { useState } from 'react';
import axios from 'axios';

function App() {

  const [inputValue, setInputValue] = useState('');
  const [items, setItems] = useState([]);

  const handleInputChange = (event) => {
    setInputValue(event.target.value);
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    axios
        .get(`http://localhost:8000/search?q=${inputValue}`)
        .then((response) => {
          //console.log(response.data);
          setItems(response.data);
        })
        .catch((error) => console.log(error));
  };

  return (
  <div>
    <form onSubmit={handleSubmit}>
      <label>
        Search items:
        <input type="text" value={inputValue} onChange={handleInputChange} />
      </label>
      <button type="submit">Search</button>
    </form>
    {items.length > 0 ? (
        <ul>
          {items.map((item) => (
              <li key={item._id}>{item.name}</li>
          ))}
        </ul>
    ) : null}
  </div>
  );
}

export default App;
