const express = require('express');
const app = express();


const port = process.env.PORT || 3000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));


app.get('/', (req, res) => {
    res.send('Hello World');
});

app.get('/get', (req, res) => {
    res.send({
        name: 'John',
        age: 30

    });
});

app.post('/post', (req, res) => {
    console.log(req.body);
    res.send(req.body);
});

app.post('/postform', (req, res) => {
    res.send(JSON.stringify(req.body));
});

app.listen(port, () => {
    console.log(`Server is running on port ${port}`);
})


