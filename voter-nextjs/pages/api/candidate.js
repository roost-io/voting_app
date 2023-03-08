console.log(process.env)

export default function handler(req, res) {
    fetch(`http://${process.env.ec_server_endpoint}`, {
      method: 'GET',
    })
      .then((response) => response.json())
      .then((response) => {
        res.status(200).json({ Candidates: response.Candidates })
      })
      .catch((error) => {
        console.error(
          'ballot service is not reachable at http://' + process.env.ec_server_endpoint
        );
        res.status(400).json({ Candidates: [] })
      });

    
  
 // res.status(200).json({ name: 'John Doe' })
}
