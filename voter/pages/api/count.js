

export default function handler(req, res) {
    let ballot_endpoint = process.env.BALLOT_ENDPOINT
    if (!/^(?:f|ht)tps?\:\/\//.test(ballot_endpoint)) {
      ballot_endpoint = "http://" + ballot_endpoint;
      }
    if(req.method === 'POST'){
    fetch(`${ballot_endpoint}`, {
      method: 'POST',
      body: JSON.stringify(JSON.parse(req.body))
    })
      .then((response) => response.json())
      .then((response) => {
        res.status(200).json({success:true})
      })
      .catch((error) => {
        console.error(
          'ballot service is not reachable at http://' + ballot_endpoint
        );
        res.status(400).json({ sucess:false })
      });
    }else {
      fetch(`${ballot_endpoint}`, {
        method: 'GET',
      })
        .then((response) => response.json())
        .then((response) => {
          res.status(200).json(response)
        })
        .catch((error) => {
          console.error(
            'ballot service is not reachable at http://' + ballot_endpoint
          );
          res.status(400).json({  })
        });
    }

    
  
 // res.status(200).json({ name: 'John Doe' })
}
