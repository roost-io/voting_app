

export default function handler(req, res) {
    console.log(req)
    fetch(`${process.env.BALLOT_ENDPOINT}`, {
      method: 'POST',
      body: JSON.stringify(JSON.parse(req.body))
    })
      .then((response) => response.json())
      .then((response) => {
        res.status(200).json({success:true})
      })
      .catch((error) => {
        console.error(
          'ballot service is not reachable at http://' + process.env.BALLOT_ENDPOINT
        );
        res.status(400).json({ sucess:false })
      });

    
  
 // res.status(200).json({ name: 'John Doe' })
}
