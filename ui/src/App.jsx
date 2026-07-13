import { useState, useEffect } from 'react'
import { getLatestDevices } from './services/api';
import DeviceTable from './components/DeviceTable';

function App() {
  const [devices, setDevices] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {

    // fetch data on first render
    getLatestDevices()
      .then((data) => {
        setDevices(data);
        setLoading(false);
      }).catch((error) => {
        setError(error.message);
        setLoading(false);
        console.error(error);
      });
  }, []);

  if (loading) {
    return <h2>Loading...</h2>
  }

  if (error) {
    return <h2>Error: {error}</h2>
  }

  return(
    <>
      <h1>Network Monitor</h1>

      <DeviceTable devices={devices} />
    </>
  )

}

export default App;