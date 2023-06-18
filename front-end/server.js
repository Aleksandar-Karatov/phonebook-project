// const express = require('express');
// const fetch = require('node-fetch');
import express from 'express';
import fetch from 'node-fetch';

// const { fetchContacts } = require('../queries');

const app = express();
const port = 3000;


// Define the HTML template
const mainTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Phone book</title>
    <link rel="stylesheet" href="./styles/style.css" > 
    <script src="index.js"></script>
</head>
<body>
      
<div class = "container" >
    <img src="images/istockphoto-1312128591-612x612.jpg">
    
    <div class = "nav">
    <div class= "search">
        <input id = "searchContact" placeholder = "Search">
        <button class = "searchContact"> SEARCH </button>
    </div>
    <button class ="addContact"> ADD NEW CONTACT</button>
    </div>
    <div class = "menu" >
    <h1> CONTACTS </h1>
    <ul> 
        <li>
    <div id = "contactInfo"> 
    <% contacts.forEach(function(contact) { %>
      <span>Name: <%= contact.name %></span><br>
      <br>
  <% }); %>
    </div>
</li>
</ul>
   </div>
</div>

</body>
</html>
`;
app.set('view engine', 'ejs');
app.use(express.static("public"));
// Serve the HTML template
app.get('/', (req, res) => {
  var query = `query {
      contacts() {
        name, id
      }
    }`
  ;
fetchContactsData(query, res)
 
});
app.get('/contact', (req, res) => {
  var query = `
  query {
      contact(id: "`+ req.query.contactid +`"){
        id
        name
        phones
        notes
        email
        imageID
      }
    }`
  ;
fetchContactData(query, res)
 
});






app.listen(port, () => {
  console.log(`Server is running on http://localhost:${port}`);
});

async function fetchContactsData(query, res) {
    const response = await fetch("http://localhost:50000/query", {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ query }),
    });
  
    const data = await response.json();
    console.log(data.data.contacts);
    const contactData = [];
    data.data.contacts.forEach(contact => {
      const c ={
        name: contact.name,
        contactid: contact.id
      };
      contactData.push(c);
    });
    
    res.render('main', {contactData});
    
}

async function fetchContactData(query, res) {
  const response = await fetch("http://localhost:50000/query", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ query }),
  });

  const data = await response.json();
  console.log(data.data);
  var cont = data.data.contact;
  var base64image = fetchImage(data.data.contact.imageID.toString());

  

  const contactData ={
    name: cont.name,
    phones : cont.phones,
    email : cont.email,
    notes: cont.notes,
    image : base64image
  }
  
  res.render('index',{contactData});
  
}

async function fetchImage(imageID){
  var query  =`
  query {
      image(imageID:"`+ imageID  +`"){
        imageId
        name
        imageData
        contentType
        
      }
    }`;
  const response = await fetch("http://localhost:50000/query", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ query }),
  });

  const data = await response.json();
  const base64Image = Buffer.from(data.data.image.imageData).toString('base64');
  return base64Image
}