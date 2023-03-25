import React, {useState, useEffect} from 'react'
import axios from 'axios'

const TimeBillingModule =() =>{
    const [clients, setClients] = useState([]);
    const [selectedClient, setSelectedClient] = useState(null);
    const [hours, setHours] = useState("");
    const [rate, setRate] = useState("");
    const [total, setTotal] =useState(0);
    const [billed, setBilled] = useState(false);


    useEffect(()=>{
        axios.get("/api/clients").then((response) =>{
            setClients(response.data);
        });
    },[])

    const handleClientChange= (e) =>{
        setSelectedClient(e.target.value);

    };
    const handleHoursChange = (e)=>{
        setHours(e.target.value);
    };
    const handleRateChange = (e) =>{
        setRate(e.target.value);
    };
     const handleCalculateTotal =()=>{
        const hoursNum = parseFloat(hours);
        const rateNum = parseFloat(rate);
        const totalNum = hoursNum * rateNum;
        setTotal(totalNum);
         };

         const handleBill = async() =>{
            try{
                const response = await axios.post("api/bill",{
                    client_id :selectedClient,
                    hours: parseFloat(hours),
                    rate: parseFloat(rate),
                    total: parseFloat(total)

                });
                if (response.status === 200){
                    setBilled(true);
                } 
            }
            catch (error){
                console.log(error);
            }
         };
         return(
            <div>
                <h2>Time and Billing Module</h2>
                {billed ?(
                    <div>
                        <h3>Billing complete</h3>
                        <button onClick = {() => setBilled(false)}>New bill</button>
                        </div>
                ):(
                    <div>
                        <label>
                            Client:
                            <select value={selectedClient} onChange = {handleClientChange}>
                                <option value="" disabled>
                                    Select client
                                </option>
                                {clients.map((client) =>(
                                    <option key ={client.id} value ={client.id}>
                                        {client.name}
                                    </option>
                                ))}
                            </select>
                        </label>
                        <br/>
                        <label>
                            Hours:
                            <input type="number" value={hours} onChange= {handleHoursChange}/>

                        </label>
                        <br/>
                        <label>
                            Rate:
                            <input type="number" value={rate} onChange={handleRateChange}/>
                        </label>
                        <br/>
                        <button onClick= {handleCalculateTotal}>Calculate Total</button>
                        <br/>
                        <label>
                            Total:
                            <input type="number" value={total} disabled/>

                        </label>
                        <br/>
                        <button onClick ={handleBill}>Bill</button>
                        </div>

                )}
                                
    
            </div>
         );
    }
      export default TimeBillingModule;   
