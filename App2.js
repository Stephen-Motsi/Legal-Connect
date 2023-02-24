import React, { useState, useEffect } from 'react';

function App() {
  const [lawyers, setLawyers] = useState([]);
  const [appointments, setAppointments] = useState([]);
  const [selectedLawyer, setSelectedLawyer] = useState(null);
  const [selectedDate, setSelectedDate] = useState(null);
  const [selectedTime, setSelectedTime] = useState(null);
  const [selectedDuration, setSelectedDuration] = useState(30);
  const [errorMessage, setErrorMessage] = useState(null);

  useEffect(() => {
    fetch('/api/lawyers')
      .then(response => response.json())
      .then(data => setLawyers(data))
      .catch(error => console.error(error));
  }, []);

  useEffect(() => {
    fetch('/api/appointments')
      .then(response => response.json())
      .then(data => setAppointments(data))
      .catch(error => console.error(error));
  }, []);

  const handleLawyerChange = event => {
    setSelectedLawyer(event.target.value);
  };

  const handleDateChange = event => {
    setSelectedDate(event.target.value);
  };

  const handleTimeChange = event => {
    setSelectedTime(event.target.value);
  };

  const handleDurationChange = event => {
    setSelectedDuration(event.target.value);
  };

  const handleSubmit = event => {
    event.preventDefault();
    const data = {
      lawyerId: selectedLawyer,
      date: selectedDate,
      time: selectedTime,
      duration: selectedDuration
    };
    fetch('/api/appointments', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data)
    })
      .then(response => {
        if (response.ok) {
          setErrorMessage(null);
          return response.json();
        } else if (response.status === 409) {
          setErrorMessage('This time slot is already booked. Please choose another time.');
        } else {
          setErrorMessage('An error occurred. Please try again later.');
          console.error(response);
        }
      })
      .then(data => {
        if (data) {
          setAppointments([...appointments, data]);
          setSelectedLawyer(null);
          setSelectedDate(null);
          setSelectedTime(null);
          setSelectedDuration(30);
        }
      })
      .catch(error => console.error(error));
  };

  return (
    <div>
      <h1>Book an Appointment</h1>
      {errorMessage && <div className="error">{errorMessage}</div>}
      <form onSubmit={handleSubmit}>
        <label>
          Lawyer:
          <select value={selectedLawyer} onChange={handleLawyerChange}>
            <option value=""></option>
            {lawyers.map(lawyer => (
              <option key={lawyer.id} value={lawyer.id}>{lawyer.name}</option>
            ))}
          </select>
        </label>
        <br />
        <label>
          Date:
          <input type="date" value={selectedDate} onChange={handleDateChange} />
        </label>
        <br />
        <label>
          Time:
          <input type="time" value={selectedTime} onChange={handleTimeChange} />
        </label>
        <br />
        <label>
          Duration (in minutes):
          <input type="number" value={selectedDuration} onChange={handleDurationChange} />
        </label>
        <br />

         
            <ul>
                {appointments.map((appointment, index)=>(
                  <li key ={index}>
                        {appointment.name} - {appointment.time}</li>
                ))}
                        </ul>
      </form>
      <h2>Booked</h2>
      </div>
              
  )
 }

  export default App;
