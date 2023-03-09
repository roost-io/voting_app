

export default function handler(req, res) {
  let ec_server_url = process.env.EC_SERVER_ENDPOINT
  if (!/^(?:f|ht)tps?\:\/\//.test(ec_server_url)) {
    ec_server_url = "http://" + ec_server_url;
    }
    fetch(`${ec_server_url}`, {
      method: 'GET',
    })
      .then((response) => response.json())
      .then((response) => {
        res.status(200).json({ Candidates: response.Candidates })
      })
      .catch((error) => {
        console.error(
          'ballot service is not reachable at http://' + ec_server_url
        );
        res.status(400).json({ Candidates: [] })
      });

    
  
 // res.status(200).json({ name: 'John Doe' })
}
