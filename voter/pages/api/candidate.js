

export default function handler(req, res) {
    fetch(`${process.env.EC_SERVER_ENDPOINT}`, {
      method: 'GET',
    })
      .then((response) => response.json())
      .then((response) => {
        res.status(200).json({ Candidates: response.Candidates })
      })
      .catch((error) => {
        console.error(
          'ballot service is not reachable at http://' + process.env.EC_SERVER_ENDPOINT
        );
        res.status(400).json({ Candidates: [] })
      });

    
  
 // res.status(200).json({ name: 'John Doe' })
}
