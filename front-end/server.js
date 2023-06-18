// const express = require('express');
// const fetch = require('node-fetch');
import express from 'express';
import fetch from 'node-fetch';
import url from 'url';
import parser from 'body-parser';

// const { fetchContacts } = require('../queries');

const app = express();
const port = 3000;
app.use(parser.urlencoded({ extended: false }))
app.use(parser.json())

app.set('view engine', 'ejs');
app.use(express.static("public"));

//main page
app.get('/', (req, res) => {
  var query = `query {
      contacts() {
        name, id
      }
    }`
  ;
fetchContactsData(query, res)
 
});

//contact page
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



app.get('/contactdelete', (req, res) => {
  var query = `
  mutation{
    deleteContact(id: "`+ req.query.contactid+ `")
  }
` ;

  console.log(query);
  deleteContact(query, res);
  
 
});

app.get('/contactmerge', (req, res) => {
  var query = `
  mutation merge{
      mergeContacts(id: "`+ req.query.contactid+ `"){
        id
      }
    }
  `;

  console.log(query);
  mergeContacts(query, res);
  
 
});

app.get('/createcontactpage', (req, res) => {
  
  res.render('addcontact');
});


app.post('/submit-contactinfo', function(req, res) {
  console.log(req.body);
  var query = `
  mutation {
      createContact(name: "` + req.body.name + `", phones: ["` + req.body.phone + `"], notes: "` + req.body.notes + `", email: "` + req.body.email + `"){
        id
      }
    }
  `;
  console.log(query);
  createContact(res, query);

  res.redirect("/");
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
  var cont = data.data.contact;
  var image = await fetchImage(cont.imageID.toString());

  const contactData ={
    contactid: cont.id ,
    name: cont.name,
    phones : cont.phones,
    email : cont.email,
    notes: cont.notes,
    imageData: image.imageData,
    imageContentType: image.contentType
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
 const imageData = {
    contentType : data.data.image.contentType,
    imageData: data.data.image.imageData,
    name: data.data.image.name
 };
 return imageData;
}


async function deleteContact(query, res) {
  const response = await fetch("http://localhost:50000/query", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ query }),
  });

  const data = await response.json();
  console.log(data);
  res.redirect("/");
  
}


async function mergeContacts(query, res) {
  const response = await fetch("http://localhost:50000/query", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ query }),
  });

  const data = await response.json();
  console.log(data);
  res.redirect(url.format({
    pathname:"/contact",
    query: {
       "contactid": data.data.mergeContacts.id
     }
  }));
  
}


async function createContact(res, query) {
  const response = await fetch("http://localhost:50000/query", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ query }),
  });

  const data = await response.json();
  console.log(data);
 
  
}