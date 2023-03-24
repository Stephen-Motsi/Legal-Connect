import React, { useState, useEffect } from 'react';

function App() {
  const [cases, setCases] = useState([]);
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');

  // Fetch cases from backend API on mount
  useEffect(() => {
    fetch('/api/cases')
      .then(response => response.json())
      .then(data => setCases(data))
      .catch(error => console.error(error));
  }, []);

  const handleSubmit = (event) => {
    event.preventDefault();
    const newCase = { title, description };
    fetch('/api/cases', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newCase),
    })
      .then(response => response.json())
      .then(data => {
        setCases([...cases, data]);
        setTitle('');
        setDescription('');
      })
      .catch(error => console.error(error));
  };

  const handleDelete = (id) => {
    fetch(`/api/cases/${id}`, { method: 'DELETE' })
      .then(() => {
        const newCases = cases.filter(caseData => caseData.id !== id);
        setCases(newCases);
      })
      .catch(error => console.error(error));
  };

  return (
    <div className="App">
      <h1>Case Management System</h1>
      <form onSubmit={handleSubmit}>
        <label htmlFor="title">Title:</label>
        <input type="text" id="title" value={title} onChange={(event) => setTitle(event.target.value)} />
        <label htmlFor="description">Description:</label>
        <textarea id="description" value={description} onChange={(event) => setDescription(event.target.value)} />
        <button type="submit">Add Case</button>
      </form>
      <ul>
        {cases.map(caseData => (
          <li key={caseData.id}>
            <h3>{caseData.title}</h3>
            <p>{caseData.description}</p>
            <button onClick={() => handleDelete(caseData.id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
