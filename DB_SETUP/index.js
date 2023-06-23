const express = require("express");
const mongoose = require("mongoose");
const { User, Image } = require('./schemas');
const app = express();
const dotenv = require('dotenv');
const { createUser, getUserById, updateUser, deleteUser } = require("./CRUD");
const {  contacts, images, fillImagesData} = require("./userData");


async function connect() {
  try {
    const result = dotenv.config();
    if (result.error) {
      throw result.error;
    }

    await mongoose.connect(process.env.DB_CONNECTION, {
      useNewUrlParser: true,
      useUnifiedTopology: true,
    });
    console.log('Connected to MongoDB');
    startServer();
    for (let index = 0; index < contacts.length; index++) {
      var imagesToInsert = fillImagesData();
      var createdUser =  await createUser(contacts[index], imagesToInsert[index]);
      console.log(createdUser);
    }
  } catch (error) {
    console.error(error);
  }
}
  
connect();

function startServer() {
    app.listen(3000, () => {
        console.log("Server started on port 3000");
    });
}


app.get('/users/:id/image', async (req, res) => {
  try {
      const userId = req.params.id;

      const user = await User.findById(userId).exec();

      if (!user) {
          return res.status(404).send('User not found');
      }

      if (!user.image_id) {
          return res.status(404).send('Image not found for the user');
      }

      const image = await Image.findById(user.image_id).exec();

      if (!image) {
          return res.status(404).send('Image not found');
      }

      res.set('Content-Type', image.contentType);

      res.send(image.data);
  } catch (error) {
      console.error('Error retrieving image:', error);
      res.status(500).send('Internal Server Error');
  }
});