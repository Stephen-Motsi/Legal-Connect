import React, { useState } from "react";

const lawyers = [
  { id: 1, name: "John Smith", specialty: "Criminal Law", location: "New York" },
  { id: 2, name: "Mary Jones", specialty: "Family Law", location: "Los Angeles" },
  { id: 3, name: "David Lee", specialty: "Personal Injury Law", location: "Chicago" },
  { id: 4, name: "Susan Chen", specialty: "Immigration Law", location: "Miami" },
];

const LawyerSearch = () => {
  const [searchTerm, setSearchTerm] = useState("");
  const [searchResults, setSearchResults] = useState([]);
  const [selectedLawyer, setSelectedLawyer] = useState(null);
  const [appointmentDate, setAppointmentDate] = useState(null);


  const handleSearch = (event) => {
    setSearchTerm(event.target.value);

    const results = lawyers.filter((lawyer) =>
      lawyer.name.toLowerCase().includes(searchTerm.toLowerCase())
    );

    setSearchResults(results);
  };

  const handleSelectLawyer = (lawyer) => {
    setSelectedLawyer(lawyer);
  };

  const handleSelectDate = (event) => {
    setAppointmentDate(event.target.value);
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    alert(`Appointment booked with ${selectedLawyer.name} on ${appointmentDate}`);
  };

  return (
    <div>
      <h1>Find a Lawyer</h1>
      <input type="text" placeholder="Search lawyers by name" onChange={handleSearch} />
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Specialty</th>
            <th>Location</th>
            <th>Select</th>
          </tr>
        </thead>
        <tbody>
          {searchResults.map((lawyer) => (
            <tr key={lawyer.id}>
              <td>{lawyer.name}</td>
              <td>{lawyer.specialty}</td>
              <td>{lawyer.location}</td>
              
              <td>
                <button onClick={() => handleSelectLawyer(lawyer)}>Select</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      {selectedLawyer && (
        <form onSubmit={handleSubmit}>
          <h2>Book an appointment with {selectedLawyer.name}</h2>
          <label htmlFor="appointment-date">Select a date:</label>
          <input type="date" id="appointment-date" onChange={handleSelectDate} />
          <button type="submit">Book Appointment</button>
        </form>
      )}
    </div>
  );
};

export default LawyerSearch;