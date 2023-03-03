const express = require('express');
const bodyParser = require('body-parser');

const app = express();
const port = 3000;

// Sample data
const lawyers = [
    { id: 1, name: 'John Doe' },
    { id: 2, name: 'Jane Smith' },
    { id: 3, name: 'Bob ashley' }
];

const appointments = [];

// Middleware
app.use(bodyParser.json());

// Endpoints
app.get('/api/lawyers', (req, res) => {c
    res.json(lawyers);
});

app.get('/api/appointments', (req, res) => {
    res.json(appointments);
});

app.post('/api/appointments', (req, res) => {
    const { lawyerId, date, time, duration } = req.body;

    // Check if the requested time slot is already booked
    const conflictingAppointment = appointments.find(appointment => {
        return (
            appointment.lawyerId === lawyerId &&
            appointment.date === date &&
            appointment.time === time
        );
    });
    if (conflictingAppointment) {
        res.status(409).send('This time slot is already booked');
        return;
    }

    // Create a new appointment
    const id = appointments.length + 1;
    const newAppointment = { id, lawyerId, date, time, duration };
    appointments.push(newAppointment);
    res.json(newAppointment);
});

// Start the server
app.listen(port, () => {
    console.log(`Server listening at http://localhost:${port}`);
});
