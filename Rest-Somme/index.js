const express = require('express')
const app = express()
const port = 9000

app.get('/somme', (req, res) => {
    let nb1 = req.query.nombre1
    let nb2 = req.query.nombre2
    let result = parseFloat(nb1) + parseFloat(nb2)
    res.status(200).json(result)
})

app.listen(port, () => {
    console.log('Listen')
})