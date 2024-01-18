var myHeaders = new Headers();
myHeaders.append("Content-Type", "application/json");

var raw = JSON.stringify({
  "representation": "binary",
  "mutation": "bit_flip",
  "generation": "random",
  "payload": {
    "length": 10,
    "generationsize": 10,
    "desired_score": 0.3
  }
});

var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: raw,
  redirect: 'follow'
};

fetch("http://127.0.0.1:5000/run", requestOptions)
  .then(response => console.log(response))
  .then(result => console.log(result))
  .catch(error => console.log('error', error));