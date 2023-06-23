const fs = require("fs");

//Users
const boyanna = {
  name: 'Boyanna',
  phoneNumbers: '0889876543',
  notes: 'Student',
  email: 'boyanna@gmail.com'
};
const yana = {
  name: 'Yana',
  phoneNumbers: '0889876542',
  notes: 'Student',
  email: 'yana@gmail.com'
};
const aleks = {
  name: 'Aleks Karatov',
  phoneNumbers: '0889876540',
  notes: 'Student',
  email: 'aleks@gmail.com'
};
const dobrin = {
  name: 'Dobrin Donev',
  phoneNumbers: '0889876500',
  notes: 'Teacher',
  email: 'dobrin@gmail.com'
};
const nevena = {
  name: 'Nevena',
  phoneNumbers: '0889876501',
  notes: 'Teacher',
  email: 'nevena@gmail.com'
};
const yavor = {
  name: 'Yavor',
  phoneNumbers: '0889876502',
  notes: 'Teacher',
  email: 'yavor@gmail.com'
};
const ivo = {
  name: 'Ivo Markov',
  phoneNumbers: '0889876700',
  notes: 'Employee',
  email: 'ivom@gmail.com'
};
const gabriela = {
  name: 'Gabi',
  phoneNumbers: ['0889876701','0888876781' ],
  notes: 'Employee',
  email: 'gabriela@gmail.com'
};
const konstantin = {
  name: 'Koceto',
  phoneNumbers: '0889876702',
  notes: 'Tenis teacher',
  email: 'konstantind@gmail.com'
};
const mariya = {
  name: 'Mariya',
  phoneNumbers: '0889876703',
  notes: 'Singer',
  email: 'mimi@gmail.com'
};
const kristiyan = {
  name: 'Kris',
  phoneNumbers: '0889876000',
  notes: 'Student',
  email: 'kristiyan@gmail.com'
};
const milena = {
  name: 'Milena',
  phoneNumbers: '0889876001',
  notes: 'Student',
  email: 'milena@gmail.com'
};
const aleksandra = {
  name: 'Aleksandra',
  phoneNumbers: '0889876003',
  notes: 'Actress',
  email: 'aleksandra@gmail.com'
};
const martin = {
  name: 'Martin',
  phoneNumbers: '0889876011',
  notes: 'Student',
  email: 'martin@gmail.com'
};
const mario = {
  name: 'Mario',
  phoneNumbers: ['0889876912', '0889111912'],
  notes: 'Actor',
  email: 'mario@gmail.com',
};

var contacts = [boyanna, yana, aleks, dobrin, 
  nevena, yavor, ivo, gabriela,
  konstantin, mariya, kristiyan, milena, 
  aleksandra, martin, mario,];


//Images
const boyannaImage = {
  name: 'profile1.jpg',
  contentType: 'image/jpeg'
};
const yanaImage = {
  name: 'profile2.jpg',
  contentType: 'image/jpeg'
};
const aleksImage = {
  name: 'profile3.jpg',

  contentType: 'image/jpeg'
};
const dobrinImage = {
  name: 'profile4.jpg',
  contentType: 'image/jpeg'
};
const nevenaImage = {
  name: 'profile5.jpg',
  contentType: 'image/jpeg'
};
const yavorImage = {
  name: 'profile6.jpg',
  contentType: 'image/jpeg'
};
const ivoImage = {
  name: 'default.png',
  contentType: 'image/jpeg'
};
const gabrielaImage = {
  name: 'default.png',
  contentType: 'image/jpeg'
};
const konstantinImage = {
  name: 'default.png',
  contentType: 'image/jpeg'
};
const mariyaImage = {
  name: 'profile10.jpg',
  contentType: 'image/jpeg'
};
const kristiyanImage = {
  name: 'default.png',
  contentType: 'image/jpeg'
};
const milenaImage = {
  name: 'default.png',
  contentType: 'image/jpeg'
};
const aleksndraImage = {
  name: 'profile13.jpg',
  contentType: 'image/jpeg'
};
const martinImage = {
  name: 'default.png',
  contentType: 'image/jpeg'
};
const marioImage = {
  name: 'profile15.jpg',
  contentType: 'image/jpeg'
};


var images = [ boyannaImage,yanaImage,aleksImage,
  dobrinImage,nevenaImage,yavorImage,
  ivoImage, gabrielaImage, konstantinImage,
  mariyaImage, kristiyanImage, milenaImage,
  aleksndraImage, martinImage, marioImage];

function fillImagesData(){
  for (let index = 0; index < images.length; index++) {
    images[index].data = fs.readFileSync("./images/" + images[index].name)
  }
  return images
}


module.exports = {
   contacts, images,
   fillImagesData
};