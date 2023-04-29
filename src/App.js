import React, { useState } from 'react';
import axios from 'axios';

function App() {
  const [inputValue, setInputValue] = useState('');
  const [items, setItems] = useState([]);

  const handleInputChange = (event) => {
    setInputValue(event.target.value);
  };

  const handleSubmitSearch = (event) => {
    event.preventDefault();
    axios
        .get(`http://localhost:8000/search?q=${inputValue}`)
        .then((response) => setItems(response.data))
        .catch((error) => console.log(error));
  };

  const [newItemSku, setNewItemSku] = useState('');
  const [newItemName, setNewItemName] = useState('');
  const [newItemCategory, setNewItemCategory] = useState('');

  const handleAddItem = (event) => {
    event.preventDefault();
    const newItem = {
      sku: newItemSku,
      name: newItemName,
      category: newItemCategory,
    };
    axios
        .post('http://localhost:8000/add-item', newItem)
        .then((response) => {
          console.log(response);
          setNewItemSku('');
          setNewItemName('');
          setNewItemCategory('');
        })
        .catch((error) => console.log(error));
  };

  return (
      <div>
        <form onSubmit={handleSubmitSearch}>
          <label>
            Search items:
            <input
                type="text"
                value={inputValue}
                onChange={handleInputChange}
            />
          </label>
          <button type="submit">Search</button>
        </form>
        {items.length > 0 ? (
            <ul>
              {items.map((item) => (
                  <li key={item._id}>
                    {item.name} ({item.category})
                  </li>
              ))}
            </ul>
        ) : null}

        <div style={{ margin: '2em 0' }}>
          <hr style={{ margin: '0' }} />
        </div>

        <form onSubmit={handleAddItem}>
          <label>
            SKU:
            <input
                type="text"
                value={newItemSku}
                onChange={(e) => setNewItemSku(e.target.value)}
            />
          </label>
          <label>
            Name:
            <input
                type="text"
                value={newItemName}
                onChange={(e) => setNewItemName(e.target.value)}
            />
          </label>
          <label>
            Category:
            <input
                type="text"
                value={newItemCategory}
                onChange={(e) => setNewItemCategory(e.target.value)}
            />
          </label>
          <button type="submit">Add Item</button>
        </form>
      </div>
  );
}

export default App;
