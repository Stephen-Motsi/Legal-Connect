import React, {useState} from 'react';


const Appointments=() =>{
    const[appointments, setAppointments] = useState([]);
    const[ lawyer, setLawyer] =useState("")
    const [name, setName] =useState("");
    const [time, setTime] = useState("");
    const [duration ,setDuration]= useState("");

    const handleSubmit = (e) =>{
        e.preventDefault();
        setAppointments([...appointments, {lawyer,name, time ,duration}]);
        setLawyer("");
        setName("");
       setTime("");
       setDuration("");
    };
    return(
        <div>
          <h2>Book an Appointment</h2>
            <form onSubmit= {handleSubmit}>
              <input type="text" placeholder="Lawyer" value={lawyer} onChange={(e) =>setLawyer(e.target.value)}/>
               <input type="text" placeholder="Name" value={name} onChange={(e) =>setName(e.target.value)}/>
                <input type="text" placeholder="Time" value={time} onChange={(e) =>setTime(e.target.value)}/>
                <input type="text" placeholder="Duration" value={duration} onChange={(e) =>setDuration(e.target.value)}/>
                <button type="submit">Book Appointment</button>
              
           </form>
            <h2>Appointments</h2>
            <ul>
                {appointments.map((appointment, index)=>(
                  <li key ={index}>
                        {appointment.lawyer} , {appointment.name} - {appointment.time} , {appointment.duration}</li>
              ))}
           </ul>
        </div>
    )
}

export default Appointments;