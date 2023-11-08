const express = require('express');
const app = express();
const port = process.env.PORT || 5000;

app.get('/', (req, res) => {
  res.send('<h2> Test app2 (node.js) </h2>');
});

app.listen(port, () => {
  console.log(`Server running on port ${port}`);
});
